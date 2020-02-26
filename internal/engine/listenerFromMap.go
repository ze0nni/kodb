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
