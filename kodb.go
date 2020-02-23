package main

import (
	"log"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/run/web"
)

func main() {
	eng := engine.New(driver.InMemory())
	userLib := eng.GetLibrary(engine.LibraryName("user"))
	userLib.NewColumn("firstname")
	userLib.NewColumn("secondName")
	userLib.NewColumn("age")

	invLib := eng.GetLibrary(engine.LibraryName("inventory"))
	invLib.NewColumn("name")
	invLib.NewColumn("type")
	invLib.NewColumn("title")
	invLib.NewColumn("picture")

	err := web.Run(eng)
	if nil != err {
		log.Fatal(err)
	}
}
