package engine

func listenerNil() Listener {
	return &listenerNilInst{}
}

type listenerNilInst struct{}

func (lm *listenerNilInst) NewLibrary(name LibraryName) {

}

func (lm *listenerNilInst) NewRow(name LibraryName, row RowID) {

}