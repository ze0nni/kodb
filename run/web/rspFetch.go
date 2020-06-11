package web

import (
	"log"

	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

func rspFetch(engine engine.Engine) (*simplejson.Json, error) {
	types := engine.Types()
	typesBody := simplejson.New()

TypesLoop:
	for _, name := range types.Names() {
		t, err := types.Get(name)
		if nil != err {
			log.Print(err)
			continue TypesLoop
		}
		body := simplejson.New()
		t.FillJson(body)
		typesBody.Set(name.String(), body)
	}

	librarysBody := simplejson.New()

LibrarysLoop:
	for _, name := range engine.Librarys() {
		library, err := engine.Library(name)
		if nil != err {
			log.Print(err)
			continue LibrarysLoop
		}
		libraryType, err := library.Type()
		if nil != err {
			log.Print(err)
			continue LibrarysLoop
		}
		librarysBody.Set(name.ToString(), libraryType.Name().String())
	}

	resp := simplejson.New()
	resp.Set("command", "fetch")
	resp.Set("types", typesBody)
	resp.Set("librarys", librarysBody)

	return resp, nil
}
