package types

import (
	"fmt"

	"github.com/bitly/go-simplejson"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/ze0nni/kodb/internal/driver"
)

func newCommonType(name TypeName, lens driver.Lens, lisrener TypesListener) Type {
	return &commonType{
		name:     name,
		lens:     lens,
		lisrener: lisrener,
		fields:   make(map[FieldID]Field),
		cases:    []FieldCase{""},
	}
}

type commonType struct {
	name     TypeName
	lens     driver.Lens
	lisrener TypesListener
	fields   map[FieldID]Field
	cases    []FieldCase
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
	newField, err := t.register(FieldID(uuid.String()), field)

	if nil != err {
		return nil, err
	}

	t.lisrener.OnNewField(t.name, newField.ID())

	return newField, nil
}

func (t *commonType) register(id FieldID, field Field) (Field, error) {
	if FieldID("") != field.ID() {
		return nil, fmt.Errorf("Field <%s> used", field.ID())
	}

	field.newID(id)

	t.fields[id] = field

	field.setListener(func() {
		t.lisrener.OnChangedField(t.name, id)
	})

	return field, nil
}

func (t *commonType) Get(id FieldID) (Field, error) {
	stored, ok := t.fields[id]
	if false == ok {
		return nil, fmt.Errorf("Field <%s> not registerd", id)
	}
	return stored, nil
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
	field.setListener(nil)

	t.lisrener.OnDeleteField(t.name, field.ID())

	return nil
}

func (t *commonType) FillJson(body *simplejson.Json) {
	body.Set("name", t.name.String())

	cases := make([]string, 0, len(t.cases))
	for _, c := range t.cases {
		cases = append(cases, c.String())
	}
	body.Set("cases", cases)

	fields := simplejson.New()
	for n, f := range t.fields {
		fbody := simplejson.New()
		f.FillJson(fbody)
		fields.Set(n.String(), fbody)
	}

	body.Set("fields", fields)
}

func (t *commonType) Cases() []FieldCase {
	out := make([]FieldCase, 0, len(t.cases))
	return append(out, t.cases...)
}

func (t *commonType) UpdateCases(value []FieldCase) {
	copy := make([]FieldCase, 0, len(value))
	t.cases = append(copy, value...)
}
