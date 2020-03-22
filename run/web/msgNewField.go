package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/types"
)

func msgNewFieldFromJson(
	client ClientID,
	msg *simplejson.Json,
) (ClientMsg, error) {
	typeName, err := msg.Get("type").String()
	if nil != err {
		return nil, err
	}
	fieldCase, err := msg.Get("case").String()
	if nil != err {
		return nil, err
	}
	return &msgNewField{
		typeName:  types.TypeName(typeName),
		fieldCase: types.FieldCase(fieldCase),
	}, nil
}

type msgNewField struct {
	typeName  types.TypeName
	fieldCase types.FieldCase
}

func (msg *msgNewField) Perform(srv *serverInstance) error {
	t, err := srv.engine.Types().Get(msg.typeName)
	if nil != err {
		return err
	}

	newType := types.NewValueFieldData("newField")
	newType.SetCase(msg.fieldCase)

	_, err = t.New(newType)
	if nil != err {
		return err
	}

	return nil
}
