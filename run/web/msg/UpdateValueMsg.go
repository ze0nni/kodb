package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
)

func UpdateValueMsgOf(
	library engine.LibraryName,
	rowId engine.RowID,
	columnId engine.ColumnID,
	exists bool,
	value string,
) *UpdateValueMsg {
	return &UpdateValueMsg{
		Command:  "updateValue",
		Library:  library,
		RowID:    rowId,
		ColumnID: columnId,
		Exists:   exists,
		Value:    value,
	}
}

type UpdateValueMsg struct {
	Command  string             `json:"command"`
	Library  engine.LibraryName `json:"library"`
	RowID    engine.RowID       `json:"rowId"`
	ColumnID engine.ColumnID    `json:"columnId"`
	Exists   bool               `json:"exists"`
	Value    string             `json:"value"`
}
