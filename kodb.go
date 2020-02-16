package main

import (
	"log"

	"github.com/ze0nni/kodb/run"
)

func main() {
	err := run.Web()
	if nil != err {
		log.Fatal(err)
	}
	//run.Interact(nil)
}
