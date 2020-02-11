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
