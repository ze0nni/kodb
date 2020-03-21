package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/types"
)

func rspSetTypes(types types.Types) (*simplejson.Json, error) {
	typesBody := simplejson.New()

TypesLoop:
	for _, name := range types.Names() {
		t, err := types.Get(name)
		if nil != err {
			continue TypesLoop
		}
		body := simplejson.New()
		t.FillJson(body)
		typesBody.Set(name.String(), body)
	}

	resp := simplejson.New()
	resp.Set("command", "setTypes")
	resp.Set("types", typesBody)

	return resp, nil
}
