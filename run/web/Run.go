package web

import (
	"log"
	"net/http"
)

func Run() error {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws/", clientHandle())

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)

	return err
}
