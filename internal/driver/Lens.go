package driver

import (
	"github.com/ze0nni/kodb/internal/entry"
)

// Lens type
type Lens interface {
	Keys() ([]string, error)
	Get(id string) (entry.Entry, error)
	Put(id string, entry entry.Entry) error
	Delete(id string) error
}

// LensOf make 'DriverLens' from 'Driver'
func LensOf(prefix string, driver Driver) Lens {
	return &driverLens{
		prefix: prefix,
		driver: driver,
	}
}

type driverLens struct {
	prefix string
	driver Driver
}

func (lens *driverLens) Keys() ([]string, error) {
	return lens.driver.IDs(lens.prefix)
}

func (lens *driverLens) Get(id string) (entry.Entry, error) {
	return lens.driver.Get(lens.prefix, id)
}

func (lens *driverLens) Put(id string, entry entry.Entry) error {
	return lens.driver.Put(lens.prefix, id, entry)
}

func (lens *driverLens) Delete(id string) error {
	return lens.driver.Delete(lens.prefix, id)
}

// UpdateLens entry
func UpdateLens(l Lens, id string, consumer func(entry.Entry)) error {
	e, err := l.Get(id)
	if nil != err {
		return err
	}
	consumer(e)

	return l.Put(id, e)
}
