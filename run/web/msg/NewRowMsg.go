package msg

import "github.com/ze0nni/kodb/internal/engine"

type NewRowMsg struct {
	Command string             `json:"command"`
	Library engine.LibraryName `json:"library"`
	RowID   engine.RowID       `json:"rowId"`
}

func NewRowMsgOf(
	library engine.LibraryName,
	rowId engine.RowID,
) *NewRowMsg {
	return &NewRowMsg{
		Command: "newRow",
		Library: library,
		RowID:   rowId,
	}
}
