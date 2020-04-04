package engine

import (
	"github.com/ze0nni/kodb/internal/types"
)

// ColumnContext type
type ColumnContext interface {
	GetType(types.TypeName) (types.Type, error)
	HasRow(library, row string) (bool, error)
	//GetValue(library, row, col string) (string, bool, error)
}
