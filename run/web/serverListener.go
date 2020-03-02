package web

import (
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/run/web/msg"
)

type serverListener struct {
	server *serverInstance
}

func (l *serverListener) NewLibrary(engine.LibraryName) {
	panic("not implements")
}

func (l *serverListener) NewRow(name engine.LibraryName, row engine.RowID) {
	newRowMsg := msg.NewRowMsgOf(
		name,
		row,
	)

	for _, client := range l.server.clients {
		client.NewRow(newRowMsg)
	}
}

func (l *serverListener) DeleteRow(name engine.LibraryName, row engine.RowID) {
	deleteRowMsg := msg.DeleteRowMsgOf(
		name,
		row,
	)
	for _, client := range l.server.clients {
		client.DeleteRow(deleteRowMsg)
	}
}

func (l *serverListener) UpdateValue(
	name engine.LibraryName,
	row engine.RowID,
	col engine.ColumnID,
	exists bool,
	value string,
	cellErr error,
) {
	updateValueMsg := msg.UpdateValueMsgOf(
		name,
		row,
		col,
		exists,
		value,
		cellErr,
	)

	for _, client := range l.server.clients {
		client.UpdateValue(updateValueMsg)
	}
}
