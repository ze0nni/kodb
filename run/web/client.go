package web

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

func clientHandle() func(http.ResponseWriter, *http.Request) {
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

		clientRead(ws)
	}

}

func clientRead(
	ws *websocket.Conn,
) {
	for {
		msgType, msgRaw, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error read message %d: %s", msgType, err)
			break
		}
		msg, err := simplejson.NewJson(msgRaw)
		if err != nil {
			log.Printf("Error decode message: %s", err)
			log.Printf("body: %s", msgRaw)
			break
		}
		log.Printf("recieve: %s", msg)
	}
	log.Print("Disconnected")
}
