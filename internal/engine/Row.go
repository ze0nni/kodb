package engine

type (
	// RowID type
	RowID string

	// Row type
	Row interface {
		Exists() (bool, error)

		// Get(ColumnID) (string, bool, error)
		// Has(ColumnID) bool
		// Set(ColumnID, string) error
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

func (r *rowImpl) private() {

}
