package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
)

type LibrarySchema struct {
	Name engine.LibraryName `json:"name"`
	Type types.TypeName     `json:"type"`
}

func NewLibrarySchemaFromEngine(
	name engine.LibraryName,
	engine engine.Engine,
) *LibrarySchema {
	l, err := engine.Library(name)
	if nil != err {
		panic(err)
	}

	var typeName types.TypeName
	if t, err := l.Type(); nil == err {
		typeName = t.Name()
	}

	schema := &LibrarySchema{
		Name: name,
		Type: typeName,
	}

	return schema
}
