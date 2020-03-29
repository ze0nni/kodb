package engine

import (
	"fmt"
	"strconv"

	"github.com/ze0nni/kodb/internal/driver"
)

type rowsByOrder struct {
	data driver.Lens
	rows []RowID
}

func (o *rowsByOrder) Len() int {
	return len(o.rows)
}

func (o *rowsByOrder) Less(i, j int) bool {
	ie, err := o.data.Get(o.rows[i].ToString())
	if nil != err {
		panic(err)
	}
	iOrd, err := strconv.Atoi(ie["order"])

	je, err := o.data.Get(o.rows[j].ToString())
	if nil != err {
		panic(err)
	}
	jOrd, err := strconv.Atoi(je["order"])

	return iOrd < jOrd
}

func (o *rowsByOrder) Swap(i, j int) {
	o.rows[i], o.rows[j] = o.rows[j], o.rows[i]
}

func SwapRowsByID(library Library, r, r0 RowID) error {
	i, ok := library.RowIndex(r)
	if false == ok {
		return fmt.Errorf("Row <%s> not found in <%s>", r.ToString(), library.Name().ToString())
	}
	j, ok := library.RowIndex(r0)
	if false == ok {
		return fmt.Errorf("Row <%s> not found in <%s>", r0.ToString(), library.Name().ToString())
	}

	return library.Swap(i, j)
}
