package engine

import (
	"strconv"

	"github.com/ze0nni/kodb/internal/entry"
)

type (
	Library interface {
		Name() string

		Columns() int
		AddColumn(columnName string) error
	}
)

func newLibraryInst(
	name string,
	schema Lens,
	data Lens,
	meta Lens,
) *libraryImp {
	return &libraryImp{
		name:   name,
		schema: schema,
		data:   data,
		meta:   meta,
	}
}

type libraryImp struct {
	name   string
	schema Lens
	data   Lens
	meta   Lens
}

func (self *libraryImp) Name() string {
	return self.name
}

func (self *libraryImp) Columns() int {
	root, err := self.getSchemaRoot()
	if nil != err {
		return 0
	}
	count, err := strconv.Atoi(root["count"])
	if nil != err {
		return 0
	}
	return count
}

func (self *libraryImp) AddColumn(columnName string) error {
	return nil
}

func (self *libraryImp) getSchemaRoot() (entry.Entry, error) {
	root, err := self.schema.Get("root")
	if nil != err {
		return nil, err
	}
	if nil == err {
		return make(entry.Entry), nil
	}
	return root, nil
}
