package web

import "log"

func newServer() *serverInstance {
	return &serverInstance{
		clientConnectedCh:    make(chan *clientConnection),
		clientDisconnectedCh: make(chan *clientConnection),
		clients:              make(map[int]*clientConnection),
	}
}

type serverInstance struct {
	clientConnectedCh    chan *clientConnection
	clientDisconnectedCh chan *clientConnection
	clients              map[int]*clientConnection
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

func (server *serverInstance) listen() {
	for {
		select {
		case client := <-server.clientConnectedCh:
			server.clientConnected(client)
		case client := <-server.clientDisconnectedCh:
			server.clientDisconnected(client)
		}
	}
}
