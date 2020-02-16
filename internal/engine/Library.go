package engine

import (
	"errors"
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/ze0nni/kodb/internal/entry"
)

type (
	LibraryName string

	ColumnId string
	RowId string

	Library interface {
		Name() LibraryName

		Columns() int
		NewColumn(columnName string) (ColumnId, error)
		AddColumn(id ColumnId, columnName string) error
		Column(index int) (string, error)

		Rows() int
		NewRow() (RowId, error)
		AddRow(RowId) error
		HasRow(RowId) bool
	}
)

func (name LibraryName) ToString() string {
	return string(name)
}

func (columnId ColumnId) ToString() string {
	return string(columnId)
}

func newLibraryInst(
	name LibraryName,
	schema Lens,
	data Lens,
	meta Lens,
) *libraryImp {
	schema.Put("root", make(entry.Entry))
	return &libraryImp{
		name:   name,
		schema: schema,
		data:   data,
		meta:   meta,
		rows:   []RowId{},
	}
}

type libraryImp struct {
	name   LibraryName
	schema Lens
	data   Lens
	meta   Lens
	rows   []RowId
}

func (self *libraryImp) Name() LibraryName {
	return self.name
}

func (self *libraryImp) Columns() int {
	root, err := self.getSchemaRoot()
	if nil != err {
		return 0
	}
	return entry.IntDef("columns", 0, root)
}

func (self *libraryImp) NewColumn(columnName string) (ColumnId, error) {
	columnV4, err := uuid.NewV4()
	if nil != err {
		return ColumnId(""), err
	}

	columnId := ColumnId(columnV4.String())
	if err := self.AddColumn(columnId, columnName); nil != err {
		return ColumnId(""), err
	}

	return columnId, nil
}

func (lib *libraryImp) AddColumn(id ColumnId, name string) error {
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

func (self *libraryImp) Column(index int) (string, error) {
	root, err := self.getSchemaRoot()
	if nil != err {
		return "", err
	}
	if columnIdentity, ok := root["column_"+strconv.Itoa(index)]; ok {
		columnEntry, err := self.schema.Get(columnIdentity)
		if nil != err {
			return "", err
		}
		return columnEntry["name"], nil
	}

	return "", errors.New("not found")
}

func (self *libraryImp) getSchemaRoot() (entry.Entry, error) {
	root, err := self.schema.Get("root")
	if nil != err {
		return nil, err
	}
	if nil == root {
		return make(entry.Entry), nil
	}
	return root, nil
}

func (self *libraryImp) Rows() int {
	return len(self.rows)
}

func (self *libraryImp) NewRow() (RowId, error) {
	rowV4, err := uuid.NewV4()
	if nil != err {
		return RowId(""), err
	}
	rowId := RowId(rowV4.String())

	err = self.AddRow(rowId)
	if nil != err {
		return RowId(""), err
	}

	return rowId, nil
}

func (self *libraryImp) AddRow(rowId RowId) error {

	self.rows = append(self.rows, rowId)

	return nil
}

func (self *libraryImp) HasRow(rowId RowId) bool {
	for _, r := range self.rows {
		if r == rowId {
			return true
		}
	}
	return false
}
