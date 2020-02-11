package engine

import (
	"strings"

	"github.com/ze0nni/kodb/internal/driver"
)

func New(driver driver.Driver) Engine {
	e := &engine{
		driver:   driver,
		librarys: make(map[string]*libraryImp),
	}

	loadLibrarys(e)

	return e
}

type Engine interface {
	Librarys() []string
	GetLibrary(name string) Library
}

type engine struct {
	driver   driver.Driver
	librarys map[string]*libraryImp
}

func loadLibrarys(e *engine) {
	if ps, err := e.driver.Prefixes(); nil == err {
		for _, p := range ps {
			if strings.HasSuffix(p, "$schema") {
				e.GetLibrary(p[:len(p)-7])
			}
		}
	}
}

func (self *engine) Librarys() []string {
	out := []string{}

	for k, _ := range self.librarys {
		out = append(out, k)
	}

	return out
}

func (self *engine) GetLibrary(name string) Library {
	if storedLib := self.librarys[name]; nil != storedLib {
		return storedLib
	}
	newLib := newLibraryInst(
		name,
		LensOf(name+"$schema", self.driver),
		LensOf(name+"$data", self.driver),
		LensOf(name+"$meta", self.driver),
	)
	self.librarys[name] = newLib
	return newLib
}
