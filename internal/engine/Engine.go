package engine

import (
	"strings"

	"github.com/ze0nni/kodb/internal/driver"
)

func New(driver driver.Driver) Engine {
	e := &engine{
		driver:   driver,
		librarys: make(map[LibraryName]*libraryImp),
	}

	loadLibrarys(e)

	return e
}

type Engine interface {
	Librarys() []LibraryName
	GetLibrary(name LibraryName) Library
}

type engine struct {
	driver   driver.Driver
	librarys map[LibraryName]*libraryImp
}

func loadLibrarys(e *engine) {
	if ps, err := e.driver.Prefixes(); nil == err {
		for _, p := range ps {
			if strings.HasSuffix(p, "$schema") {
				libraryName := LibraryName(p[:len(p)-7])
				e.GetLibrary(libraryName)
			}
		}
	}
}

func (self *engine) Librarys() []LibraryName {
	out := []LibraryName{}

	for k, _ := range self.librarys {
		out = append(out, k)
	}

	return out
}

func (self *engine) GetLibrary(name LibraryName) Library {
	if storedLib := self.librarys[name]; nil != storedLib {
		return storedLib
	}
	newLib := newLibraryInst(
		name,
		LensOf(name.ToString()+"$schema", self.driver),
		LensOf(name.ToString()+"$data", self.driver),
		LensOf(name.ToString()+"$meta", self.driver),
	)
	self.librarys[name] = newLib
	return newLib
}
