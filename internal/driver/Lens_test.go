package driver

import (
	"testing"

	"github.com/ze0nni/kodb/internal/entry"

	"github.com/stretchr/testify/assert"
)

func TestLens_Get(t *testing.T) {
	l := LensOf("foo", InMemory())

	e, err := l.Get("bar")

	assert.NoError(t, err)
	assert.Nil(t, e)
}

func TestLens_Put(t *testing.T) {
	l := LensOf("foo", InMemory())

	err := l.Put("bar", make(entry.Entry))

	e, err := l.Get("bar")

	assert.NoError(t, err)
	assert.NotNil(t, e)
}

func TestLens_Update(t *testing.T) {
	l := LensOf("foo", InMemory())

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
	d := InMemory()

	foo := LensOf("foo", d)
	bar := LensOf("bar", d)

	err := foo.Put("x", make(entry.Entry))

	assert.NoError(t, err)

	e, err := bar.Get("x")

	assert.NoError(t, err)

	assert.Nil(t, e)
}

func TestLens_driver(t *testing.T) {
	d := InMemory()

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

func TestLens_Keys(t *testing.T) {
	l := LensOf("prefix", InMemory())

	l.Put("1", entry.Entry{})
	l.Put("2", entry.Entry{})
	l.Put("6", entry.Entry{})
	l.Put("1", entry.Entry{})

	keys, err := l.Keys()

	assert.NoError(t, err)
	assert.ElementsMatch(
		t,
		[]string{"1", "2", "6"},
		keys,
	)
}
