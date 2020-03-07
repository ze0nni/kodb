package msg

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

type LibrarySchema struct {
	Name    engine.LibraryName `json:"name"`
	Columns []*simplejson.Json `json:"columns"`
}

func NewLibrarySchemaFromEngine(
	name engine.LibraryName,
	engine engine.Engine,
) *LibrarySchema {
	l := engine.GetLibrary(name)

	schema := &LibrarySchema{
		Name:    name,
		Columns: []*simplejson.Json{},
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

func NewColumnSchemaFromLibrary(index int, library engine.Library) *simplejson.Json {
	schema := simplejson.New()

	col, err := library.ColumnData(index)
	if nil == err {
		col.FillJson(schema)
	}

	return schema
}
