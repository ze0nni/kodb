package run

import (
	"log"
	"net/http"
)

func Web() error {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return err
	}
	return nil
}
