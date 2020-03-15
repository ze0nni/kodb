package web

import (
	"errors"

	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

func msgNewColumnFromJson(
	clientID ClientID,
	json *simplejson.Json,
) (msgNewColumn, error) {
	libraryName := engine.LibraryName(json.Get("library").MustString())
	columnName := json.Get("name").MustString()
	switch engine.ColumnTypeOf(json.Get("type").MustString()) {
	case engine.Literal:
		return msgNewColumn{
			clientID,
			libraryName,
			engine.NewLiteralColumn(columnName),
		}, nil
	case engine.Reference:
		ref := engine.LibraryName(json.Get("ref").MustString())
		return msgNewColumn{
			clientID,
			libraryName,
			engine.NewRefColumn(
				columnName,
				ref,
			),
		}, nil
	case engine.List:
		ref := engine.LibraryName(json.Get("ref").MustString())
		return msgNewColumn{
			clientID,
			libraryName,
			engine.NewListColumn(
				columnName,
				ref,
			),
		}, nil
	}
	return msgNewColumn{}, errors.New("Unknown library type")
}

type msgNewColumn struct {
	clientID    ClientID
	libraryName engine.LibraryName
	columnData  engine.ColumnData
}
