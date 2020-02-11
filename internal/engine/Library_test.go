package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/kodb/internal/driver"
)

func TestLibrary_Name(t *testing.T) {
	l := newLibraryInst("foo", nil, nil, nil)

	assert.Equal(t, "foo", l.Name())
}

func TestLibrary_Columns(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", LensOf("schema", d), nil, nil)

	assert.Equal(t, 0, l.Columns())
}

func TestLibrary_NewColumn(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", LensOf("schema", d), nil, nil)

	err := l.AddColumn("bar")

	assert.NoError(t, err)

	assert.Equal(t, 1, l.Columns())
}

func TestLibrary_NewColumn_duplicate(t *testing.T) {
	assert.Fail(t, "todo")
}

func TestLibrary_Column(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", LensOf("schema", d), nil, nil)

	l.AddColumn("foo")
	l.AddColumn("bar")

	c1, err := l.Column(0)
	assert.NoError(t, err)

	c2, err := l.Column(1)
	assert.NoError(t, err)

	assert.Equal(t, "foo", c1)
	assert.Equal(t, "bar", c2)
}
