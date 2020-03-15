package engine

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/entry"
)

func NewListColumn(name string, ref LibraryName) ColumnData {
	e := make(entry.Entry)
	e["name"] = name
	e["type"] = List.ToString()
	e["ref"] = ref.ToString()
	return ColumnData{e}
}

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
	refLib, err := eng.AddLibrary(cl.Ref())
	if nil != err {
		return err
	}

	parentCol := ColumnID("parent")
	return refLib.AddColumn(parentCol, "parent") //TODO validate

	return nil
}

func (cl ColumnList) IsDependent(
	library LibraryName,
) bool {
	return library == cl.Ref()
}
