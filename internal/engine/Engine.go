package engine

import (
	"github.com/ze0nni/kodb/internal/driver"
)

func New(driver driver.Driver) Engine {
	e := engine{
		driver:   driver,
		librarys: make(map[string]*libraryImp),
	}

	return &e
}

type Engine interface {
	GetLibrary(name string) Library
}

type engine struct {
	driver   driver.Driver
	librarys map[string]*libraryImp
}

func (self *engine) GetLibrary(name string) Library {
	if storedLib := self.librarys[name]; nil != storedLib {
		return storedLib
	}
	newLib := newLibraryInst(name)
	self.librarys[name] = newLib
	return newLib
}
