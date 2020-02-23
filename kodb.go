package main

import (
	"log"

	"github.com/ze0nni/kodb/run/web"
)

func main() {
	err := web.Run()
	if nil != err {
		log.Fatal(err)
	}
}
