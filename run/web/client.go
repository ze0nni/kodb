package web

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"github.com/ze0nni/kodb/run/web/msg"
)

type ClientID int

type serverController interface {
	ClientConnected(client *clientConnection)
	ClientDisconnected(client *clientConnection)

	GetSchema(ClientID)
	GetLibraryRows(ClientID, string)
	NewRow(ClientID, string)
}

func clientHandle(server serverController) func(http.ResponseWriter, *http.Request) {
	var clientIDCounter = 0

	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

		if nil != ws {
			defer ws.Close()
		}

		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(w, "Not a websocket handshake", 400)
			return
		} else if err != nil {
			return
		}

		clientIDCounter++

		client := &clientConnection{
			id:     ClientID(clientIDCounter),
			server: server,
			ws:     ws,

			responseCh: make(chan interface{}),
		}

		client.listen()
	}

}

type clientConnection struct {
	id     ClientID
	server serverController
	ws     *websocket.Conn

	responseCh chan interface{}
}

func (client *clientConnection) listen() {
	client.server.ClientConnected(client)
	defer client.server.ClientDisconnected(client)

	go client.write()
	client.read()
}

func (client *clientConnection) read() {
	log.Printf("[%d] Connected", client.id)
	for {
		msgType, msgRaw, err := client.ws.ReadMessage()
		if err != nil {
			log.Printf("[%d] Error read message %d: %s", client.id, msgType, err)
			break
		}
		msg, err := simplejson.NewJson(msgRaw)
		if err != nil {
			log.Printf("[%d] Error decode message: %s", client.id, err)
			log.Printf("[%d] body: %s", client.id, msgRaw)
			break
		}
		log.Printf("[%d] Message recieved", client.id)
		client.clientRecieveMessage(msg)
	}
	log.Printf("[%d] Disconnected", client.id)
}

func (client *clientConnection) clientRecieveMessage(
	msg *simplejson.Json,
) {
	commandRaw, ok := msg.CheckGet("command")
	if false == ok {
		log.Printf("[%d] Broken message %s", client.id, msg)
		return
	}
	command, err := commandRaw.String()
	if nil != err {
		log.Printf("[%d] Broken message %s", client.id, msg)
		return
	}

	switch command {
	case "getSchema":
		client.server.GetSchema(client.id)
	case "getLibraryRows":
		libraryName := msg.Get("library").MustString()
		client.server.GetLibraryRows(client.id, libraryName)
	case "newRow":
		libraryName := msg.Get("library").MustString()
		client.server.NewRow(client.id, libraryName)
	default:
		log.Printf("[%d] Unknown message %s", client.id, msg)
	}
}

func (client *clientConnection) write() {
	for {
		select {
		case msg := <-client.responseCh:
			err := client.ws.WriteJSON(msg)
			if err != nil {
				log.Print("[%d] Message sending error: %s", client.id, err)
				break
			}
			log.Printf("[%d] Message sended", client.id)
		}
	}
}

func (client *clientConnection) SetSchema(msg *msg.SetSchemaMsg) {
	client.responseCh <- msg
}

func (client *clientConnection) SetLibraryRows(msg *msg.SetLibraryRowsMsg) {
	client.responseCh <- msg
}

func (client *clientConnection) NewRow(msg *msg.NewRowMsg) {
	client.responseCh <- msg
}
