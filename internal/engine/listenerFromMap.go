package engine

func newListenerFromMap() *listenerFromMap {
	return &listenerFromMap{
		listeners:        make(map[Listener]struct{}),
		libraryListeners: make(map[LibraryName]map[Listener]struct{}),
	}
}

type listenerFromMap struct {
	listeners        map[Listener]struct{}
	libraryListeners map[LibraryName]map[Listener]struct{}
}

func (e *listenerFromMap) listen(listener Listener) func() {
	e.listeners[listener] = struct{}{}
	return func() {
		delete(e.listeners, listener)
	}
}

func (e *listenerFromMap) listenLibrary(library LibraryName, listener Listener) func() {
	libraryListenersMap, exists := e.libraryListeners[library]
	if false == exists {
		libraryListenersMap = make(map[Listener]struct{})
		e.libraryListeners[library] = libraryListenersMap
	}
	libraryListenersMap[listener] = struct{}{}

	return func() {
		delete(libraryListenersMap, listener)
	}
}

func (e *listenerFromMap) forLibrary(
	library LibraryName,
	consumer func(Listener),
) {
	if libraryListenersMap, ok := e.libraryListeners[library]; ok {
		for l := range libraryListenersMap {
			consumer(l)
		}
	}
}

func (lm *listenerFromMap) NewLibrary(name LibraryName) {
	for l, _ := range lm.listeners {
		l.NewLibrary(name)
	}
	lm.forLibrary(name, func(l Listener) {
		l.NewLibrary(name)
	})
}

func (lm *listenerFromMap) NewRow(name LibraryName, row RowID) {
	for l, _ := range lm.listeners {
		l.NewRow(name, row)
	}
	lm.forLibrary(name, func(l Listener) {
		l.NewRow(name, row)
	})
}

func (lm *listenerFromMap) DeleteRow(name LibraryName, row RowID) {
	for l, _ := range lm.listeners {
		l.DeleteRow(name, row)
	}
	lm.forLibrary(name, func(l Listener) {
		l.DeleteRow(name, row)
	})
}

func (lm *listenerFromMap) UpdateValue(name LibraryName, row RowID, col ColumnID, exixts bool, value string, cellErr error) {
	for l, _ := range lm.listeners {
		l.UpdateValue(name, row, col, exixts, value, cellErr)
	}
	lm.forLibrary(name, func(l Listener) {
		l.UpdateValue(name, row, col, exixts, value, cellErr)
	})
}
