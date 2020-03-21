package types

import (
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/ze0nni/kodb/internal/entry"
)

// FieldDataKind type
type FieldDataKind string

func (k FieldDataKind) String() string {
	return string(k)
}

const (
	// ValueFieldKind field kind
	ValueFieldKind = FieldDataKind("value")

	// ReferenceFieldKind field kind
	ReferenceFieldKind = FieldDataKind("reference")

	// ListFieldKind field kind
	ListFieldKind = FieldDataKind("list")

	// InstanceFieldKind field kind
	InstanceFieldKind = FieldDataKind("instance")

	// UnknownFieldKind field kind
	UnknownFieldKind = FieldDataKind("unknown")
)

type fieldData struct {
	id       FieldID
	name     string
	kind     FieldDataKind
	fcase    FieldCase
	listener func()
}

func (fd fieldData) ID() FieldID {
	return fd.id
}

func (fd *fieldData) newID(id FieldID) {
	fd.id = id
}

func (fd fieldData) Name() string {
	return fd.name
}

func (fd *fieldData) Rename(name string) {
	fd.name = name

	fd.onChanged()
}

func (fd fieldData) Kind() FieldDataKind {
	return fd.kind
}

func (fd fieldData) Case() FieldCase {
	return fd.fcase
}

func (fd *fieldData) SetCase(value FieldCase) {
	fd.fcase = value

	fd.onChanged()
}

func (fd fieldData) fillJson(body *simplejson.Json) {
	body.Set("id", fd.id.String())
	body.Set("kind", fd.kind.String())
	body.Set("name", fd.name)
	body.Set("case", fd.fcase.String())
}

func (fd *fieldData) readEntry(e entry.Entry) error {
	id := FieldID(e["id"])
	if id != fd.id {
		return fmt.Errorf("IDs <%s> and <%s> not match", id, fd.id)
	}

	kind := FieldDataKind(e["kind"])
	if kind != fd.kind {
		return fmt.Errorf("Kinds <%s> and <%s> not match", kind, fd.kind)
	}

	fd.name = e["name"]
	return nil
}

func (fd fieldData) createEntry() entry.Entry {
	e := make(entry.Entry)

	e["id"] = fd.id.String()
	e["kind"] = fd.kind.String()
	e["name"] = fd.name

	return e
}

func (fd *fieldData) setListener(l func()) {
	fd.listener = l
}

func (fd *fieldData) onChanged() {
	if nil != fd.listener {
		fd.listener()
	}
}

//ReferenceFieldData type
type ReferenceFieldData struct {
	fieldData
}

//ListFieldData type
type ListFieldData struct {
	fieldData
}
