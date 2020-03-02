package engine

import "github.com/ze0nni/kodb/internal/entry"

type ColumnType string

const Literal = ColumnType("literal")

const Reference = ColumnType("reference")

const Unknown = ColumnType("unknown")

func (t ColumnType) ToString() string {
	return string(t)
}

type ColumnData entry.Entry

func (d ColumnData) Name() string {
	e := entry.Entry(d)
	return e["name"]
}

func (d ColumnData) Type() ColumnType {
	e := entry.Entry(d)
	t := e["type"]
	switch t {
	case "literal":
		return Literal
	case "reference":
		return Reference
	default:
		return Unknown
	}
}
