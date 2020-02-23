package msg

import "github.com/ze0nni/kodb/internal/engine"

type SetLibraryRowsMsg struct {
	Command string             `json:"command"`
	Library engine.LibraryName `json:"library"`
	Rows    []*struct{}        `json:"rows"`
}

func SetLibraryRowsMsgFromEngine(
	name engine.LibraryName,
	engine engine.Engine,
) *SetLibraryRowsMsg {
	l := engine.GetLibrary(name)

	msg := &SetLibraryRowsMsg{
		Command: "setLibraryRows",
		Library: name,
		Rows:    []*struct{}{},
	}

	rows := l.Rows()
	for i := 0; i < rows; i++ {
		msg.Rows = append(
			msg.Rows,
			&struct{}{},
		)
	}

	return msg
}
