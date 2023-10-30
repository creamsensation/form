package form

import (
	"fmt"
	"reflect"
	
	"github.com/iancoleman/strcase"
)

const (
	baseFormFieldName = "Form"
)

func Build[T any](b *FormBuilder) (T, error) {
	if b.request == nil {
		return buildForm[T](b), nil
	}
	b.submitted = isFormSubmitted(b.request)
	reqFormData, reqFormFiles, err := processRequest(b.request, b.limit)
	if err != nil {
		return *new(T), fmt.Errorf("error processing request to form: %w", err)
	}
	if len(reqFormData) > 0 {
		processFormData(b, reqFormData)
	}
	if len(reqFormFiles) > 0 {
		err := processFormFiles(b, reqFormFiles)
		if err != nil {
			return *new(T), err
		}
	}
	return buildForm[T](b), nil
}

func buildForm[T any](b *FormBuilder) T {
	form := new(T)
	formRef := reflect.ValueOf(form)
	for i, fb := range b.fields {
		fb.validatorError = b.validatorError
		errors := buildFormField(formRef, fb)
		b.fields[i].valid = len(errors) == 0
	}
	buildBaseForm(formRef, b)
	return *form
}

func buildFormField(formRef reflect.Value, fb *FieldBuilder) []error {
	errors := make([]error, 0)
	formField := formRef.Elem().FieldByName(strcase.ToCamel(fb.name))
	if !formField.IsValid() {
		return errors
	}
	switch fb.dataType {
	case fieldDataTypeString:
		if fb.multiple {
			field := createFormField[[]string](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
		if !fb.multiple {
			field := createFormField[string](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
	case fieldDataTypeFloat:
		if fb.multiple {
			field := createFormField[[]float64](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
		if !fb.multiple {
			field := createFormField[float64](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
	case fieldDataTypeInt:
		if fb.multiple {
			field := createFormField[[]int](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
		if !fb.multiple {
			field := createFormField[int](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
	case fieldDataTypeBool:
		if fb.multiple {
			field := createFormField[[]bool](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
		if !fb.multiple {
			field := createFormField[bool](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
	case fieldDataTypeFile:
		if fb.multiple {
			field := createFormField[[]Multipart](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
		if !fb.multiple {
			field := createFormField[Multipart](fb)
			formField.Set(reflect.ValueOf(field))
			return field.Errors
		}
	}
	return errors
}

func buildBaseForm(formRef reflect.Value, b *FormBuilder) {
	if formRef.Kind() != reflect.Ptr {
		return
	}
	baseFormField := formRef.Elem().FieldByName(baseFormFieldName)
	if !baseFormField.IsValid() {
		return
	}
	baseFormField.Set(reflect.ValueOf(createBaseForm(b)))
}

func createBaseForm(b *FormBuilder) Form {
	var csrfName, csrfToken string
	if b.request != nil {
		csrfName = b.request.FormValue(CsrfName)
		csrfToken = b.request.FormValue(CsrfToken)
	}
	return Form{
		Method:      b.method,
		Action:      b.action,
		IsValid:     b.isValid(),
		IsSubmitted: b.submitted,
		CsrfName:    csrfName,
		CsrfToken:   csrfToken,
	}
}

func createFormField[T any](fb *FieldBuilder) Field[T] {
	return Field[T]{
		Id:       fb.id,
		Name:     fb.name,
		Type:     fb.fieldType,
		DataType: fb.dataType,
		Value:    fb.value.(T),
		Multiple: fb.multiple,
		Errors:   validateField(fb),
	}
}
