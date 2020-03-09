package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenerFromMap_ListenLibrary(t *testing.T) {
	l := newListenerFromMap()

	ll := newLogListener()

	foo := LibraryName("foo")
	bar := LibraryName("bar")
	l.listenLibrary(foo, ll)

	l.NewLibrary(foo)
	l.NewLibrary(bar)
	l.NewRow(foo, RowID("row"))
	l.NewRow(bar, RowID("row"))
	l.UpdateValue(foo, RowID("row"), ColumnID("col"), true, "value", nil)
	l.UpdateValue(bar, RowID("row"), ColumnID("col"), true, "value", nil)
	l.DeleteRow(foo, RowID("row"))
	l.DeleteRow(bar, RowID("row"))

	assert.Equal(
		t,
		[]string{"newLibrary foo", "newRow foo:row", "updateRow foo:row:col true value", "deleteRow foo:row"},
		ll.getLog(),
	)
}
