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

func (tl *typesListener) OnNewType(name TypeName) {
	for l, _ := range tl.listeners {
		l.OnNewType(name)
	}
}

func (tl *typesListener) OnDeleteType(name TypeName) {
	for l, _ := range tl.listeners {
		l.OnDeleteType(name)
	}
}

func (tl *typesListener) OnChangedType(name TypeName) {
	for l, _ := range tl.listeners {
		l.OnChangedType(name)
	}
}

func (tl *typesListener) OnNewField(name TypeName, id FieldID) {
	for l, _ := range tl.listeners {
		l.OnNewField(name, id)
	}
}

func (tl *typesListener) OnDeleteField(name TypeName, id FieldID) {
	for l, _ := range tl.listeners {
		l.OnDeleteField(name, id)
	}
}

func (tl *typesListener) OnChangedField(name TypeName, id FieldID) {
	for l, _ := range tl.listeners {
		l.OnChangedField(name, id)
	}
}
