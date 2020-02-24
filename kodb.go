package main

import (
	"log"
	"strconv"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/run/web"
)

func main() {
	eng := engine.New(driver.InMemory())
	userLib := eng.GetLibrary(engine.LibraryName("user"))
	firstname, _ := userLib.NewColumn("firstname")
	secondName, _ := userLib.NewColumn("secondName")
	age, _ := userLib.NewColumn("age")
	for i := 0; i < 5; i++ {
		userLib.NewRow()
		row, _ := userLib.Row(i)
		row.Set(firstname, "First name")
		row.Set(secondName, "Second name")
		row.Set(age, strconv.Itoa(i*5))
	}

	invLib := eng.GetLibrary(engine.LibraryName("inventory"))
	invLib.NewColumn("name")
	invLib.NewColumn("type")
	invLib.NewColumn("title")
	invLib.NewColumn("picture")
	for i := 0; i < 20; i++ {
		invLib.NewRow()
	}

	err := web.Run(eng)
	if nil != err {
		log.Fatal(err)
	}
}
