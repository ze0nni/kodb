package types

func newTypesListener() *typesListener {
	return &typesListener{
		listeners: make(map[TypesListener]struct{}),
	}
}

type typesListener struct {
	listeners map[TypesListener]struct{}
}

func (tl *typesListener) Listen(l TypesListener) func() {
	tl.listeners[l] = struct{}{}
	return func() {
		delete(tl.listeners, l)
	}
}
