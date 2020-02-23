package web

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

type serverController interface {
	ClientConnected(client *clientConnection)
	ClientDisconnected(client *clientConnection)
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
			id:     clientIDCounter,
			server: server,
			ws:     ws,
		}

		client.listen()
	}

}

type clientConnection struct {
	id     int
	server serverController
	ws     *websocket.Conn
}

func (client *clientConnection) listen() {
	client.server.ClientConnected(client)
	defer client.server.ClientDisconnected(client)

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
		log.Printf("[%d] Recieve: %s", client.id, msg)
		client.clientRecieveMessage(msg)
	}
	log.Printf("[%d] Disconnected", client.id)
}

func (client *clientConnection) clientRecieveMessage(
	msg *simplejson.Json,
) {

}
