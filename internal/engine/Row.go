package engine

import (
	"errors"
)

type (
	// RowID type
	RowID string

	// Row type
	Row interface {
		Exists() (bool, error)

		Get(ColumnID) (string, bool, error)
		// Has(ColumnID) (bool, error)
		Set(ColumnID, string) error
		// Delete(ColumnID) error

		private()
	}
)

func (id RowID) ToString() string {
	return string(id)
}

func RowOf(
	data Lens,
	id RowID,
) Row {
	return &rowImpl{
		data: data,
		id:   id,
	}
}

type rowImpl struct {
	data Lens
	id   RowID
}

func (r *rowImpl) Exists() (bool, error) {
	e, err := r.data.Get(r.id.ToString())
	if nil != err {
		return false, err
	}
	return nil != e, nil
}

func (r *rowImpl) Get(id ColumnID) (string, bool, error) {
	e, err := r.data.Get(r.id.ToString())
	if nil != err {
		return "", false, err
	}
	if nil == e {
		return "", false, errors.New("Row not exists:" + r.id.ToString())
	}
	v, ok := e[id.ToString()]
	if false == ok {
		return "", false, nil
	}
	return v, true, nil
}

func (r *rowImpl) Set(id ColumnID, value string) error {
	e, err := r.data.Get(r.id.ToString())
	if nil != err {
		return err
	}
	if nil == e {
		return errors.New("Row not exists:" + r.id.ToString())
	}
	e[id.ToString()] = value
	return r.data.Put(r.id.ToString(), e)
}

func (r *rowImpl) private() {

}
