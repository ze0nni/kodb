package main

import (
	"log"
	"strconv"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/run/web"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	eng := engine.New(driver.InMemory())
	userLib := eng.GetLibrary(engine.LibraryName("user"))
	firstname, _ := userLib.NewColumn("firstName")
	secondName, _ := userLib.NewColumn("secondName")
	age, _ := userLib.NewColumn("age")
	for i := 0; i < 10; i++ {
		row, _ := userLib.NewRow()
		userLib.UpdateValue(row, firstname, randomdata.FirstName(0))
		userLib.UpdateValue(row, secondName, randomdata.LastName())
		userLib.UpdateValue(row, age, strconv.Itoa(randomdata.Number(16, 40)))
	}

	invLib := eng.GetLibrary(engine.LibraryName("location"))
	name, _ := invLib.NewColumn("name")
	invLib.NewColumn("type")
	invLib.NewColumn("title")
	invLib.NewColumn("picture")
	//invLib.NewRefColumn("owner", userLib.Name())
	for i := 0; i < 20; i++ {
		row, _ := invLib.NewRow()
		invLib.UpdateValue(row, name, randomdata.Adjective())
	}

	err := web.Run(eng)
	if nil != err {
		log.Fatal(err)
	}
}
