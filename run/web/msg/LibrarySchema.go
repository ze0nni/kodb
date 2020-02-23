package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
)

type LibrarySchema struct {
	Name    engine.LibraryName `json:"name"`
	Columns []ColumnSchema     `json:"columns"`
}

type ColumnSchema struct {
	ID   engine.ColumnID `json:"id"`
	Name string          `json:"name"`
}

func NewLibrarySchemaFromEngine(
	name engine.LibraryName,
	engine engine.Engine,
) *LibrarySchema {
	l := engine.GetLibrary(name)

	schema := &LibrarySchema{
		Name:    name,
		Columns: []ColumnSchema{},
	}

	columnds := l.Columns()
	for i := 0; i < columnds; i++ {
		schema.Columns = append(
			schema.Columns,
			NewColumnSchemaFromLibrary(i, l),
		)
	}

	return schema
}

func NewColumnSchemaFromLibrary(index int, library engine.Library) ColumnSchema {
	columnID, _ := library.Column(index)
	columnName, _ := library.ColumnName(index)

	schema := ColumnSchema{
		ID:   columnID,
		Name: columnName,
	}

	return schema
}
