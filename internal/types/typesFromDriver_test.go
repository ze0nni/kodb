package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func Test_typesOfDriver_Names_returns_emptyList(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	list := types.Names()

	assert.Equal(t, []string{}, list)
}

func Test_typesOfDriver_New(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	tp, err := types.New("newType")
	list := types.Names()

	assert.NotNil(t, tp)
	assert.NoError(t, err)
	assert.Equal(t, "newType", tp.Name())
	assert.Equal(t, []string{"newType"}, list)
}

func Test_typesOfDriver_New_error_on_duplicate(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	types.New("newType")
	tp, err := types.New("newType")

	assert.Nil(t, tp)
	assert.Error(t, err)
}

func Test_typesOfDriver_Get_returns_error_when_type_not_exists(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	tp, err := types.Get("newType")

	assert.Nil(t, tp)
	assert.Error(t, err)
}

func Test_typesOfDriver_Get_returns_exists_type(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	types.New("newType")
	tp, err := types.Get("newType")

	assert.NotNil(t, tp)
	assert.Equal(t, "newType", tp.Name())
	assert.NoError(t, err)
}

func Test_typesOfDriver_Delete_returns_error_when_type_not_exists(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	err := types.Delete("newType")

	assert.Error(t, err)
}

func Test_typesOfDriver_Delete_exists_type(t *testing.T) {
	types, _ := typesOfDriver(driver.InMemory())

	types.New("newType")
	err := types.Delete("newType")
	names := types.Names()

	assert.NoError(t, err)
	assert.Equal(t, []string{}, names)
}

func Test_typesOfDriver_restore_types_from_Driver(t *testing.T) {
	dr := driver.InMemory()
	types, _ := typesOfDriver(dr)

	types.New("type1")
	types.New("type2")
	types.New("type3")
	types.Delete("type2")

	newTypes, _ := typesOfDriver(dr)
	list := newTypes.Names()

	assert.ElementsMatch(t, []string{"type1", "type3"}, list)
}
