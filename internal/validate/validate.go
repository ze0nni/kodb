package validate

import "github.com/ze0nni/kodb/internal/engine"

// Validate returns errors in data
func Validate(
	eng engine.Engine,
	consumer func(engine.LibraryName, engine.RowID, engine.ColumnID, error),
) error {
	for _, libraryName := range eng.Librarys() {
		library, err := eng.Library(libraryName)
		if nil != err {
			return err
		}

		err = libraryValidate(
			eng,
			library,
			func(row engine.RowID, col engine.ColumnID, err error) {
				consumer(libraryName, row, col, err)
			},
		)
		if nil != err {
			return err
		}
	}
	return nil
}

func libraryValidate(
	eng engine.Engine,
	lib engine.Library,
	consumer func(engine.RowID, engine.ColumnID, error),
) error {
	cols, err := engine.Columns(lib)
	if nil != err {
		return err

	}
	rows := lib.Rows()

	for i := 0; i < rows; i++ {
		row, err := lib.RowID(i)
		if nil != err {
			return err
		}
		for _, c := range cols {
			value, _, _ := lib.GetValue(row, c.ID())
			err := c.Validate(eng.Context(), value)
			if nil != err {
				consumer(row, c.ID(), err)
			}
		}
	}

	return nil
}
