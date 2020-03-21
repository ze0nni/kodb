package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/types"
)

func msgDeleteFieldFromJson(
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
	return &msgDeleteField{
		typeName: types.TypeName(typeName),
		fieldID:  types.FieldID(fieldID),
	}, nil
}

type msgDeleteField struct {
	typeName types.TypeName
	fieldID  types.FieldID
}

func (msg *msgDeleteField) Perform(srv *serverInstance) error {
	t, err := srv.engine.Types().Get(msg.typeName)
	if nil != err {
		return err
	}
	f, err := t.Get(msg.fieldID)
	if nil != err {
		return err
	}

	err = t.Delete(f)
	if nil != err {
		return err
	}

	return nil
}
