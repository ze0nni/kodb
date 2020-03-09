package engine

func newListenerFromMap() *listenerFromMap {
	return &listenerFromMap{
		listeners: make(map[Listener]struct{}),
	}
}

type listenerFromMap struct {
	listeners map[Listener]struct{}
}

func (e *listenerFromMap) listen(listener Listener) func() {
	e.listeners[listener] = struct{}{}
	return func() {
		delete(e.listeners, listener)
	}
}

func (lm *listenerFromMap) NewLibrary(name LibraryName) {
	for l, _ := range lm.listeners {
		l.NewLibrary(name)
	}
}

func (lm *listenerFromMap) NewRow(name LibraryName, row RowID) {
	for l, _ := range lm.listeners {
		l.NewRow(name, row)
	}
}

func (lm *listenerFromMap) DeleteRow(name LibraryName, row RowID) {
	for l, _ := range lm.listeners {
		l.DeleteRow(name, row)
	}
}

func (lm *listenerFromMap) UpdateValue(name LibraryName, row RowID, col ColumnID, exixts bool, value string, cellErr error) {
	for l, _ := range lm.listeners {
		l.UpdateValue(name, row, col, exixts, value, cellErr)
	}
}
