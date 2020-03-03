package engine

func newNilColumnContext() ColumnContext {
	return &nilColumnContext{}
}

type nilColumnContext struct{}
