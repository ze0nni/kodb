package engine

import (
	"errors"
	"fmt"
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/ze0nni/kodb/internal/entry"
)

type (
	// LibraryName type
	LibraryName string

	// ColumnID type
	ColumnID string

	// Library type
	Library interface {
		Name() LibraryName

		Columns() int
		NewColumn(columnName string) (ColumnID, error)
		AddColumn(id ColumnID, columnName string) error
		Column(index int) (ColumnID, error)
		ColumnName(index int) (string, error)

		Rows() int
		NewRow() (RowID, error)
		AddRow(RowID) error
		HasRow(RowID) bool
		DeleteRow(RowID) error

		Row(int) (Row, error)
		RowID(int) (RowID, error)

		GetRowColumn(int, ColumnID) (string, bool, error)
		GetValue(RowID, ColumnID) (string, bool, error)
		UpdateValue(RowID, ColumnID, string) error
	}
)

func (name LibraryName) ToString() string {
	return string(name)
}

func (id ColumnID) ToString() string {
	return string(id)
}

func newLibraryInst(
	name LibraryName,
	listener Listener,
	schema Lens,
	data Lens,
	meta Lens,
) *libraryImp {
	//TODO: error or panic
	if root, _ := schema.Get("root"); nil == root {
		schema.Put("root", make(entry.Entry))
	}
	return &libraryImp{
		name:     name,
		listener: listener,
		schema:   schema,
		data:     data,
		meta:     meta,
		rows:     []RowID{},
	}
}

// ColumnIDs return slice for ColumnID
func ColumnIDs(library Library) ([]ColumnID, error) {
	columns := library.Columns()
	out := make([]ColumnID, columns)
	for i := 0; i < columns; i++ {
		c, err := library.Column(i)
		if nil != err {
			return nil, err
		}
		out = append(out, c)
	}
	return out, nil
}

type libraryImp struct {
	name     LibraryName
	listener Listener
	schema   Lens
	data     Lens
	meta     Lens
	rows     []RowID
}

func (lib *libraryImp) Name() LibraryName {
	return lib.name
}

func (lib *libraryImp) Columns() int {
	root, err := lib.getSchemaRoot()
	if nil != err {
		return 0
	}
	return entry.IntDef("columns", 0, root)
}

func (lib *libraryImp) NewColumn(columnName string) (ColumnID, error) {
	columnV4, err := uuid.NewV4()
	if nil != err {
		return ColumnID(""), err
	}

	columnID := ColumnID(columnV4.String())
	if err := lib.AddColumn(columnID, columnName); nil != err {
		return ColumnID(""), err
	}

	return columnID, nil
}

func (lib *libraryImp) AddColumn(id ColumnID, name string) error {
	root, err := lib.getSchemaRoot()
	if nil != err {
		return err
	}

	if s, _ := lib.schema.Get(id.ToString()); nil != s {
		return errors.New("duplicate column " + id.ToString())
	}

	num := entry.IntDef("columns", 0, root)
	entry.SetInt("columns", num+1, root)

	entry.SetString("column_"+strconv.Itoa(num), id.ToString(), root)

	columnEntry := make(entry.Entry)
	entry.SetString("name", name, columnEntry)

	lib.schema.Put(id.ToString(), columnEntry)
	lib.schema.Put("root", root)

	return nil
}

func (lib *libraryImp) Column(index int) (ColumnID, error) {
	root, err := lib.getSchemaRoot()
	if nil != err {
		return ColumnID(""), err
	}
	if columnIdentity, ok := root["column_"+strconv.Itoa(index)]; ok {
		return ColumnID(columnIdentity), nil
	}

	return ColumnID(""), errors.New("not found")
}

func (lib *libraryImp) ColumnName(index int) (string, error) {
	root, err := lib.getSchemaRoot()
	if nil != err {
		return "", err
	}
	if columnIdentity, ok := root["column_"+strconv.Itoa(index)]; ok {
		columnEntry, err := lib.schema.Get(columnIdentity)
		if nil != err {
			return "", err
		}
		return columnEntry["name"], nil
	}

	return "", errors.New("not found")
}

func (lib *libraryImp) getSchemaRoot() (entry.Entry, error) {
	root, err := lib.schema.Get("root")
	if nil != err {
		return nil, err
	}
	if nil == root {
		return make(entry.Entry), nil
	}
	return root, nil
}

func (lib *libraryImp) Rows() int {
	return len(lib.rows)
}

func (lib *libraryImp) NewRow() (RowID, error) {
	rowV4, err := uuid.NewV4()
	if nil != err {
		return RowID(""), err
	}
	rowId := RowID(rowV4.String())

	err = lib.AddRow(rowId)
	if nil != err {
		return RowID(""), err
	}

	return rowId, nil
}

func (l *libraryImp) AddRow(rowID RowID) error {
	err := l.data.Put(rowID.ToString(), make(entry.Entry))
	if nil != err {
		return err
	}

	l.rows = append(l.rows, rowID)

	l.listener.NewRow(l.name, rowID)

	return nil
}

func (self *libraryImp) HasRow(rowID RowID) bool {
	for _, r := range self.rows {
		if r == rowID {
			return true
		}
	}
	return false
}

func (l *libraryImp) DeleteRow(id RowID) error {
	e, err := l.data.Get(id.ToString())
	if nil != err {
		return err
	}
	if nil == e {
		return fmt.Errorf("Row not exists: %s", id.ToString())
	}
	err = l.data.Delete(id.ToString())
	if nil != err {
		return err
	}

	rows := l.rows
	rowsCount := len(rows)
	for i := 0; i < rowsCount; i++ {
		if id == rows[i] {
			l.rows = append(rows[:i], rows[i+1:]...)
			break
		}
	}

	l.listener.DeleteRow(l.name, id)

	return nil
}

func (lib *libraryImp) Row(index int) (Row, error) {
	id, err := lib.RowID(index)
	if nil != err {
		return nil, err
	}
	return RowOf(
		lib.data,
		id,
	), nil
}

func (lib *libraryImp) RowID(index int) (RowID, error) {
	if index < 0 || index >= len(lib.rows) {
		return RowID(""), errors.New("Row index out of range")
	}

	return lib.rows[index], nil
}

func (lib *libraryImp) GetRowColumn(
	index int,
	column ColumnID,
) (string, bool, error) {
	id, err := lib.RowID(index)
	if nil != err {
		return "", false, err
	}
	e, err := lib.data.Get(id.ToString())
	if nil != err {
		return "", false, err
	}
	if nil == e {
		return "", false, errors.New("Row not exists")
	}
	v, ok := e[column.ToString()]
	return v, ok, nil
}

func (lib *libraryImp) GetValue(
	id RowID,
	col ColumnID,
) (string, bool, error) {
	e, err := lib.data.Get(id.ToString())
	if nil != err {
		return "", false, err
	}
	if nil == e {
		return "", false, errors.New("Row not exists")
	}

	v, exists := e[col.ToString()]
	if false == exists {
		return "", false, nil
	}
	return v, true, nil
}

func (lib *libraryImp) UpdateValue(
	id RowID,
	col ColumnID,
	value string,
) error {
	e, err := lib.data.Get(id.ToString())
	if nil != err {
		return err
	}
	if nil == e {
		return errors.New("Row not exists")
	}
	e[col.ToString()] = value

	err = lib.data.Put(id.ToString(), e)

	if nil != err {
		return err
	}

	lib.listener.UpdateValue(lib.name, id, col, true, value)

	return nil
}
