package types

import (
	"fmt"

	"github.com/bitly/go-simplejson"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/ze0nni/kodb/internal/driver"
)

func newCommonType(name TypeName, lens driver.Lens) Type {
	return &commonType{
		name:   name,
		lens:   lens,
		fields: make(map[FieldID]Field),
	}
}

type commonType struct {
	name   TypeName
	lens   driver.Lens
	fields map[FieldID]Field
}

func (t *commonType) Name() TypeName {
	return t.name
}

func (t *commonType) Fields() []Field {
	out := []Field{}

	for _, f := range t.fields {
		out = append(out, f)
	}

	return out
}

func (t *commonType) New(field Field) (Field, error) {
	uuid, err := uuid.NewV4()
	if nil != err {
		return nil, err
	}
	return t.register(FieldID(uuid.String()), field)
}

func (t *commonType) register(id FieldID, field Field) (Field, error) {
	if FieldID("") != field.ID() {
		return nil, fmt.Errorf("Field <%s> used", field.ID())
	}

	field.newID(id)

	t.fields[id] = field

	return field, nil
}

func (t *commonType) Delete(field Field) error {
	stored, ok := t.fields[field.ID()]

	if false == ok {
		return fmt.Errorf("Field <%s> not registerd", field.ID())
	}

	if stored != field {
		return fmt.Errorf("Fields <%s> and <%s> not match", stored, field)
	}

	delete(t.fields, field.ID())

	return nil
}

func (t *commonType) FillJson(body *simplejson.Json) {
	body.Set("name", t.name.String())

	fields := simplejson.New()
	for n, f := range t.fields {
		fbody := simplejson.New()
		f.FillJson(fbody)
		fields.Set(n.String(), fbody)
	}

	body.Set("fields", fields)
}
