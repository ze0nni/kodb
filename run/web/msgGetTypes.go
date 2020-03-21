package web

import (
	"github.com/bitly/go-simplejson"
)

type msgGetTypes struct {
	clientID ClientID
}

func (m *msgGetTypes) Perform(srv *serverInstance) error {
	client := srv.clients[m.clientID]
	if nil == client {
		return nil
	}

	types := simplejson.New()

TypesLoop:
	for _, name := range srv.engine.Types().Names() {
		t, err := srv.engine.Types().Get(name)
		if nil != err {
			continue TypesLoop
		}
		body := simplejson.New()
		t.FillJson(body)
		types.Set(name.String(), body)
	}

	resp := simplejson.New()
	resp.Set("command", "setTypes")
	resp.Set("types", types)

	client.Send(resp)

	return nil
}
