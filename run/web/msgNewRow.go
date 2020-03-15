package web

import (
	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

func msgNewRowFromJson(
	clientID ClientID,
	json *simplejson.Json,
) (msgNewRow, error) {
	libraryName := json.Get("library").MustString()
	if data, ok := json.CheckGet("parentLibrary"); ok {
		parentLibrary := engine.LibraryName(data.MustString())
		parentRow := engine.RowID(json.Get("parentRow").MustString())
		parentColumn := engine.ColumnID(json.Get("parentColumn").MustString())
		return msgNewRow{
			ClientID:          clientID,
			LibraryName:       engine.LibraryName(libraryName),
			HasParent:         true,
			ParentLibraryName: parentLibrary,
			ParentRowID:       parentRow,
			ParentColumnID:    parentColumn,
		}, nil
	}
	return msgNewRow{
		ClientID:    clientID,
		LibraryName: engine.LibraryName(libraryName),
		HasParent:   false,
	}, nil
}

type msgNewRow = struct {
	ClientID          ClientID
	LibraryName       engine.LibraryName
	HasParent         bool
	ParentLibraryName engine.LibraryName
	ParentRowID       engine.RowID
	ParentColumnID    engine.ColumnID
}
