package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ze0nni/kodb/internal/driver"
)

func Test_typesOfDriver_Names_returns_emptyList(t *testing.T) {
	types := typesOfDriver(driver.InMemory())

	list, err := types.Names()

	assert.Equal(t, []string{}, list)
	assert.NoError(t, err)
}

func Test_typesOfDriver_Add(t *testing.T) {
	types := typesOfDriver(driver.InMemory())

	tp, err := types.Add("newType")

	assert.NotNil(t, tp)
	assert.NoError(t, err)
}
