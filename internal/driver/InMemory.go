package driver

func InMemory() Driver {
	return &inMemory{
		data: make(map[string]map[string]Entry),
	}
}

type inMemory struct {
	data map[string]map[string]Entry
}

func (d *inMemory) Get(prefix string, id string) (error, Entry) {
	entrys := d.data[prefix]

	if nil == entrys {
		return nil, nil
	}

	return nil, entrys[id]
}

func (d *inMemory) Put(prefix string, id string, entry Entry) error {
	entrys := d.data[prefix]

	if nil == entrys {
		entrys = make(map[string]Entry)
		d.data[prefix] = entrys
	}

	entrys[id] = entry

	return nil
}
