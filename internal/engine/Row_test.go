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
