package engine

import (
	"github.com/bitly/go-simplejson"
)

// ColumnList manage other column rows
type ColumnList struct {
	data ColumnData
}

// Ref library name
func (cl ColumnList) Ref() LibraryName {
	return LibraryName(cl.data.entry["ref"])
}

// UpdateRef change ref library name
func (cl ColumnList) UpdateRef(name LibraryName) {
	cl.data.entry["ref"] = name.ToString()
}
func (cl ColumnList) FillJson(json *simplejson.Json) {
	json.Set("reference", cl.Ref().ToString())
}

func (cl ColumnList) Initilize(
	eng Engine,
) error {
	refLib := eng.GetLibrary(cl.Ref())

	parentCol := ColumnID("parent")
	return refLib.AddColumn(parentCol, "parent") //TODO validate

	return nil
}

func (cl ColumnList) IsDependent(
	library LibraryName,
) bool {
	return library == cl.Ref()
}
