package types

import (
	"fmt"
	"strings"

	"github.com/ze0nni/kodb/internal/entry"

	"github.com/ze0nni/kodb/internal/driver"
)

const typePrefix = "type_"

func typesOfDriver(
	driver driver.Driver,
) (Types, error) {
	types := &types{
		driver: driver,
		dict:   make(map[TypeName]Type),
	}

	ps, err := driver.Prefixes()
	if nil != err {
		return nil, err
	}

	for _, p := range ps {
		if strings.HasPrefix(p, typePrefix) {
			name := TypeName(p[len(typePrefix):])
			types.dict[name] = newCommonType(
				name,
			)
		}
	}

	return types, nil
}

type types struct {
	driver driver.Driver
	dict   map[TypeName]Type
}

func (ts *types) Names() []TypeName {
	out := []TypeName{}

	for k := range ts.dict {
		out = append(out, k)
	}

	return out
}

func (ts *types) New(name TypeName) (Type, error) {
	if _, ok := ts.dict[name]; ok {
		return nil, fmt.Errorf("Duplicate type <%s>", name)
	}

	ts.driver.Put(typePrefix+name.String(), "root", make(entry.Entry))

	t := newCommonType(name)
	ts.dict[name] = t
	return t, nil
}

func (ts *types) Get(name TypeName) (Type, error) {
	if t, ok := ts.dict[name]; ok {
		return t, nil
	}
	return nil, fmt.Errorf("Type <%s> not exists", name)
}

func (ts *types) Delete(name TypeName) error {

	if _, ok := ts.dict[name]; ok {
		err := ts.driver.DeletePrefix(typePrefix + name.String())
		if nil != err {
			return err
		}

		delete(ts.dict, name)
		return nil
	}
	return fmt.Errorf("Type <%s> not exists", name)
}
