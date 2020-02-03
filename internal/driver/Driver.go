package driver

type Driver interface {
	Get(prefix string, id string) (error, Entry)
	Put(prefix string, id string, entry Entry) error
}

type Entry = map[string]string

func CopyEntry(e Entry) Entry {
	copy := make(Entry)
	for k, v := range e {
		copy[k] = v
	}
	return copy
}
