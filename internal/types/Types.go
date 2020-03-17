package types

type Type interface {
	Name() string
}

type Types interface {
	Names() ([]string, error)
	Add(name string) (Type, error)
}
