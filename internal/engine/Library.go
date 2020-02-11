package engine

type Library interface {
	Name() string
}

func newLibraryInst(name string) *libraryImp {
	return &libraryImp{
		name: name,
	}
}

type libraryImp struct {
	name string
}

func (self *libraryImp) Name() string {
	return self.name
}
