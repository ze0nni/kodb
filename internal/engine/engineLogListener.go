package engine

import "fmt"

func newLogListener() *logListener {
	return &logListener{
		log: []string{},
	}
}

type logListener struct {
	log []string
}

func (l *logListener) getLog() []string {
	return l.log
}

func (l *logListener) OnNewLibrary(name LibraryName) {
	l.log = append(l.log, "newLibrary "+name.ToString())
}

func (l *logListener) OnNewColumn(LibraryName, FieldID) {
	panic("not inplements")
}

func (l *logListener) OnNewRow(name LibraryName, row RowID) {
	l.log = append(l.log, fmt.Sprintf("newRow %s:%s", name.ToString(), row.ToString()))
}

func (l *logListener) OnDeleteRow(name LibraryName, row RowID) {
	l.log = append(l.log, fmt.Sprintf("deleteRow %s:%s", name.ToString(), row.ToString()))
}

func (l *logListener) OnUpdateValue(name LibraryName, row RowID, field FieldID, exists bool, value string, cellErr error) {
	if nil != cellErr {
		l.log = append(l.log, fmt.Sprintf("updateRow %s:%s:%s error %s",
			name.ToString(),
			row.ToString(),
			field.String(),
			cellErr,
		))
		return
	}

	l.log = append(l.log, fmt.Sprintf("updateRow %s:%s:%s %t %s",
		name.ToString(),
		row.ToString(),
		field.String(),
		exists,
		value,
	))
}

func (l *logListener) OnSwap(name LibraryName, i, j int, iID, jID RowID) {
	l.log = append(l.log, fmt.Sprintf("swap %s %d %d %s %s",
		name.ToString(),
		i,
		j,
		iID.ToString(),
		jID.ToString(),
	))
}
