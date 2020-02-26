package engine

import (
	"strings"

	"github.com/ze0nni/kodb/internal/driver"
)

func New(driver driver.Driver) Engine {
	e := &engine{
		driver:    driver,
		librarys:  make(map[LibraryName]*libraryImp),
		listeners: make(map[Listener]struct{}),
	}

	loadLibrarys(e)

	return e
}

type Engine interface {
	Librarys() []LibraryName
	GetLibrary(name LibraryName) Library
	Listen(listener Listener) func()
}

type Listener interface {
	NewLibrary(LibraryName)
	NewRow(LibraryName, RowID)
}

type engine struct {
	driver    driver.Driver
	librarys  map[LibraryName]*libraryImp
	listeners map[Listener]struct{}
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

func (e *engine) GetLibrary(name LibraryName) Library {
	if storedLib := e.librarys[name]; nil != storedLib {
		return storedLib
	}
	newLib := newLibraryInst(
		name,
		newListenerFromMap(e.listeners),
		LensOf(name.ToString()+"$schema", e.driver),
		LensOf(name.ToString()+"$data", e.driver),
		LensOf(name.ToString()+"$meta", e.driver),
	)
	e.librarys[name] = newLib

	for l, _ := range e.listeners {
		l.NewLibrary(name)
	}

	return newLib
}

func (e *engine) Listen(listener Listener) func() {
	e.listeners[listener] = struct{}{}
	return func() {
		delete(e.listeners, listener)
	}
}
