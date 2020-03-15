package engine

func listenerNil() Listener {
	return &listenerNilInst{}
}

type listenerNilInst struct{}

func (lm *listenerNilInst) OnNewLibrary(name LibraryName) {

}

func (lm *listenerNilInst) OnNewColumn(LibraryName, ColumnID) {

}

func (lm *listenerNilInst) OnNewRow(name LibraryName, row RowID) {

}

func (lm *listenerNilInst) OnDeleteRow(LibraryName, RowID) {

}

func (lm *listenerNilInst) OnUpdateValue(LibraryName, RowID, ColumnID, bool, string, error) {

}
