package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/kodb/internal/driver"
)

func emptyLens() driver.Lens {
	return driver.LensOf("prefix", driver.InMemory())
}

func Test_commonType_Fields_empty(t *testing.T) {
	f := newCommonType("newType", emptyLens(), newTypesListener())

	fs := f.Fields()

	assert.NotNil(t, fs)
	assert.Len(t, fs, 0)
}

func Test_commonType_New(t *testing.T) {
	f := newCommonType("newType", emptyLens(), newTypesListener())

	fs, err := f.New(NewValueFieldData("fName"))
	fields := f.Fields()

	assert.NoError(t, err)
	assert.NotEqual(t, FieldID(""), fs.ID())
	assert.Equal(t, "fName", fs.Name())
	assert.Equal(t, ValueFieldKind, fs.Kind())
	assert.Equal(t, fields, []Field{fs})
}

func Test_commonType_New_error_if_FieldNotNew(t *testing.T) {
	f := newCommonType("newType", emptyLens(), newTypesListener())

	fs, _ := f.New(NewValueFieldData("fName"))
	fs2, err := f.New(fs)
	fields := f.Fields()

	assert.Error(t, err)
	assert.Nil(t, fs2)
	assert.Len(t, fields, 1)
}

func Test_commonType_Delete_error_if_field_not_exists(t *testing.T) {
	tp := newCommonType("newType", emptyLens(), newTypesListener())

	err := tp.Delete(NewValueFieldData("fname"))

	assert.Error(t, err)
}

func Test_commonType_Delete(t *testing.T) {
	tp := newCommonType("newType", emptyLens(), newTypesListener())

	f1, _ := tp.New(NewValueFieldData("f1"))
	f2, _ := tp.New(NewValueFieldData("f2"))
	f3, _ := tp.New(NewValueFieldData("f3"))

	err := tp.Delete(f2)
	fields := tp.Fields()

	assert.NoError(t, err)
	assert.ElementsMatch(t, []Field{f1, f3}, fields)
}
