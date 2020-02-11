package driver

type (
	// Entry type
	Entry = map[string]string

	// Driver type
	Driver interface {
		Get(prefix string, id string) (Entry, error)
		Put(prefix string, id string, entry Entry) error
	}
)

func copyEntry(e Entry) Entry {
	copy := make(Entry)
	for k, v := range e {
		copy[k] = v
	}
	return copy
}
