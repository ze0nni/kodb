package types

import (
	"errors"

	"github.com/ze0nni/kodb/internal/driver"
)

const typePrefix = "type_"

func typesOfDriver(
	driver driver.Driver,
) Types {
	return &types{}
}

type types struct {
}

func (ts *types) Names() ([]string, error) {
	return []string{}, nil
}

func (ts *types) Add(name string) (Type, error) {
	return nil, errors.New("")
}
