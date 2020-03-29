package engine

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/entry"
)

type (
	// LibraryName type
	LibraryName string

	// ColumnID type
	ColumnID string

	// RowID type
	RowID string

	// Library type
	Library interface {
		Name() LibraryName

		Columns() int

		NewColumn(ColumnData) (ColumnData, error)
		AddColumn(ColumnID, ColumnData) (ColumnData, error)

		Column(index int) (ColumnID, error)
		ColumnName(index int) (string, error)
		ColumnData(int) (ColumnData, error)
		ColumnDataOf(id ColumnID) (ColumnData, error)
		UpdateColumnData(data ColumnData) (ColumnData, error)

		Rows() int
		NewRow() (RowID, error)
		AddRow(RowID) error
		HasRow(RowID) bool
		DeleteRow(RowID) error

		RowID(int) (RowID, error)

		Swap(int, int) error

		GetValueAt(int, ColumnID) (string, bool, error)
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

func (id RowID) ToString() string {
	return string(id)
}

func newLibraryInst(
	name LibraryName,
	context ColumnContext,
	listener Listener,
	schema driver.Lens,
	data driver.Lens,
	meta driver.Lens,
) *libraryImp {
	//TODO: error or panic
	if root, _ := schema.Get("root"); nil == root {
		schema.Put("root", make(entry.Entry))
	}

	rowKeys, err := data.Keys()
	if nil != err {
		panic(err)
	}

	rows := []RowID{}
	for _, id := range rowKeys {
		rows = append(rows, RowID(id))
	}
	sort.Sort(&rowsByOrder{data, rows})

	lib := &libraryImp{
		name:     name,
		context:  context,
		listener: listener,
		schema:   schema,
		data:     data,
		meta:     meta,
		rows:     rows,
	}

	return lib
}

// Columns return slice for ColumnData
func Columns(library Library) ([]ColumnData, error) {
	columns := library.Columns()
	out := make([]ColumnData, columns)
	for i := 0; i < columns; i++ {
		c, err := library.ColumnData(i)
		if nil != err {
			return nil, err
		}
		out = append(out, c)
	}
	return out, nil
}

type libraryImp struct {
	name     LibraryName
	context  ColumnContext
	listener Listener
	schema   driver.Lens
	data     driver.Lens
	meta     driver.Lens
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

func (lib *libraryImp) NewColumn(data ColumnData) (ColumnData, error) {
	columnV4, err := uuid.NewV4()
	if nil != err {
		return ColumnData{nil}, err
	}
	id := ColumnID(columnV4.String())
	return lib.AddColumn(id, data)
}

func (lib *libraryImp) AddColumn(id ColumnID, data ColumnData) (ColumnData, error) {
	root, err := lib.getSchemaRoot()
	if nil != err {
		return ColumnData{nil}, err
	}

	if s, _ := lib.schema.Get(id.ToString()); nil != s {
		return ColumnData{nil}, errors.New("duplicate column " + id.ToString())
	}

	num := entry.IntDef("columns", 0, root)
	entry.SetInt("columns", num+1, root)

	entry.SetString("column_"+strconv.Itoa(num), id.ToString(), root)

	data = data.NewID(id)

	lib.schema.Put(id.ToString(), data.entry)
	lib.schema.Put("root", root)

	lib.listener.OnNewColumn(lib.Name(), id)

	return data, nil
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
	data, err := lib.ColumnData(index)
	if nil != err {
		return "", err
	}
	return data.Name(), nil
}

func (lib *libraryImp) ColumnData(index int) (ColumnData, error) {
	root, err := lib.getSchemaRoot()
	if nil != err {
		return ColumnData{nil}, err
	}
	if columnIdentity, ok := root["column_"+strconv.Itoa(index)]; ok {
		return lib.ColumnDataOf(ColumnID(columnIdentity))
	}

	return ColumnData{nil}, fmt.Errorf("Column <%d> not exists", index)
}

func (lib *libraryImp) ColumnDataOf(id ColumnID) (ColumnData, error) {
	columnEntry, err := lib.schema.Get(id.ToString())
	if nil != err {
		return ColumnData{nil}, err
	}
	return ColumnData{columnEntry}, nil
}

func (lib *libraryImp) UpdateColumnData(data ColumnData) (ColumnData, error) {
	columnEntry, err := lib.schema.Get(data.ID().ToString())
	if nil != err {
		return ColumnData{nil}, err
	}
	if nil == columnEntry {
		return ColumnData{nil}, fmt.Errorf("column %s not exists", data.Name())
	}

	lib.schema.Put(data.ID().ToString(), data.entry)

	return lib.ColumnDataOf(data.ID())
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
	e := make(entry.Entry)
	e["order"] = strconv.Itoa(len(l.rows))

	err := l.data.Put(rowID.ToString(), e)
	if nil != err {
		return err
	}

	l.rows = append(l.rows, rowID)

	l.listener.OnNewRow(l.name, rowID)

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

	l.listener.OnDeleteRow(l.name, id)

	return nil
}

func (lib *libraryImp) RowID(index int) (RowID, error) {
	if index < 0 || index >= len(lib.rows) {
		return RowID(""), errors.New("Row index out of range")
	}

	return lib.rows[index], nil
}

func (lib *libraryImp) Swap(i, j int) error {
	if 0 == len(lib.rows) {
		return errors.New("Library is empty")
	}

	rowI, err := lib.RowID(i)
	if nil != err {
		return err
	}

	rowJ, err := lib.RowID(j)
	if nil != err {
		return err
	}

	// Todo: transaction!
	err = driver.UpdateLens(lib.data, rowI.ToString(), func(e entry.Entry) {
		e["order"] = strconv.Itoa(j)
	})
	if nil != err {
		return err
	}

	err = driver.UpdateLens(lib.data, rowJ.ToString(), func(e entry.Entry) {
		e["order"] = strconv.Itoa(i)
	})
	if nil != err {
		return err
	}

	lib.rows[i] = rowJ
	lib.rows[j] = rowI

	return nil
}

func (lib *libraryImp) GetValueAt(
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

	colData, err := lib.ColumnDataOf(col)
	if nil != err {
		return err
	}

	e[col.ToString()] = value

	err = lib.data.Put(id.ToString(), e)

	if nil != err {
		return err
	}

	cellErr := colData.Validate(lib.context, value)

	lib.listener.OnUpdateValue(lib.name, id, col, true, value, cellErr)

	return nil
}
