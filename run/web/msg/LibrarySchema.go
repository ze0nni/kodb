package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
)

type LibrarySchema struct {
	Name    engine.LibraryName `json:"name"`
	Columns []ColumnSchema     `json:"columns"`
}

type ColumnSchema struct {
	Name string `json:"name"`
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
	columnName, _ := library.Column(index)

	schema := ColumnSchema{
		Name: columnName,
	}

	return schema
}
