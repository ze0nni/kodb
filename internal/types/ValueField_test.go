package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValueField_Raname(t *testing.T) {
	f := NewValueFieldData("name")

	f.Rename("newName")

	assert.Equal(t, "newName", f.Name())
}

func Test_ValueField_toEntry(t *testing.T) {
	f := NewValueFieldData("value")

	e := f.toEntry()

	assert.Equal(t, ValueFieldKind.String(), e["kind"])
	assert.Equal(t, "value", e["name"])
}

func Test_ValueField_fromEntry(t *testing.T) {
	f0 := NewValueFieldData("value")

	e := f0.toEntry()

	f := NewValueFieldData("")
	err := f.fromEntry(e)

	assert.NoError(t, err)
	assert.Equal(t, ValueFieldKind, f.Kind())
	assert.Equal(t, "value", f.Name())
}

func Test_ValueField_fromEntry_with_wrong_kind(t *testing.T) {
	f0 := NewValueFieldData("value")

	e := f0.toEntry()
	e["kind"] = "wrongKind"

	f := NewValueFieldData("")
	err := f.fromEntry(e)

	assert.Error(t, err)
}

func Test_ValueField_fromEntry_with_wrong_id(t *testing.T) {
	f0 := NewValueFieldData("value")

	e := f0.toEntry()
	e["id"] = f0.ID().String() + "_salt"

	f := NewValueFieldData("")
	err := f.fromEntry(e)

	assert.Error(t, err)
}

func Test_ValueField_call_listener_on_rename(t *testing.T) {
	f := NewValueFieldData("oldName")

	calls := 0
	f.setListener(func() {
		calls++
	})

	f.Rename("name1")
	f.Rename("name2")
	f.Rename("name3")

	assert.Equal(t, 3, calls)
}

func Test_ValueField_no_call_after_clean_listener(t *testing.T) {
	f := NewValueFieldData("oldName")

	calls := 0
	f.setListener(func() {
		calls++
	})

	f.Rename("name1")
	f.setListener(nil)

	f.Rename("name2")
	f.Rename("name3")

	assert.Equal(t, 1, calls)
}
