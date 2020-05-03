package validate

import "github.com/ze0nni/kodb/internal/engine"

//  Dependency returns dependency for column
func Dependency(
	masterLibrary engine.LibraryName,
	masterCol engine.FieldID,
	eng engine.Engine,
	consumer func(engine.LibraryName, engine.FieldID),
) error {
	for _, libraryName := range eng.Librarys() {
		library, err := eng.Library(libraryName)
		if nil != err {
			return err
		}

		err = libraryDependency(masterLibrary, masterCol, library, func(col engine.FieldID) {
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
	masterCol engine.FieldID,
	lib engine.Library,
	consumer func(engine.FieldID),
) error {
	// cols, err := engine.Columns(lib)
	// if nil != err {
	// 	return err

	// }

	// for _, c := range cols {
	// 	if c.IsDependent(masterLibrary, masterCol) {
	// 		consumer(c.ID())
	// 	}
	// }

	return nil
}
