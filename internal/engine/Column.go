package engine

import (
	"fmt"

	"github.com/ze0nni/kodb/internal/entry"
)

// ColumnType type
type ColumnType string

// Literal value
const Literal = ColumnType("literal")

// Reference to library row or rows
const Reference = ColumnType("reference")

// Unknown column type
const Unknown = ColumnType("unknown")

// ColumnContext type
type ColumnContext interface {
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

func (d ColumnData) Type() ColumnType {
	t := d.entry["type"]
	switch t {
	case "literal":
		return Literal
	case "reference":
		return Reference
	default:
		return Unknown
	}
}

func (d ColumnData) FillJson(json *simplejson.Json) {
	json.Set("id", d.ID().ToString())
	json.Set("name", d.Name())
	json.Set("type", d.Type().ToString())
}

// Validate cell
func (d ColumnData) Validate(
	context ColumnContext,
	value string,
) error {
	t := d.Type()
	switch t {
	case Literal:
		return nil
	case Reference:
		return fmt.Errorf("Unknown refer: %s", value)
	default:
		return fmt.Errorf("Unknown type: %s", t)
	}
}
