package driver

import "github.com/ze0nni/kodb/internal/entry"

type (
	// Driver type
	Driver interface {
		Get(prefix string, id string) (entry.Entry, error)
		Put(prefix string, id string, entry entry.Entry) error
	}
)

func copyEntry(e entry.Entry) entry.Entry {
	copy := make(entry.Entry)
	for k, v := range e {
		copy[k] = v
	}
	return copy
}
