package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/types"
)

func msgUpdateFieldFromJson(
	client ClientID,
	msg *simplejson.Json,
) (ClientMsg, error) {
	typeName, err := msg.Get("type").String()
	if nil != err {
		return nil, err
	}
	fieldID, err := msg.Get("field").String()
	if nil != err {
		return nil, err
	}
	name, err := msg.Get("name").String()
	if nil != err {
		return nil, err
	}
	return &msgUpdateField{
		typeName: types.TypeName(typeName),
		fieldID:  types.FieldID(fieldID),

		name: name,
	}, nil
}

type msgUpdateField struct {
	typeName types.TypeName
	fieldID  types.FieldID

	name string
}

func (msg *msgUpdateField) Perform(srv *serverInstance) error {
	t, err := srv.engine.Types().Get(msg.typeName)
	if nil != err {
		return err
	}

	f, err := t.Get(msg.fieldID)
	if nil != err {
		return err
	}

	f.Rename(msg.name)

	return nil
}
