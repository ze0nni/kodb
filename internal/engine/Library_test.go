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

	_, err := l.NewColumn("bar")

	assert.NoError(t, err)

	assert.Equal(t, 1, l.Columns())
}

func TestLibrary_ColumnName(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", LensOf("schema", d), nil, nil)

	l.NewColumn("foo")
	l.NewColumn("bar")

	c1, err := l.ColumnName(0)
	assert.NoError(t, err)

	c2, err := l.ColumnName(1)
	assert.NoError(t, err)

	assert.Equal(t, "foo", c1)
	assert.Equal(t, "bar", c2)
}

func TestLibrary_AddCoumn_error_on_duplicate(t *testing.T) {
	l := newLibraryInst("foo", LensOf("schema", driver.InMemory()), nil, nil)

	err1 := l.AddColumn(ColumnID("foo"), "foo")
	err2 := l.AddColumn(ColumnID("foo"), "foo")

	assert.NoError(t, err1)
	assert.Error(t, err2)
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

	l.AddRow(RowID("foo"))

	assert.Equal(t, 1, l.Rows())
}

func TestLibrary_AddRow_errorOnDuplicate(t *testing.T) {
	l, _ := emptyUsersLibrary()

	l.AddRow(RowID("foo"))
	err := l.AddRow(RowID("foo"))

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

	assert.False(t, l.HasRow(RowID("foo")))
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
	l.NewColumn("first_name")
	l.NewColumn("second_name")
	l.NewColumn("age")
	return l, d
}
