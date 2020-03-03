package engine

func newEngineColumnContext(engine Engine) ColumnContext {
	return &engineColumnContext{}
}

type engineColumnContext struct {
	engine Engine
}
