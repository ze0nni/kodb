package engine

import (
	"testing"

	"github.com/ze0nni/kodb/internal/entry"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func TestLens_Get(t *testing.T) {
	l := LensOf("foo", driver.InMemory())

	e, err := l.Get("bar")

	assert.NoError(t, err)
	assert.Nil(t, e)
}

func TestLens_Put(t *testing.T) {
	l := LensOf("foo", driver.InMemory())

	err := l.Put("bar", make(entry.Entry))

	e, err := l.Get("bar")

	assert.NoError(t, err)
	assert.NotNil(t, e)
}

func TestLens_Update(t *testing.T) {
	l := LensOf("foo", driver.InMemory())

	e1 := make(entry.Entry)
	e1["value"] = "foo"

	e2 := make(entry.Entry)
	e2["value"] = "bar"

	err := l.Put("foo", e1)
	assert.NoError(t, err)

	err = l.Put("foo", e2)
	assert.NoError(t, err)

	e3, err := l.Get("foo")
	assert.NoError(t, err)

	assert.Equal(t, "bar", e3["value"])
}

func TestLens_prefix(t *testing.T) {
	d := driver.InMemory()

	foo := LensOf("foo", d)
	bar := LensOf("bar", d)

	err := foo.Put("x", make(entry.Entry))

	assert.NoError(t, err)

	e, err := bar.Get("x")

	assert.NoError(t, err)

	assert.Nil(t, e)
}

func TestLens_driver(t *testing.T) {
	d := driver.InMemory()

	foo1 := LensOf("foo", d)

	e1 := make(entry.Entry)
	e1["hello"] = "world"

	err := foo1.Put("x", e1)

	assert.NoError(t, err)

	foo2 := LensOf("foo", d)

	e2, err := foo2.Get("x")

	assert.NoError(t, err)
	assert.NotNil(t, e2)
	assert.Equal(t, "world", e2["hello"])
}
