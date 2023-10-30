package form

import (
	"net/http"
)

type FormBuilder struct {
	fields         []*FieldBuilder
	request        *http.Request
	method         string
	action         string
	name           string
	limit          int
	submitted      bool
	validatorError validatorError
}

const (
	defaultBodyLimit = 256
)

func New(fields ...*FieldBuilder) *FormBuilder {
	return &FormBuilder{
		fields:         fields,
		validatorError: createDefaultErrors(),
		limit:          defaultBodyLimit,
	}
}

func (b *FormBuilder) Action(action string) *FormBuilder {
	b.action = action
	return b
}

func (b *FormBuilder) Add(name string) *FieldBuilder {
	return Add(name)
}

func (b *FormBuilder) Get(name string) *FieldBuilder {
	for _, f := range b.fields {
		if f.name != name {
			continue
		}
		return f
	}
	return nil
}

func (b *FormBuilder) Limit(limit int) *FormBuilder {
	b.limit = limit
	return b
}
func (b *FormBuilder) Method(method string) *FormBuilder {
	b.method = method
	return b
}

func (b *FormBuilder) Name(name string) *FormBuilder {
	b.name = name
	return b
}

func (b *FormBuilder) Request(request *http.Request) *FormBuilder {
	b.request = request
	return b
}

func (b *FormBuilder) isValid() bool {
	if !b.submitted {
		return true
	}
	for _, field := range b.fields {
		if !field.valid {
			return false
		}
	}
	return true
}
