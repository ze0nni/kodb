package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func Test_TypesOfDriver_Names_returns_emptyList(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	list := types.Names()

	assert.Equal(t, []TypeName{}, list)
}

func Test_TypesOfDriver_New(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	tp, err := types.New("newType")
	list := types.Names()

	assert.NotNil(t, tp)
	assert.NoError(t, err)
	assert.Equal(t, TypeName("newType"), tp.Name())
	assert.Equal(t, []TypeName{TypeName("newType")}, list)
}

func Test_TypesOfDriver_New_error_on_duplicate(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	types.New(TypeName("newType"))
	tp, err := types.New(TypeName("newType"))

	assert.Nil(t, tp)
	assert.Error(t, err)
}

func Test_TypesOfDriver_Get_returns_error_when_type_not_exists(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	tp, err := types.Get(TypeName("newType"))

	assert.Nil(t, tp)
	assert.Error(t, err)
}

func Test_TypesOfDriver_Get_returns_exists_type(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	types.New(TypeName("newType"))
	tp, err := types.Get(TypeName("newType"))

	assert.NotNil(t, tp)
	assert.Equal(t, TypeName("newType"), tp.Name())
	assert.NoError(t, err)
}

func Test_TypesOfDriver_Delete_returns_error_when_type_not_exists(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	err := types.Delete(TypeName("newType"))

	assert.Error(t, err)
}

func Test_TypesOfDriver_Delete_exists_type(t *testing.T) {
	types, _ := TypesOfDriver(driver.InMemory())

	types.New(TypeName("newType"))
	err := types.Delete(TypeName("newType"))
	names := types.Names()

	assert.NoError(t, err)
	assert.Equal(t, []TypeName{}, names)
}

func Test_TypesOfDriver_restore_types_from_Driver(t *testing.T) {
	dr := driver.InMemory()
	types, _ := TypesOfDriver(dr)

	types.New(TypeName("type1"))
	types.New(TypeName("type2"))
	types.New(TypeName("type3"))
	types.Delete(TypeName("type2"))

	newTypes, _ := TypesOfDriver(dr)
	list := newTypes.Names()

	assert.ElementsMatch(t, []TypeName{TypeName("type1"), TypeName("type3")}, list)
}
