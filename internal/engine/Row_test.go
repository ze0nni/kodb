package engine

import (
	"testing"

	"github.com/ze0nni/kodb/internal/entry"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func TestRow_NotExists(t *testing.T) {
	r := RowOf(
		LensOf("x", driver.InMemory()),
		RowID("foo"),
	)

	e, err := r.Exists()

	assert.NoError(t, err)
	assert.False(t, e)
}

func TestRow_Exists(t *testing.T) {
	l := LensOf("x", driver.InMemory())
	id := RowID("foo")

	l.Put(id.ToString(), make(entry.Entry))

	r := RowOf(
		l,
		id,
	)

	e, err := r.Exists()

	assert.NoError(t, err)
	assert.True(t, e)
}

func TestRow_Get_Error_when_row_not_exists(t *testing.T) {
	r := RowOf(
		LensOf("x", driver.InMemory()),
		RowID("foo"),
	)
	_, _, err := r.Get(ColumnID("key"))

	assert.Error(t, err)
}

func TestRow_Get_row_not_found(t *testing.T) {
	l := LensOf("x", driver.InMemory())
	id := RowID("foo")

	l.Put(id.ToString(), make(entry.Entry))

	r := RowOf(
		l,
		id,
	)

	_, ok, err := r.Get(ColumnID("key"))

	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestRow_Get(t *testing.T) {
	l := LensOf("x", driver.InMemory())
	id := RowID("foo")
	key := ColumnID("key")

	e := make(entry.Entry)
	e[key.ToString()] = "value"
	l.Put(id.ToString(), e)

	r := RowOf(
		l,
		id,
	)

	v, ok, err := r.Get(key)

	assert.Equal(t, "value", v)
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestRow_Set_error_when_row_not_exists(t *testing.T) {
	r := RowOf(
		LensOf("x", driver.InMemory()),
		RowID("foo"),
	)

	err := r.Set(ColumnID("key"), "value")

	assert.Error(t, err)
}

func TestRow_Set(t *testing.T) {
	l := LensOf("x", driver.InMemory())
	id := RowID("foo")
	l.Put(id.ToString(), make(entry.Entry))

	r := RowOf(
		l,
		id,
	)

	key := ColumnID("key")
	err := r.Set(key, "value")
	v, ok, _ := r.Get(key)

	assert.NoError(t, err)
	assert.Equal(t, "value", v)
	assert.True(t, ok)
}
