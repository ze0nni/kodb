package engine

import (
	"fmt"

	"github.com/ze0nni/kodb/internal/entry"
)

type ColumnType string

const Literal = ColumnType("literal")

const Reference = ColumnType("reference")

const Unknown = ColumnType("unknown")

type ColumnContext interface {
	//GetValue(library, row, col string) (string, bool, error)
}

func (t ColumnType) ToString() string {
	return string(t)
}

type ColumnData struct {
	entry entry.Entry
}

func (d ColumnData) Name() string {
	return d.entry["name"]
}

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
