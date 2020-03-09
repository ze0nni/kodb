package engine

func libraryDependency(
	masterLibrary LibraryName,
	masterCol ColumnID,
	lib *libraryImp,
	consumer func(ColumnID),
) error {
	cols, err := Columns(lib)
	if nil != err {
		return err

	}

	for _, c := range cols {
		if c.IsDependent(masterLibrary, masterCol) {
			consumer(c.ID())
		}
	}

	return nil
}
