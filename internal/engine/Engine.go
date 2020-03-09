package engine

import (
	"strings"

	"github.com/ze0nni/kodb/internal/driver"
)

func New(driver driver.Driver) Engine {
	e := &engine{
		driver:    driver,
		librarys:  make(map[LibraryName]*libraryImp),
		listeners: newListenerFromMap(),
	}

	loadLibrarys(e)

	return e
}

type Engine interface {
	Context() ColumnContext
	Librarys() []LibraryName
	GetLibrary(LibraryName) Library
	Listen(Listener) func()
	ListenLibrary(LibraryName, Listener) func()

	// Validate returns cells with errors
	Validate(errorConsumer func(LibraryName, RowID, ColumnID, error)) error

	// Dependency returns list of columns who dependency from LibraryName and ColumnID
	Dependency(LibraryName, ColumnID, func(LibraryName, ColumnID)) error
}

type Listener interface {
	OnNewLibrary(LibraryName)
	OnNewRow(LibraryName, RowID)
	OnDeleteRow(LibraryName, RowID)
	OnUpdateValue(LibraryName, RowID, ColumnID, bool, string, error)
}

type engine struct {
	context         *engineColumnContext
	driver          driver.Driver
	librarys        map[LibraryName]*libraryImp
	listeners       *listenerFromMap
	currentListener func()
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

func (e *engine) Context() ColumnContext {
	if nil == e.context {
		e.context = newEngineColumnContext(e)
	}
	return e.context
}

func (self *engine) Librarys() []LibraryName {
	out := []LibraryName{}

	for k := range self.librarys {
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
		e.Context(),
		e.listeners,
		LensOf(name.ToString()+"$schema", e.driver),
		LensOf(name.ToString()+"$data", e.driver),
		LensOf(name.ToString()+"$meta", e.driver),
	)
	e.librarys[name] = newLib

	e.listeners.OnNewLibrary(name)

	return newLib
}

func (e *engine) Listen(listener Listener) func() {
	return e.listeners.listen(listener)
}

func (e *engine) ListenLibrary(library LibraryName, listener Listener) func() {
	return e.listeners.listenLibrary(library, listener)
}

func (e *engine) Validate(consumer func(LibraryName, RowID, ColumnID, error)) error {
	for _, lib := range e.librarys {
		libraryName := lib.name
		err := libraryValidate(lib, func(row RowID, col ColumnID, err error) {
			consumer(libraryName, row, col, err)
		})
		if nil != err {
			return err
		}
	}
	return nil
}

func (e *engine) Dependency(masterLibrary LibraryName, masterCol ColumnID, consumer func(LibraryName, ColumnID)) error {
	for _, lib := range e.librarys {
		libraryName := lib.name
		err := libraryDependency(masterLibrary, masterCol, lib, func(col ColumnID) {
			consumer(libraryName, col)
		})
		if nil != err {
			return err
		}
	}
	return nil
}
