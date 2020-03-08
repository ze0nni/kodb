package engine

import "fmt"

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
