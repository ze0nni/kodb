package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
)

type SetSchemaMsg struct {
	Command  string           `json:"command"`
	Librarys []*LibrarySchema `json:"librarys"`
}

func SetSchemaMsgFromEngine(engine engine.Engine) *SetSchemaMsg {
	msg := &SetSchemaMsg{
		Command:  "setSchema",
		Librarys: []*LibrarySchema{},
	}

	for _, libraryName := range engine.Librarys() {
		msg.Librarys = append(
			msg.Librarys,
			NewLibrarySchemaFromEngine(libraryName, engine),
		)
	}

	return msg
}
