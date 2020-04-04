package engine

import (
	"fmt"

	"github.com/ze0nni/kodb/internal/types"
)

func newEngineColumnContext(eng *engine) *engineColumnContext {
	return &engineColumnContext{eng}
}

type engineColumnContext struct {
	eng *engine
}

func (c *engineColumnContext) HasRow(library, row string) (bool, error) {
	lib, exists := c.eng.librarys[LibraryName(library)]
	if false == exists {
		return false, fmt.Errorf("Library <%s> not exists", library)
	}
	return lib.HasRow(RowID(row)), nil
}

func (c *engineColumnContext) GetType(name types.TypeName) (types.Type, error) {
	return c.eng.types.Get(name)
}
