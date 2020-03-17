package engine

import "github.com/ze0nni/kodb/internal/driver"

func typesOfDriver(
	driver driver.Driver,
) Types {
	return &types{}
}

type types struct {
}
