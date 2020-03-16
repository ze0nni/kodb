package engine

import (
	"fmt"
	"strings"

	"github.com/bitly/go-simplejson"
)

//ColumnMeta type
type ColumnMeta struct {
	data ColumnData
}

// FetchAll fetch full metadata
func (m ColumnMeta) FetchAll(consumer func(string, string)) {
	for k, v := range m.data.entry {
		if strings.HasPrefix(k, "meta:") {
			consumer(k, v)
		}
	}
}

//Update update or delete meta-key
func (m ColumnMeta) Update(key, value string) error {
	if false == strings.HasPrefix(key, "meta:") {
		return fmt.Errorf("<%s> bad mate-key", key)
	}
	if "" == value {
		delete(m.data.entry, key)
	} else {
		m.data.entry[key] = value
	}
	return nil
}

//FillJSON with meta-keys
func (m ColumnMeta) FillJSON(json *simplejson.Json) {
	m.FetchAll(func(k, v string) {
		json.Set(k, v)
	})
}
