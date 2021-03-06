package engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/types"
)

func TestEngine_GetLibrary(t *testing.T) {
	eng := New(driver.InMemory())

	lib, _ := eng.AddLibrary("foo", types.TypeName(""))

	assert.NotNil(t, lib)
}

func TestEngine_GetSameLibrarys(t *testing.T) {
	eng := New(driver.InMemory())

	lib1, _ := eng.AddLibrary("foo", types.TypeName(""))
	lib2, _ := eng.AddLibrary("foo", types.TypeName(""))

	assert.Same(t, lib1, lib2)
}

func TestEngine_Library_Name(t *testing.T) {
	eng := New(driver.InMemory())

	foo, _ := eng.AddLibrary("foo", types.TypeName(""))

	assert.Equal(t, LibraryName("foo"), foo.Name())
}

func TestEngine_Librarys_empty(t *testing.T) {
	eng := New(driver.InMemory())

	ls := eng.Librarys()

	assert.Equal(t, []LibraryName{}, ls)
}

func TestEngine_Librarys(t *testing.T) {
	eng := New(driver.InMemory())

	eng.AddLibrary("foo", types.TypeName(""))
	eng.AddLibrary("bar", types.TypeName(""))
	ls := eng.Librarys()

	assert.ElementsMatch(t, []LibraryName{LibraryName("foo"), LibraryName("bar")}, ls)
}

func TestEngine_Librarys_onLoad(t *testing.T) {
	m := driver.InMemory()
	eng1 := New(m)

	eng1.AddLibrary("foo", types.TypeName(""))
	eng1.AddLibrary("bar", types.TypeName(""))

	eng2 := New(m)
	ls := eng2.Librarys()

	assert.ElementsMatch(t, []LibraryName{LibraryName("foo"), LibraryName("bar")}, ls)
}

func TestEngine_Listener(t *testing.T) {
	eng := New(driver.InMemory())

	l := eng.Listen(nil)

	assert.NotNil(t, l)
}

func TestEngine_Listener_NewLibrary(t *testing.T) {
	eng := New(driver.InMemory())

	ll := newLogListener()
	eng.Listen(ll)

	eng.AddLibrary(LibraryName("foo"), types.TypeName(""))

	assert.Equal(
		t,
		[]string{"newLibrary foo"},
		ll.getLog(),
	)
}

func TestEngine_Listener_StopListen(t *testing.T) {
	eng := New(driver.InMemory())

	ll := newLogListener()
	handle := eng.Listen(ll)
	handle()

	eng.AddLibrary(LibraryName("foo"), types.TypeName(""))

	assert.Equal(
		t,
		[]string{},
		ll.getLog(),
	)
}

func TestEngine_Listener_NewRow(t *testing.T) {
	eng := New(driver.InMemory())

	ll := newLogListener()
	eng.Listen(ll)

	foo, _ := eng.AddLibrary(LibraryName("foo"), types.TypeName(""))
	foo.AddRow(RowID("bar"))

	assert.Equal(
		t,
		[]string{"newLibrary foo", "newRow foo:bar"},
		ll.getLog(),
	)
}

func TestEngine_Listener_DeleteRow(t *testing.T) {
	eng := New(driver.InMemory())

	ll := newLogListener()
	eng.Listen(ll)

	foo, _ := eng.AddLibrary(LibraryName("foo"), types.TypeName(""))
	foo.AddRow(RowID("bar"))
	foo.DeleteRow(RowID("bar"))

	assert.Equal(
		t,
		[]string{"newLibrary foo", "newRow foo:bar", "deleteRow foo:bar"},
		ll.getLog(),
	)
}

func TestEngine_Listener_UpdateValue(t *testing.T) {
	eng := New(driver.InMemory())
	tp, _ := eng.Types().New("typeName")
	field, _ := tp.New(types.NewValueFieldData("field"))

	foo, _ := eng.AddLibrary(LibraryName("foo"), tp.Name())
	rowID, _ := foo.NewRow()

	ll := newLogListener()
	eng.Listen(ll)

	foo.UpdateValue(rowID, field.ID(), "hello world")

	assert.Equal(
		t,
		[]string{fmt.Sprintf("updateRow foo:%s:%s true hello world", rowID.ToString(), field.ID().String())},
		ll.getLog(),
	)
}

func TestEngine_Listener_Swap(t *testing.T) {
	eng := New(driver.InMemory())
	l, _ := eng.AddLibrary(LibraryName("library"), types.TypeName(""))

	l.AddRow(RowID("r1"))
	l.AddRow(RowID("r2"))
	l.AddRow(RowID("r3"))
	l.AddRow(RowID("r4"))

	ll := newLogListener()
	eng.Listen(ll)

	l.Swap(1, 3)

	assert.Equal(
		t,
		[]string{"swap library 1 3 r4 r2"},
		ll.getLog(),
	)
}
