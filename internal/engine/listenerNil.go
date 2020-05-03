package engine

func listenerNil() Listener {
	return &listenerNilInst{}
}

type listenerNilInst struct{}

func (lm *listenerNilInst) OnNewLibrary(name LibraryName) {

}

func (lm *listenerNilInst) OnNewColumn(LibraryName, FieldID) {

}

func (lm *listenerNilInst) OnNewRow(name LibraryName, row RowID) {

}

func (lm *listenerNilInst) OnDeleteRow(LibraryName, RowID) {

}

func (lm *listenerNilInst) OnUpdateValue(LibraryName, RowID, FieldID, bool, string, error) {

}

func (lm *listenerNilInst) OnSwap(LibraryName, int, int, RowID, RowID) {

}
