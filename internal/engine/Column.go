package engine

type ColumnType string

const Literal = ColumnType("literal")

const Reference = ColumnType("reference")

func (t ColumnType) ToString() string {
	return string(t)
}
