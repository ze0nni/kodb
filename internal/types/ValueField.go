package types

import "github.com/ze0nni/kodb/internal/entry"

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

func (vfd *ValueFieldData) fromEntry(e entry.Entry) error {
	err := vfd.readEntry(e)
	if nil != err {
		return err
	}

	return nil
}

func (vfd *ValueFieldData) toEntry() entry.Entry {
	return vfd.createEntry()
}
