package engine

func libraryValidate(
	lib *libraryImp,
	consumer func(RowID, ColumnID, error),
) error {
	cols, err := Columns(lib)
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
			err := c.Validate(lib.context, value)
			if nil != err {
				consumer(row, c.ID(), err)
			}
		}
	}

	return nil
}
