package engine

import "github.com/ze0nni/kodb/internal/types"

func newNilColumnContext() ColumnContext {
	return &nilColumnContext{}
}

type nilColumnContext struct{}

func (c *nilColumnContext) HasRow(library, row string) (bool, error) {
	return false, nil
}

func (c *nilColumnContext) GetType(types.TypeName) (types.Type, error) {
	panic("No type")
}
