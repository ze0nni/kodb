package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/kodb/internal/entry"
)

func TestInMemory(t *testing.T) {
	m := InMemory()
	assert.NotNil(t, m)
}

func TestInMemory_EntryNotExists(t *testing.T) {
	m := InMemory()

	e, err := m.Get("foo", "bar")

	assert.NoError(t, err)

	assert.Nil(t, e)
}

func TestInMemory_Put(t *testing.T) {
	m := InMemory()

	err := m.Put("foo", "bar", make(entry.Entry))

	assert.NoError(t, err)

	e, err := m.Get("foo", "bar")

	assert.NoError(t, err)

	assert.NotNil(t, e)
}

func TestInMemory_PutEquals(t *testing.T) {
	m := InMemory()

	e1 := make(entry.Entry)
	e1["hello"] = "world"

	err := m.Put("foo", "bar", e1)

	assert.NoError(t, err)

	e2, err := m.Get("foo", "bar")

	assert.NoError(t, err)

	assert.Equal(t, "world", e2["hello"])
	assert.Equal(t, "", e2["world"])
}

func TestInMemory_Prefixes_Empty(t *testing.T) {
	m := InMemory()

	p, err := m.Prefixes()

	assert.NoError(t, err)
	assert.Equal(t, []string{}, p)
}

func TestInMemory_Prefixes(t *testing.T) {
	m := InMemory()
	m.Put("foo", "value", make(entry.Entry))

	p, err := m.Prefixes()

	assert.NoError(t, err)
	assert.Equal(t, []string{"foo"}, p)
}

func TestInMemory_Delete_error_when_not_exists(t *testing.T) {
	m := InMemory()

	err := m.Delete("foo", "value")

	assert.Error(t, err)
}

func TestInMemory_Delete(t *testing.T) {
	m := InMemory()
	m.Put("foo", "value", make(entry.Entry))

	err := m.Delete("foo", "value")

	assert.NoError(t, err)

	e, err := m.Get("foo", "value")
	assert.NoError(t, err)
	assert.Nil(t, e)
}

func TestInMemory_DeletePrefix(t *testing.T) {
	m := InMemory()

	m.Put("foo", "1", make(entry.Entry))
	m.Put("foo", "2", make(entry.Entry))
	m.Put("foo", "3", make(entry.Entry))
	m.Put("bar", "1", make(entry.Entry))
	m.Put("bar", "2", make(entry.Entry))
	m.Put("bar", "3", make(entry.Entry))

	m.DeletePrefix("foo")

	f1, _ := m.Get("foo", "1")
	f2, _ := m.Get("foo", "2")
	f3, _ := m.Get("foo", "3")
	b1, _ := m.Get("bar", "1")
	b2, _ := m.Get("bar", "2")
	b3, _ := m.Get("bar", "3")

	assert.Nil(t, f1)
	assert.Nil(t, f2)
	assert.Nil(t, f3)
	assert.NotNil(t, b1)
	assert.NotNil(t, b2)
	assert.NotNil(t, b3)
}
