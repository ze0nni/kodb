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
}

//Type type
type Type interface {
	Name() TypeName
	Fields() []Field
	New(FieldData) (FieldData, error)
}

//Types type
type Types interface {
	Names() []TypeName
	New(TypeName) (Type, error)
	Get(TypeName) (Type, error)
	Delete(TypeName) error
}
