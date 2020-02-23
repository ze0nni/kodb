package web

import (
	"log"

	"github.com/ze0nni/kodb/run/web/msg"

	"github.com/ze0nni/kodb/internal/engine"
)

func newServer(engine engine.Engine) *serverInstance {
	return &serverInstance{
		engine:               engine,
		clientConnectedCh:    make(chan *clientConnection),
		clientDisconnectedCh: make(chan *clientConnection),
		clients:              make(map[ClientID]*clientConnection),

		msgGetSchemaCh: make(chan msgGetSchema),
	}
}

type msgGetSchema = struct{ ClientId ClientID }

type serverInstance struct {
	engine engine.Engine

	clientConnectedCh    chan *clientConnection
	clientDisconnectedCh chan *clientConnection
	clients              map[ClientID]*clientConnection

	msgGetSchemaCh chan msgGetSchema
}

// ClientConnected
func (server *serverInstance) ClientConnected(client *clientConnection) {
	server.clientConnectedCh <- client
}

func (server *serverInstance) clientConnected(client *clientConnection) {
	if _, ok := server.clients[client.id]; false == ok {
		server.clients[client.id] = client
	} else {
		log.Panicf("Client [%d] already conected", client.id)
	}
}

// ClientDisconnected
func (server *serverInstance) ClientDisconnected(client *clientConnection) {
	server.clientDisconnectedCh <- client
}

func (server *serverInstance) clientDisconnected(client *clientConnection) {
	if storedClient, ok := server.clients[client.id]; ok && storedClient == client {
		delete(server.clients, client.id)
	} else {
		log.Panicf("Can't disconnec for client [%d]", client.id)
	}
}

func (server *serverInstance) GetSchema(clientId ClientID) {
	server.msgGetSchemaCh <- msgGetSchema{clientId}
}

func (server *serverInstance) getSchema(m msgGetSchema) {
	if client, ok := server.clients[m.ClientId]; ok {
		client.SetSchema(msg.SetSchemaMsgFromEngine(server.engine))
	}

}

//listen
func (server *serverInstance) listen() {
	for {
		select {
		case client := <-server.clientConnectedCh:
			server.clientConnected(client)
		case client := <-server.clientDisconnectedCh:
			server.clientDisconnected(client)

		case msg := <-server.msgGetSchemaCh:
			server.getSchema(msg)
		}
	}
}
