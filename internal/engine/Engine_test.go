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

	assert.Equal(t, "foo", foo.Name())
}
