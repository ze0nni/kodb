package engine

import "github.com/ze0nni/kodb/internal/entry"

type ColumnType string

const Literal = ColumnType("literal")

const Reference = ColumnType("reference")

const Unknown = ColumnType("unknown")

func (t ColumnType) ToString() string {
	return string(t)
}

type ColumnData struct {
	entry entry.Entry
}

func (d ColumnData) Name() string {
	return d.entry["name"]
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
