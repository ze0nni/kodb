package engine

func newLogListener() *logListener {
	return &logListener{
		log: []string{},
	}
}

type logListener struct {
	log []string
}

func (l *logListener) getLog() []string {
	return l.log
}

func (l *logListener) NewLibrary(name LibraryName) {
	l.log = append(l.log, "newLibrary "+name.ToString())
}
