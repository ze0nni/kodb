package engine

import (
	"fmt"

	"github.com/bitly/go-simplejson"

	"github.com/ze0nni/kodb/internal/entry"
)

// ColumnType type
type ColumnType string

// Literal value
const Literal = ColumnType("literal")

// Reference to library row or rows
const Reference = ColumnType("reference")

const List = ColumnType("list")

// Unknown column type
const Unknown = ColumnType("unknown")

// ColumnContext type
type ColumnContext interface {
	HasRow(library, row string) (bool, error)
	//GetValue(library, row, col string) (string, bool, error)
}

// ToString func
func (t ColumnType) ToString() string {
	return string(t)
}

// ColumnData type
type ColumnData struct {
	entry entry.Entry
}

// Name of column
func (d ColumnData) Name() string {
	return d.entry["name"]
}

// Rename column
func (d ColumnData) Rename(value string) {
	d.entry["name"] = value
}

// ID of column
func (d ColumnData) ID() ColumnID {
	return ColumnID(d.entry["id"])
}

//NewID copy ColumnData with new ID
func (d ColumnData) NewID(id ColumnID) ColumnData {
	e := d.entry.Copy()
	e["id"] = id.ToString()
	return ColumnData{e}
}

// Type of column
func (d ColumnData) Type() ColumnType {
	t := d.entry["type"]
	switch t {
	case "literal":
		return Literal
	case "reference":
		return Reference
	case "list":
		return List
	default:
		return Unknown
	}
}

// Match column type
func (d ColumnData) Match(
	consumeLiteral func(),
	consumeReference func(ColumnRef),
	consumeList func(ColumnList),
	consumeUnknown func(),
) {
	switch d.Type() {
	case Literal:
		consumeLiteral()
	case Reference:
		ref, err := d.ToRef()
		if nil != err {
			panic(err)
		}
		consumeReference(ref)
	case List:
		list, err := d.ToList()
		if nil != err {
			panic(err)
		}
		consumeList(list)
	default:
		consumeUnknown()
	}
}

func (d ColumnData) FillJson(json *simplejson.Json) {
	d.Match(
		func() {

		},
		func(ref ColumnRef) {
			ref.FillJson(json)
		},
		func(list ColumnList) {
			list.FillJson(json)
		},
		func() {

		},
	)

	json.Set("id", d.ID().ToString())
	json.Set("name", d.Name())
	json.Set("type", d.Type().ToString())
}

func (d ColumnData) IsDependent(
	library LibraryName,
	column ColumnID,
) (out bool) {
	d.Match(
		func() {

		},
		func(ref ColumnRef) {
			out = ref.IsDependent(library)
		},
		func(list ColumnList) {
			out = list.IsDependent(library)
		},
		func() {

		},
	)
	return
}

// Initilize column for first time
func (d ColumnData) Initilize(
	eng Engine,
) (err error) {
	d.Match(
		func() {

		},
		func(ref ColumnRef) {

		},
		func(list ColumnList) {
			err = list.Initilize(eng)
		},
		func() {

		},
	)
	return
}

// Validate cell
func (d ColumnData) Validate(
	context ColumnContext,
	value string,
) error {
	t := d.Type() //TODO: match
	switch t {
	case Literal:
		return nil
	case Reference:
		ref, err := d.ToRef()
		if nil != err {
			return err
		}
		return ref.Validate(context, value)
	default:
		return fmt.Errorf("Unknown type: %s", t)
	}
}

// ToRef convert column to ColumnRef type
func (d ColumnData) ToRef() (ColumnRef, error) {
	if Reference == d.Type() {
		return ColumnRef{d}, nil
	}
	return ColumnRef{ColumnData{nil}}, fmt.Errorf("%s is not Ref %d but", d.Name(), d.Type())
}

// ToList convert column to ColumnList type
func (d ColumnData) ToList() (ColumnList, error) {
	if List == d.Type() {
		return ColumnList{d}, nil
	}
	return ColumnList{ColumnData{nil}}, fmt.Errorf("%s is not List but %s", d.Name(), d.Type())
}
