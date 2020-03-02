package msg

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

type SetLibraryRowsMsg struct {
	Command string             `json:"command"`
	Library engine.LibraryName `json:"library"`
	Rows    []RowSchema        `json:"rows"`
}

type RowSchema struct {
	RowID engine.RowID     `json:"rowId"`
	Data  *simplejson.Json `json:"data"`
}

func SetLibraryRowsMsgFromEngine(
	name engine.LibraryName,
	eng engine.Engine,
) *SetLibraryRowsMsg {
	l := eng.GetLibrary(name)

	msg := &SetLibraryRowsMsg{
		Command: "setLibraryRows",
		Library: name,
		Rows:    []RowSchema{},
	}

	columns, err := engine.Columns(l)
	if nil != err {
		return msg
	}

	rows := l.Rows()
	for i := 0; i < rows; i++ {
		r, err := RowSchemaFromLibrary(i, columns, l)
		if nil == err {
			msg.Rows = append(msg.Rows, r)
		}
	}

	return msg
}

func RowSchemaFromLibrary(
	index int,
	columns []engine.ColumnData,
	library engine.Library,
) (RowSchema, error) {
	rowId, err := library.RowID(index)
	if nil != err {
		return RowSchema{}, err
	}

	row := RowSchema{
		RowID: rowId,
		Data:  simplejson.New(),
	}

	for _, col := range columns {
		colID := col.ID() // TODO: cache id's
		v, ok, err := library.GetValueAt(index, colID)

		colData := simplejson.New()
		colData.Set("exists", ok)
		if ok {
			colData.Set("value", v)
			cellErr := col.Validate(v)
			if nil != cellErr {
				colData.Set("error", cellErr)
			}
		}
		if nil != err {
			colData.Set("error", err)
		}

		row.Data.Set(colID.ToString(), colData)
	}

	return row, nil
}
