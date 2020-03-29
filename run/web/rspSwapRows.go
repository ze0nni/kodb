package web

import (
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/engine"
)

func rspSwapRows(
	name engine.LibraryName,
	i, j int,
	rowI, rowJ engine.RowID,
) *simplejson.Json {
	resp := simplejson.New()

	resp.Set("command", "swapRows")
	resp.Set("library", name.ToString())
	resp.Set("i", strconv.Itoa(i))
	resp.Set("j", strconv.Itoa(j))
	resp.Set("row", rowI.ToString())
	resp.Set("row0", rowJ.ToString())

	return resp
}
