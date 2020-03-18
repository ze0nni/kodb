package types

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

//Field type
type Field interface {
	ID() FieldID
	newID(id FieldID)

	Name() string
	Kind() FieldDataKind

	private()
}

//Type type
type Type interface {
	Name() TypeName
	Fields() []Field
	New(Field) (Field, error)
}

//Types type
type Types interface {
	Names() []TypeName
	New(TypeName) (Type, error)
	Get(TypeName) (Type, error)
	Delete(TypeName) error
}
