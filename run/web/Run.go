package web

import (
	"log"
	"net/http"

	"github.com/ze0nni/kodb/internal/engine"
)

func Run(engine engine.Engine) error {
	server := newServer(engine)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws/", clientHandle(server))

	log.Println("http server started on :8000")
	go server.listen()
	err := http.ListenAndServe(":8000", nil)

	return err
}
