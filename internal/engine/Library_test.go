package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/kodb/internal/driver"
)

func emptyTestLibrary() (Library, driver.Driver) {
	return testLibraryOfDriver(driver.InMemory())
}

func testLibraryOfDriver(d driver.Driver) (Library, driver.Driver) {
	schema := driver.LensOf("schema", d)
	data := driver.LensOf("data", d)
	meta := driver.LensOf("meta", d)

	return newLibraryInst(
		LibraryName("libraryName"),
		newNilColumnContext(),
		listenerNil(),
		schema,
		data,
		meta,
	), d
}

func TestLibrary_resoreDataFromDriver(t *testing.T) {
	l, d := emptyTestLibrary()

	r1, _ := l.NewRow()
	r2, _ := l.NewRow()
	r3, _ := l.NewRow()

	l2, _ := testLibraryOfDriver(d)

	r1x, _ := l2.RowID(0)
	r2x, _ := l2.RowID(1)
	r3x, _ := l2.RowID(2)

	assert.Equal(t, r1, r1x)
	assert.Equal(t, r2, r2x)
	assert.Equal(t, r3, r3x)
}

func TestLibrary_Swap_error_for_empty_library(t *testing.T) {
	l, _ := emptyTestLibrary()

	err := l.Swap(0, 0)

	assert.Error(t, err)
}

func TestLibrary_Swap_line_with_self(t *testing.T) {
	l, _ := emptyTestLibrary()

	r1, _ := l.NewRow()
	r2, _ := l.NewRow()
	r3, _ := l.NewRow()

	err := l.Swap(1, 1)

	r1x, _ := l.RowID(0)
	r2x, _ := l.RowID(1)
	r3x, _ := l.RowID(2)

	assert.NoError(t, err)

	assert.Equal(t, r1, r1x)
	assert.Equal(t, r2, r2x)
	assert.Equal(t, r3, r3x)
}

func TestLibrary_Swap_in_driver(t *testing.T) {
	l, d := emptyTestLibrary()

	r1, _ := l.NewRow()
	r2, _ := l.NewRow()
	r3, _ := l.NewRow()

	l.Swap(0, 2)
	l2, _ := testLibraryOfDriver(d)

	r1x, _ := l2.RowID(0)
	r2x, _ := l2.RowID(1)
	r3x, _ := l2.RowID(2)

	assert.Equal(t, r3, r1x)
	assert.Equal(t, r2, r2x)
	assert.Equal(t, r1, r3x)
}
