package types

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/entry"
)

//TypeName type
type TypeName string

func (tn TypeName) String() string {
	return string(tn)
}

//FieldID type
type FieldID string

func (fid FieldID) String() string {
	return string(fid)
}

type FieldCase string

func (fc FieldCase) String() string {
	return string(fc)
}

//Field type
type Field interface {
	ID() FieldID
	newID(id FieldID)

	Name() string
	Rename(string)
	Kind() FieldDataKind

	Case() FieldCase
	SetCase(FieldCase)

	FillJson(*simplejson.Json)

	fromEntry(entry.Entry) error
	toEntry() entry.Entry
	setListener(l func())
}

//Type type
type Type interface {
	Name() TypeName
	Fields() []Field
	New(Field) (Field, error)
	Get(FieldID) (Field, error)
	Delete(Field) error

	Cases() []FieldCase
	UpdateCases([]FieldCase)

	FillJson(*simplejson.Json)
}

//Types type
type Types interface {
	Names() []TypeName
	New(TypeName) (Type, error)
	Get(TypeName) (Type, error)
	Delete(TypeName) error

	Listen(TypesListener) func()
}

//TypesListener type
type TypesListener interface {
	OnNewType(TypeName)
	OnDeleteType(TypeName)
	OnChangedType(TypeName)

	OnNewField(TypeName, FieldID)
	OnDeleteField(TypeName, FieldID)
	OnChangedField(TypeName, FieldID)
}
