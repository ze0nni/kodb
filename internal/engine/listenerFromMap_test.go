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

	l.OnNewLibrary(foo)
	l.OnNewLibrary(bar)
	l.OnNewRow(foo, RowID("row"))
	l.OnNewRow(bar, RowID("row"))
	l.OnUpdateValue(foo, RowID("row"), ColumnID("col"), true, "value", nil)
	l.OnUpdateValue(bar, RowID("row"), ColumnID("col"), true, "value", nil)
	l.OnDeleteRow(foo, RowID("row"))
	l.OnDeleteRow(bar, RowID("row"))

	assert.Equal(
		t,
		[]string{"newLibrary foo", "newRow foo:row", "updateRow foo:row:col true value", "deleteRow foo:row"},
		ll.getLog(),
	)
}
