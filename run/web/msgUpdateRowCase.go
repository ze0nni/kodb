package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
)

func msgUpdateRowCaseFromJson(
	client ClientID,
	msg *simplejson.Json,
) (ClientMsg, error) {
	library, err := msg.Get("library").String()
	if nil != err {
		return nil, err
	}
	rowID, err := msg.Get("rowId").String()
	if nil != err {
		return nil, err
	}
	fieldCase, err := msg.Get("case").String()
	if nil != err {
		return nil, err
	}

	return &msgUpdateRowCase{
		library:   engine.LibraryName(library),
		rowID:     engine.RowID(rowID),
		fieldCase: types.FieldCase(fieldCase),
	}, nil
}

type msgUpdateRowCase struct {
	library   engine.LibraryName
	rowID     engine.RowID
	fieldCase types.FieldCase
}

func (msg *msgUpdateRowCase) Perform(srv *serverInstance) error {
	l, err := srv.engine.Library(msg.library)
	if nil != err {
		return err
	}
	err = l.UpdateCase(msg.rowID, msg.fieldCase)
	if nil != err {
		return err
	}

	return nil
}
