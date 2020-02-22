package run

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

func newClient(
	conn *websocket.Conn,
	serverHandle serverHandle,
) *client {
	return &client{
		conn:   conn,
		handle: serverHandle,

		messagesCh: make(chan *Message),
		doneCh:     make(chan struct{}),
	}
}

type client struct {
	conn   *websocket.Conn
	handle serverHandle

	messagesCh chan *Message
	doneCh     chan struct{}
}

type serverHandle interface {
	//Del(c *client)
}

func (c *client) Listen() {
	go c.listenWrite()
	c.listenRead()
}
func (c *client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.messagesCh:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.conn, msg)

		// receive done request
		case <-c.doneCh:
			//TODO: c.server.Del(c)
			c.doneCh <- struct{}{} // for listenRead method
			return
		}
	}
}

func (c *client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			//TODO: c.handle.Del(c)
			c.doneCh <- struct{}{} // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg Message
			err := websocket.JSON.Receive(c.conn, &msg)
			if err == io.EOF {
				c.doneCh <- struct{}{}
			} else if err != nil {
				log.Panic(err)
				//c.handle.Err(err)
			} else {
				//c.handle.SendAll(&msg)
				log.Print("Message recieved")
			}
		}
	}
}
