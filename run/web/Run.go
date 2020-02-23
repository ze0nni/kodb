package web

import (
	"log"
	"net/http"
)

func Run() error {
	server := newServer()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws/", clientHandle(server))

	log.Println("http server started on :8000")
	go server.listen()
	err := http.ListenAndServe(":8000", nil)

	return err
}
