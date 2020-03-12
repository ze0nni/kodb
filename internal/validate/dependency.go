package validate

import "github.com/ze0nni/kodb/internal/engine"

//  Dependency returns dependency for column
func Dependency(
	masterLibrary engine.LibraryName,
	masterCol engine.ColumnID,
	eng engine.Engine,
	consumer func(engine.LibraryName, engine.ColumnID),
) error {
	for _, libraryName := range eng.Librarys() {
		library := eng.GetLibrary(libraryName)
		err := libraryDependency(masterLibrary, masterCol, library, func(col engine.ColumnID) {
			consumer(libraryName, col)
		})
		if nil != err {
			return err
		}
	}
	return nil
}

func libraryDependency(
	masterLibrary engine.LibraryName,
	masterCol engine.ColumnID,
	lib engine.Library,
	consumer func(engine.ColumnID),
) error {
	cols, err := engine.Columns(lib)
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