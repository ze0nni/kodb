package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemory(t *testing.T) {
	m := InMemory()
	assert.NotNil(t, m)
}

func TestInMemory_EntryNotExists(t *testing.T) {
	m := InMemory()

	err, e := m.Get("foo", "bar")

	assert.NoError(t, err)

	assert.Nil(t, e)
}

func TestInMemory_Put(t *testing.T) {
	m := InMemory()

	err := m.Put("foo", "bar", make(Entry))

	assert.NoError(t, err)

	err, e := m.Get("foo", "bar")

	assert.NoError(t, err)

	assert.NotNil(t, e)
}
