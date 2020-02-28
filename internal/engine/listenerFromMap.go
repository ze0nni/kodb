package engine

func newListenerFromMap(listeners map[Listener]struct{}) Listener {
	return &listenerFromMap{
		listeners: listeners,
	}
}

type listenerFromMap struct {
	listeners map[Listener]struct{}
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

func (lm *listenerFromMap) UpdateValue(name LibraryName, row RowID, col ColumnID, exixts bool, value string) {
	for l, _ := range lm.listeners {
		l.UpdateValue(name, row, col, exixts, value)
	}
}
