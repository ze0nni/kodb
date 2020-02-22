package run

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func newServer() *server {
	s := &server{
		clients:     []*client{},
		addClientCh: make(chan *client),
		delClientCh: make(chan *client),
	}

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.Handle("/ws/", websocket.Handler(s.handleWebsocket))

	return s
}

type server struct {
	clients []*client

	addClientCh chan *client
	delClientCh chan *client
}

func (s *server) handleWebsocket(conn *websocket.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			//s.errCh <- err
			log.Panic(err)
		}
	}()
	c := newClient(conn, s)
	s.addClientCh <- c
	c.Listen()
	log.Print("Disconnected")
}

func (s *server) listen() {
	for {
		select {
		case newClient := <-s.addClientCh:
			log.Print("New client")
			s.clients = append(s.clients, newClient)
		}
	}
}

func (s *server) run() error {
	log.Println("ws listening")
	go s.listen()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	return err
}
