package engine

import (
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
