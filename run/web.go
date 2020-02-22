package run

func Web() error {
	s := newServer()
	return s.run()
}
