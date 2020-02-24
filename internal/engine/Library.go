package engine

import (
	"errors"
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

		Row(int) (Row, error)
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
	schema Lens,
	data Lens,
	meta Lens,
) *libraryImp {
	//TODO: error or panic
	if root, _ := schema.Get("root"); nil == root {
		schema.Put("root", make(entry.Entry))
	}
	return &libraryImp{
		name:   name,
		schema: schema,
		data:   data,
		meta:   meta,
		rows:   []RowID{},
	}
}

type libraryImp struct {
	name   LibraryName
	schema Lens
	data   Lens
	meta   Lens
	rows   []RowID
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

func (self *libraryImp) AddRow(rowID RowID) error {
	err := self.data.Put(rowID.ToString(), make(entry.Entry))
	if nil != err {
		return err
	}

	self.rows = append(self.rows, rowID)

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

func (lib *libraryImp) Row(index int) (Row, error) {
	if index < 0 || index >= len(lib.rows) {
		return nil, errors.New("Row index out of range")
	}

	id := lib.rows[index]
	return RowOf(
		lib.data,
		id,
	), nil
}
