package types

func newCommonType(name string) Type {
	return &commonType{
		name: name,
	}
}

type commonType struct {
	name string
}

func (t *commonType) Name() string {
	return t.name
}
