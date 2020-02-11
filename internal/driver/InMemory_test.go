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
