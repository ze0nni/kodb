package engine

import (
	"testing"

	"github.com/ze0nni/kodb/internal/entry"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/kodb/internal/driver"
)

func TestLibrary_Name(t *testing.T) {
	l := newLibraryInst("foo", listenerNil(), LensOf("schema", driver.InMemory()), nil, nil)

	assert.Equal(t, LibraryName("foo"), l.Name())
}

func TestLibrary_Columns(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", listenerNil(), LensOf("schema", d), nil, nil)

	assert.Equal(t, 0, l.Columns())
}

func TestLibrary_NewColumn(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", listenerNil(), LensOf("schema", d), nil, nil)

	_, err := l.NewColumn("bar")

	assert.NoError(t, err)

	assert.Equal(t, 1, l.Columns())
}

func TestLibrary_ColumnName(t *testing.T) {
	d := driver.InMemory()
	l := newLibraryInst("foo", listenerNil(), LensOf("schema", d), nil, nil)

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
	l := newLibraryInst("foo", listenerNil(), LensOf("schema", driver.InMemory()), nil, nil)

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

func TestLibrary_DeleteRow_error_when_row_not_exists(t *testing.T) {
	l, _ := emptyUsersLibrary()

	err := l.DeleteRow(RowID("foo"))

	assert.Error(t, err)
}

func TestLibrary_DeleteRow(t *testing.T) {
	l, _ := emptyUsersLibrary()
	id, _ := l.NewRow()

	assert.True(t, l.HasRow(id))

	err := l.DeleteRow(id)

	assert.NoError(t, err)
	assert.False(t, l.HasRow(id))
}

func TestLibrary_Row_out_of_range(t *testing.T) {
	l, _ := emptyLibrary("foo")
	_, err := l.Row(0)
	assert.Error(t, err)
}

func TestLibrary_Row(t *testing.T) {
	l, _ := emptyLibrary("foo")
	id, _ := l.NewRow()
	r, err := l.Row(0)
	exists, _ := r.Exists()

	assert.True(t, exists)
	assert.Equal(t, id, r.ID())
	assert.NoError(t, err)
}

func TestLibrary_Row_ref_to_same_data(t *testing.T) {
	l, _ := emptyLibrary("foo")
	l.NewRow()

	r1, _ := l.Row(0)
	r2, _ := l.Row(0)

	assert.NotSame(t, r1, r2)

	key := ColumnID("key")
	r1.Set(key, "value")
	v, _, _ := r2.Get(key)

	assert.Equal(t, "value", v)
}

func TestLibrary_RowID(t *testing.T) {
	l, _ := emptyLibrary("foo")
	id, _ := l.NewRow()
	r, err := l.RowID(0)

	assert.Equal(t, id, r)
	assert.NoError(t, err)
}

func TestLibrary_GetRowColumn_row_not_exists(t *testing.T) {
	l, _ := emptyLibrary("foo")
	key := ColumnID("key")

	v, ok, err := l.GetRowColumn(0, key)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	assert.Error(t, err)
}

func TestLibrary_GetRowColumn_column_not_exists(t *testing.T) {
	l, _ := emptyLibrary("foo")
	l.NewRow()
	key := ColumnID("key")

	v, ok, err := l.GetRowColumn(0, key)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestLibrary_GetRowColumn(t *testing.T) {
	l, _ := emptyLibrary("foo")
	l.NewRow()
	r, _ := l.Row(0)
	key := ColumnID("key")
	r.Set(key, "value")

	v, ok, err := l.GetRowColumn(0, key)
	assert.Equal(t, "value", v)
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestLibrary_GetValue_error_when_row_not_exists(t *testing.T) {
	l, _ := emptyLibrary("foo")

	_, exists, err := l.GetValue(RowID("foo"), ColumnID("bar"))

	assert.False(t, exists)
	assert.Error(t, err)
}

func TestLibrary_GetValue(t *testing.T) {
	l, d := emptyLibrary("foo")

	rowID, _ := l.NewRow()

	e := make(entry.Entry)
	e["bar"] = "baz"
	d.Put("data", rowID.ToString(), e)

	value, exists, err := l.GetValue(rowID, ColumnID("bar"))

	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, "baz", value)
}

func TestLibrary_UpdateValue_error_when_row_not_exists(t *testing.T) {
	l, _ := emptyLibrary("foo")

	err := l.UpdateValue(RowID("foo"), ColumnID("bar"), "baz")

	assert.Error(t, err)
}

func TestLibrary_UpdateValue(t *testing.T) {
	l, _ := emptyLibrary("foo")

	r, _ := l.NewRow()
	err := l.UpdateValue(r, ColumnID("bar"), "baz")

	assert.NoError(t, err)

	value, exists, _ := l.GetValue(r, ColumnID("bar"))
	assert.True(t, exists)
	assert.Equal(t, "baz", value)
}

func emptyLibrary(libraryName LibraryName) (Library, driver.Driver) {
	d := driver.InMemory()
	l := newLibraryInst(
		libraryName,
		listenerNil(),
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
