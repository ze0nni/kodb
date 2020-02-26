package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
)

type DeleteRowMsg struct {
	Command string             `json:"command"`
	Library engine.LibraryName `json:"library"`
	RowID   engine.RowID       `json:"rowId"`
}

func DeleteRowMsgOf(
	library engine.LibraryName,
	rowId engine.RowID,
) *DeleteRowMsg {
	return &DeleteRowMsg{
		Command: "deleteRow",
		Library: library,
		RowID:   rowId,
	}
}
