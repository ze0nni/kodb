package types

// FieldDataKind type
type FieldDataKind string

const (
	// ValueFieldKind field kind
	ValueFieldKind = FieldDataKind("value")

	// ReferenceFieldKind field kind
	ReferenceFieldKind = FieldDataKind("reference")

	// ListFieldKind field kind
	ListFieldKind = FieldDataKind("list")

	// UnknownFieldKind field kind
	UnknownFieldKind = FieldDataKind("unknown")
)

type fieldData struct {
	id   FieldID
	name string
	kind FieldDataKind
}

func (fd fieldData) ID() FieldID {
	return fd.id
}

func (fd *fieldData) newID(id FieldID) {
	fd.id = id
}

func (fd fieldData) Name() string {
	return fd.name
}

func (fd fieldData) Kind() FieldDataKind {
	return fd.kind
}

//NewValueFieldData returns ValueFieldData
func NewValueFieldData(name string) *ValueFieldData {
	return &ValueFieldData{
		fieldData{
			name: name,
			kind: ValueFieldKind,
		},
	}
}

//ValueFieldData type
type ValueFieldData struct {
	fieldData
}

func (vfd *ValueFieldData) private() {

}

//ReferenceFieldData type
type ReferenceFieldData struct {
	fieldData
}

//ListFieldData type
type ListFieldData struct {
	fieldData
}
