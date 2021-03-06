package web

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"github.com/ze0nni/kodb/run/web/msg"
)

type ClientID int

type ClientMsg interface {
	Perform(*serverInstance) error
}

type serverController interface {
	ClientConnected(client *clientConnection)
	ClientDisconnected(client *clientConnection)

	Perform(ClientMsg, error)

	GetSchema(ClientID)
	GetLibraryRows(ClientID, string)
	NewRow(msgNewRow)
	DeleteRow(ClientID, string, string)
	UpdateValue(ClientID, string, string, string, string)

	AddLibrary(ClientID, string)
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

	done := make(chan struct{})
	go client.write(done)
	client.read(done)
}

func (client *clientConnection) read(done chan struct{}) {
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
	done <- struct{}{}
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
	case "fetch":
		client.server.Perform(&(msgFetch{client.id}), nil)
	case "getTypes":
		client.server.Perform(&(msgGetTypes{client.id}), nil)
	case "newField":
		client.server.Perform(msgNewFieldFromJson(client.id, msg))
	case "deleteField":
		client.server.Perform(msgDeleteFieldFromJson(client.id, msg))
	case "updateField":
		client.server.Perform(msgUpdateFieldFromJson(client.id, msg))
	case "getSchema":
		client.server.GetSchema(client.id)
	case "getLibraryRows":
		libraryName := msg.Get("library").MustString()
		client.server.GetLibraryRows(client.id, libraryName)
	case "newRow":
		m, err := msgNewRowFromJson(client.id, msg)
		if nil != err {
			log.Printf("[%d] %s | %s", client.id, msg, err)
			return
		}
		client.server.NewRow(m)
	case "deleteRow":
		libraryName := msg.Get("library").MustString()
		rowID := msg.Get("rowId").MustString()
		client.server.DeleteRow(client.id, libraryName, rowID)
	case "updateValue":
		libraryName := msg.Get("library").MustString()
		rowID := msg.Get("rowId").MustString()
		columnID := msg.Get("columnId").MustString()
		value := msg.Get("value").MustString()
		client.server.UpdateValue(client.id, libraryName, rowID, columnID, value)
	case "swapRows":
		client.server.Perform(msgSwapRowsFromJson(client.id, msg))
	case "updateRowCase":
		client.server.Perform(msgUpdateRowCaseFromJson(client.id, msg))
	case "addLibrary":
		libraryName := msg.Get("library").MustString()
		client.server.AddLibrary(client.id, libraryName)
	default:
		log.Printf("[%d] Unknown message %s", client.id, msg)
	}
}

func (client *clientConnection) write(done chan struct{}) {
WriteLoop:
	for {
		select {
		case <-done:
			break WriteLoop
		case msg := <-client.responseCh:
			err := client.ws.WriteJSON(msg)
			if err != nil {
				log.Print("[%d] Message sending error: %s", client.id, err.Error())
				break WriteLoop
			}
			log.Printf("[%d] Message sended", client.id)
		}
	}
}

func (client *clientConnection) Send(msg *simplejson.Json) {
	client.responseCh <- msg
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

func (client *clientConnection) DeleteRow(msg *msg.DeleteRowMsg) {
	client.responseCh <- msg
}

func (client *clientConnection) UpdateValue(msg *msg.UpdateValueMsg) {
	client.responseCh <- msg
}
