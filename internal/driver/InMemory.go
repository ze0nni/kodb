package driver

import "github.com/ze0nni/kodb/internal/entry"

func InMemory() Driver {
	return &inMemory{
		data: make(map[string](map[string]entry.Entry)),
	}
}

type inMemory struct {
	data map[string](map[string]entry.Entry)
}

func (d *inMemory) Prefixes() ([]string, error) {
	out := []string{}
	for k, _ := range d.data {
		out = append(out, k)
	}
	return out, nil
}

func (d *inMemory) Get(prefix string, id string) (entry.Entry, error) {
	entrys := d.data[prefix]

	if nil == entrys {
		return nil, nil
	}

	return entrys[id], nil
}

func (d *inMemory) Put(prefix string, id string, e entry.Entry) error {
	entrys := d.data[prefix]

	if nil == entrys {
		entrys = make(map[string]entry.Entry)
		d.data[prefix] = entrys
	}

	entrys[id] = e

	return nil
}
