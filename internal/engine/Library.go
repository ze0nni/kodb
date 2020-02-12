package engine

import (
	"errors"
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/ze0nni/kodb/internal/entry"
)

type (
	LibraryName string

	Library interface {
		Name() LibraryName

		Columns() int
		AddColumn(columnName string) error
		Column(index int) (string, error)
	}
)

func (name LibraryName) ToString() string {
	return string(name)
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
	}
}

type libraryImp struct {
	name   LibraryName
	schema Lens
	data   Lens
	meta   Lens
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

func (self *libraryImp) AddColumn(columnName string) error {
	root, err := self.getSchemaRoot()
	if nil != err {
		return err
	}
	num := entry.IntDef("columns", 0, root)
	entry.SetInt("columns", num+1, root)

	columnV4, err := uuid.NewV4()
	if nil != err {
		return err
	}

	columnIdentity := columnV4.String()
	entry.SetString("column_"+strconv.Itoa(num), columnIdentity, root)

	columnEntry := make(entry.Entry)
	entry.SetString("name", columnName, columnEntry)

	self.schema.Put(columnIdentity, columnEntry)
	self.schema.Put("root", root)

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
