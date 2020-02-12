package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func TestEngine_GetLibrary(t *testing.T) {
	eng := New(driver.InMemory())

	lib := eng.GetLibrary("foo")

	assert.NotNil(t, lib)
}

func TestEngine_GetSameLibrarys(t *testing.T) {
	eng := New(driver.InMemory())

	lib1 := eng.GetLibrary("foo")
	lib2 := eng.GetLibrary("foo")

	assert.Same(t, lib1, lib2)
}

func TestEngine_Library_Name(t *testing.T) {
	eng := New(driver.InMemory())

	foo := eng.GetLibrary("foo")

	assert.Equal(t, LibraryName("foo"), foo.Name())
}

func TestEngine_Librarys_empty(t *testing.T) {
	eng := New(driver.InMemory())

	ls := eng.Librarys()

	assert.Equal(t, []LibraryName{}, ls)
}

func TestEngine_Librarys(t *testing.T) {
	eng := New(driver.InMemory())

	eng.GetLibrary("foo")
	eng.GetLibrary("bar")
	ls := eng.Librarys()

	assert.ElementsMatch(t, []LibraryName{LibraryName("foo"), LibraryName("bar")}, ls)
}

func TestEngine_Librarys_onLoad(t *testing.T) {
	m := driver.InMemory()
	eng1 := New(m)

	eng1.GetLibrary("foo")
	eng1.GetLibrary("bar")

	eng2 := New(m)
	ls := eng2.Librarys()

	assert.ElementsMatch(t, []LibraryName{LibraryName("foo"), LibraryName("bar")}, ls)
}
