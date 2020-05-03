package web

import (
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
	"github.com/ze0nni/kodb/run/web/msg"
)

type serverListener struct {
	server *serverInstance
}

func (l *serverListener) OnNewLibrary(engine.LibraryName) {
	// TODO: send only new schema
	setSchemaMsg := msg.SetSchemaMsgFromEngine(l.server.engine)

	for _, client := range l.server.clients {
		client.SetSchema(setSchemaMsg)
	}
}

func (l *serverListener) OnNewColumn(libraryName engine.LibraryName, field engine.FieldID) {
	setLibraryRowsMsg := msg.SetLibraryRowsMsgFromEngine(libraryName, l.server.engine)
	for _, client := range l.server.clients {
		client.SetLibraryRows(setLibraryRowsMsg)
	}

	// TODO: send only new schema
	setSchemaMsg := msg.SetSchemaMsgFromEngine(l.server.engine)

	for _, client := range l.server.clients {
		client.SetSchema(setSchemaMsg)
	}
}

func (l *serverListener) OnNewRow(name engine.LibraryName, row engine.RowID) {
	newRowMsg := msg.NewRowMsgOf(
		name,
		row,
	)

	for _, client := range l.server.clients {
		client.NewRow(newRowMsg)
	}
}

func (l *serverListener) OnDeleteRow(name engine.LibraryName, row engine.RowID) {
	deleteRowMsg := msg.DeleteRowMsgOf(
		name,
		row,
	)
	for _, client := range l.server.clients {
		client.DeleteRow(deleteRowMsg)
	}
}

func (l *serverListener) OnUpdateValue(
	name engine.LibraryName,
	row engine.RowID,
	col engine.FieldID,
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

func (l *serverListener) OnSwap(
	name engine.LibraryName,
	i, j int,
	rowI, rowJ engine.RowID,
) {
	rsp := rspSwapRows(name, i, j, rowI, rowJ)
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}

func (l *serverListener) OnNewType(types.TypeName) {
	//TODO newType
	rsp, err := rspSetTypes(l.server.engine.Types())
	if nil != err {
		return
	}
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}

func (l *serverListener) OnDeleteType(types.TypeName) {
	//TODO deleteType
	rsp, err := rspSetTypes(l.server.engine.Types())
	if nil != err {
		return
	}
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}

func (l *serverListener) OnChangedType(types.TypeName) {
	//TODO changedType
	rsp, err := rspSetTypes(l.server.engine.Types())
	if nil != err {
		return
	}
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}

func (l *serverListener) OnNewField(types.TypeName, types.FieldID) {
	//TODO newField
	rsp, err := rspSetTypes(l.server.engine.Types())
	if nil != err {
		return
	}
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}

func (l *serverListener) OnDeleteField(types.TypeName, types.FieldID) {
	//TODO deleteField
	rsp, err := rspSetTypes(l.server.engine.Types())
	if nil != err {
		return
	}
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}

func (l *serverListener) OnChangedField(types.TypeName, types.FieldID) {
	//TODO changedField
	rsp, err := rspSetTypes(l.server.engine.Types())
	if nil != err {
		return
	}
	for _, client := range l.server.clients {
		client.Send(rsp)
	}
}
