package msg

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
)

type SetLibraryRowsMsg struct {
	Command string             `json:"command"`
	Library engine.LibraryName `json:"library"`
	Rows    []RowSchema        `json:"rows"`
}

type RowSchema struct {
	RowID     engine.RowID     `json:"rowId"`
	FieldCase types.FieldCase  `json:"case"`
	Data      *simplejson.Json `json:"data"`
}

func SetLibraryRowsMsgFromEngine(
	name engine.LibraryName,
	eng engine.Engine,
) *SetLibraryRowsMsg {
	l, err := eng.Library(name)
	if nil != err {
		panic(err)
	}

	msg := &SetLibraryRowsMsg{
		Command: "setLibraryRows",
		Library: name,
		Rows:    []RowSchema{},
	}

	tp, err := l.Type()
	if nil != err {
		return msg
	}
	fields := tp.Fields()

	rows := l.Rows()
	for i := 0; i < rows; i++ {
		r, err := RowSchemaFromLibrary(i, fields, eng.Context(), l)
		if nil == err {
			msg.Rows = append(msg.Rows, r)
		}
	}

	return msg
}

func RowSchemaFromLibrary(
	index int,
	fields []types.Field,
	context engine.ColumnContext,
	library engine.Library,
) (RowSchema, error) {
	rowId, err := library.RowID(index)
	if nil != err {
		return RowSchema{}, err
	}
	fieldCase, _ := library.Case(rowId)

	row := RowSchema{
		RowID:     rowId,
		FieldCase: fieldCase,
		Data:      simplejson.New(),
	}

	for _, field := range fields {
		fieldID := field.ID()
		v, ok, err := library.GetValueAt(index, fieldID)

		colData := simplejson.New()
		colData.Set("exists", ok)
		if ok {
			colData.Set("value", v)
			//TODO: field.Validate
			//cellErr := col.Validate(context, v)
			//if nil != cellErr {
			//	colData.Set("error", cellErr.Error())
			//}
		}
		if nil != err {
			colData.Set("error", err)
		}

		row.Data.Set(fieldID.String(), colData)
	}

	return row, nil
}
