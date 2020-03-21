package engine

import (
	"fmt"
	"strings"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/types"
)

func New(driver driver.Driver) Engine {
	types, err := types.TypesOfDriver(driver)
	if nil != err {
		panic(err)
	}

	e := &engine{
		driver:    driver,
		types:     types,
		librarys:  make(map[LibraryName]*libraryImp),
		listeners: newListenerFromMap(),
	}

	loadLibrarys(e)

	return e
}

type Engine interface {
	Types() types.Types

	Context() ColumnContext
	Librarys() []LibraryName
	Library(LibraryName) (Library, error)

	AddLibrary(LibraryName) (Library, error)

	Listen(Listener) func()
	ListenLibrary(LibraryName, Listener) func()
}

type Listener interface {
	OnNewLibrary(LibraryName)
	OnNewColumn(LibraryName, ColumnID)

	OnNewRow(LibraryName, RowID)
	OnDeleteRow(LibraryName, RowID)
	OnUpdateValue(LibraryName, RowID, ColumnID, bool, string, error)
}

type engine struct {
	context         *engineColumnContext
	driver          driver.Driver
	types           types.Types
	librarys        map[LibraryName]*libraryImp
	listeners       *listenerFromMap
	currentListener func()
}

func loadLibrarys(e *engine) {
	if ps, err := e.driver.Prefixes(); nil == err {
		for _, p := range ps {
			if strings.HasSuffix(p, "$schema") {
				libraryName := LibraryName(p[:len(p)-7])
				e.AddLibrary(libraryName) //TODO: e.recoveryLibrary
			}
		}
	}
}

func (e *engine) Types() types.Types {
	return e.types
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

func (e *engine) Library(name LibraryName) (Library, error) {
	if storedLib := e.librarys[name]; nil != storedLib {
		return storedLib, nil
	}
	return nil, fmt.Errorf("Library <%s> not exits", name)
}

func (e *engine) AddLibrary(name LibraryName) (Library, error) {
	if _, exists := e.librarys[name]; exists {
		return nil, fmt.Errorf("Library <%s> already exits", name)
	}

	newLib := newLibraryInst(
		name,
		e.Context(),
		e.listeners,
		driver.LensOf(name.ToString()+"$schema", e.driver),
		driver.LensOf(name.ToString()+"$data", e.driver),
		driver.LensOf(name.ToString()+"$meta", e.driver),
	)
	e.librarys[name] = newLib

	e.listeners.OnNewLibrary(name)

	return newLib, nil
}

func (e *engine) Listen(listener Listener) func() {
	return e.listeners.listen(listener)
}

func (e *engine) ListenLibrary(library LibraryName, listener Listener) func() {
	return e.listeners.listenLibrary(library, listener)
}
