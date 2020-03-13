package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func TestEngine_GetLibrary(t *testing.T) {
	eng := New(driver.InMemory())

	lib, _ := eng.AddLibrary("foo")

	assert.NotNil(t, lib)
}

func TestEngine_GetSameLibrarys(t *testing.T) {
	eng := New(driver.InMemory())

	lib1, _ := eng.AddLibrary("foo")
	lib2, _ := eng.AddLibrary("foo")

	assert.Same(t, lib1, lib2)
}

func TestEngine_Library_Name(t *testing.T) {
	eng := New(driver.InMemory())

	foo, _ := eng.AddLibrary("foo")

	assert.Equal(t, LibraryName("foo"), foo.Name())
}

func TestEngine_Librarys_empty(t *testing.T) {
	eng := New(driver.InMemory())

	ls := eng.Librarys()

	assert.Equal(t, []LibraryName{}, ls)
}

func TestEngine_Librarys(t *testing.T) {
	eng := New(driver.InMemory())

	eng.AddLibrary("foo")
	eng.AddLibrary("bar")
	ls := eng.Librarys()

	assert.ElementsMatch(t, []LibraryName{LibraryName("foo"), LibraryName("bar")}, ls)
}

func TestEngine_Librarys_onLoad(t *testing.T) {
	m := driver.InMemory()
	eng1 := New(m)

	eng1.AddLibrary("foo")
	eng1.AddLibrary("bar")

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

	eng.AddLibrary(LibraryName("foo"))

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

	eng.AddLibrary(LibraryName("foo"))

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

	foo, _ := eng.AddLibrary(LibraryName("foo"))
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

	foo, _ := eng.AddLibrary(LibraryName("foo"))
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

	colID := ColumnID("name")

	foo, _ := eng.AddLibrary(LibraryName("foo"))
	foo.AddColumn(colID, "no name")
	foo.AddRow(RowID("bar"))

	ll := newLogListener()
	eng.Listen(ll)

	foo.UpdateValue(RowID("bar"), colID, "hello world")

	assert.Equal(
		t,
		[]string{"updateRow foo:bar:name true hello world"},
		ll.getLog(),
	)
}
