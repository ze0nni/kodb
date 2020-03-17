package types

type Type interface {
	Name() string
}

type Types interface {
	Names() []string
	New(string) (Type, error)
	Get(string) (Type, error)
	Delete(string) error
}
