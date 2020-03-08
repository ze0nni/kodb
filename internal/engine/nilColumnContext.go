package engine

func newNilColumnContext() ColumnContext {
	return &nilColumnContext{}
}

type nilColumnContext struct{}

func (c *nilColumnContext) HasRow(library, row string) (bool, error) {
	return false, nil
}
