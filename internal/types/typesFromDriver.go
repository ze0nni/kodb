package types

import (
	"fmt"

	"github.com/ze0nni/kodb/internal/driver"
)

const typePrefix = "type_"

func typesOfDriver(
	driver driver.Driver,
) Types {
	return &types{
		dict: make(map[string]Type),
	}
}

type types struct {
	dict map[string]Type
}

func (ts *types) Names() []string {
	out := []string{}

	for k := range ts.dict {
		out = append(out, k)
	}

	return out
}

func (ts *types) New(name string) (Type, error) {
	if _, ok := ts.dict[name]; ok {
		return nil, fmt.Errorf("Duplicate type <%s>", name)
	}

	t := newCommonType(name)
	ts.dict[name] = t
	return t, nil
}

func (ts *types) Get(name string) (Type, error) {
	if t, ok := ts.dict[name]; ok {
		return t, nil
	}
	return nil, fmt.Errorf("Type <%s> not exists", name)
}

func (ts *types) Delete(name string) error {
	if _, ok := ts.dict[name]; ok {
		delete(ts.dict, name)
		return nil
	}
	return fmt.Errorf("Type <%s> not exists", name)
}
