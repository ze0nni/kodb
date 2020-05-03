package msg

import (
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
)

func UpdateValueMsgOf(
	library engine.LibraryName,
	rowID engine.RowID,
	fieldID types.FieldID,
	exists bool,
	value string,
	cellErr error,
) *UpdateValueMsg {
	var cellErrorStr *string
	if nil != cellErr {
		str := cellErr.Error()
		cellErrorStr = &str
	}

	return &UpdateValueMsg{
		Command: "updateValue",
		Library: library,
		RowID:   rowID,
		FieldID: fieldID,
		Exists:  exists,
		Value:   value,
		Error:   cellErrorStr,
	}
}

type UpdateValueMsg struct {
	Command string             `json:"command"`
	Library engine.LibraryName `json:"library"`
	RowID   engine.RowID       `json:"rowId"`
	FieldID engine.FieldID     `json:"fieldId"`
	Exists  bool               `json:"exists"`
	Value   string             `json:"value"`
	Error   *string            `json:"error"`
}
