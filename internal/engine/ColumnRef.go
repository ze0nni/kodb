package engine

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

// ColumnRef reference to one or many rows from other library
type ColumnRef struct {
	data ColumnData
}

// Data return ColumnData
func (r ColumnRef) Data() ColumnData {
	return r.data
}

// Ref library name
func (r ColumnRef) Ref() LibraryName {
	return LibraryName(r.data.entry["ref"])
}

// UpdateRef change ref library name
func (r ColumnRef) UpdateRef(name LibraryName) {
	r.data.entry["ref"] = name.ToString()
}

func (r ColumnRef) FillJson(json *simplejson.Json) {
	json.Set("reference", r.Ref().ToString())
}

func (r ColumnRef) IsDependent(
	library LibraryName,
) bool {
	return library == r.Ref()
}

// Validate cell
func (r ColumnRef) Validate(
	context ColumnContext,
	value string,
) error {
	exists, err := context.HasRow(r.Ref().ToString(), value)
	if nil != err {
		return err
	}
	if false == exists {
		return fmt.Errorf("Row <%s> not exists in <%s>", value, r.Ref().ToString())
	}
	return nil
}
