package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commonType_Fields_empty(t *testing.T) {
	f := newCommonType("newType")

	fs := f.Fields()

	assert.NotNil(t, fs)
	assert.Len(t, fs, 0)
}

func Test_commonType_New(t *testing.T) {
	f := newCommonType("newType")

	fs, err := f.New(NewValueFieldData("fName"))
	fields := f.Fields()

	assert.NoError(t, err)
	assert.Equal(t, "fName", fs.Name())
	assert.Equal(t, ValueFieldKind, fs.Kind())
	assert.Equal(t, fields, []Field{fs})
}
