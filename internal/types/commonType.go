package types

func newCommonType(name TypeName) Type {
	return &commonType{
		name:   name,
		fields: make(map[FieldID]Field),
	}
}

type commonType struct {
	name   TypeName
	fields map[FieldID]Field
}

func (t *commonType) Name() TypeName {
	return t.name
}

func (t *commonType) Fields() []Field {
	return []Field{}
}

func (t *commonType) New(field Field) (Field, error) {
	return field, nil
}
