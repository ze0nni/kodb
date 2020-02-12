package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/kodb/internal/driver"
)

func TestLibrary_Name(t *testing.T) {
	l := newLibraryInst("foo", LensOf("schema", driver.InMemory()), nil, nil)

	assert.Equal(t, LibraryName("foo"), l.Name())
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

func TestLibrary_Rows_empty(t *testing.T) {
	l, _ := emptyUsersLibrary()

	assert.Equal(t, 0, l.Rows())
}

func TestLibrary_AddRow_NewRow(t *testing.T) {
	l, _ := emptyUsersLibrary()

	_, err := l.NewRow()

	assert.NoError(t, err)
	assert.Equal(t, 1, l.Rows())
}

func TestLibrary_AddRow_NewRow_IdsNotEqual(t *testing.T) {
	l, _ := emptyUsersLibrary()

	id1, _ := l.NewRow()
	id2, _ := l.NewRow()

	assert.NotEqual(t, id1, id2)
}

func TestLibrary_AddRow(t *testing.T) {
	l, _ := emptyUsersLibrary()

	l.AddRow(RowId("foo"))

	assert.Equal(t, 1, l.Rows())
}

func TestLibrary_AddRow_errorOnDuplicate(t *testing.T) {
	l, _ := emptyUsersLibrary()

	l.AddRow(RowId("foo"))
	err := l.AddRow(RowId("foo"))

	assert.Error(t, err)
	assert.Equal(t, 1, l.Rows())
}

func TestLibrary_HasRow(t *testing.T) {
	l, _ := emptyUsersLibrary()

	rowId, _ := l.NewRow()

	assert.True(t, l.HasRow(rowId))
}

func TestLibrary_HasRow_notFound(t *testing.T) {
	l, _ := emptyUsersLibrary()

	assert.False(t, l.HasRow(RowId("foo")))
}

func emptyLibrary(libraryName LibraryName) (Library, driver.Driver) {
	d := driver.InMemory()
	l := newLibraryInst(
		libraryName,
		LensOf("schema", d),
		LensOf("data", d),
		LensOf("meta", d),
	)
	return l, d
}

func emptyUsersLibrary() (Library, driver.Driver) {
	l, d := emptyLibrary(LibraryName("users"))
	l.AddColumn("first_name")
	l.AddColumn("second_name")
	l.AddColumn("age")
	return l, d
}
