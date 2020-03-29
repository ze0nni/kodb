package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

func msgSwapRowsFromJson(
	client ClientID,
	msg *simplejson.Json,
) (ClientMsg, error) {
	library, err := msg.Get("library").String()
	if nil != err {
		return nil, err
	}

	row, err := msg.Get("row").String()
	if nil != err {
		return nil, err
	}

	row0, err := msg.Get("row0").String()
	if nil != err {
		return nil, err
	}

	return &msgSwapRows{
		library: engine.LibraryName(library),
		row:     engine.RowID(row),
		row0:    engine.RowID(row0),
	}, nil
}

type msgSwapRows struct {
	library engine.LibraryName
	row     engine.RowID
	row0    engine.RowID
}

func (msg *msgSwapRows) Perform(srv *serverInstance) error {
	l, err := srv.engine.Library(msg.library)
	if nil != err {
		return err
	}
	return engine.SwapRowsByID(l, msg.row, msg.row0)
}
