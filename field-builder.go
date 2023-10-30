package form

import (
	"fmt"
	
	"golang.org/x/exp/constraints"
)

type FieldBuilder struct {
	dataType       string
	fieldType      string
	id             string
	multiple       bool
	valid          bool
	name           string
	value          any
	validators     []validator
	validatorError validatorError
}

type FieldConfig struct {
	fieldType string
	dataType  string
	value     any
}

const (
	fieldTypeButton        = "button"
	fieldTypeCheckbox      = "checkbox"
	fieldTypeColor         = "color"
	fieldTypeDate          = "date"
	fieldTypeDateTimeLocal = "datetime-local"
	fieldTypeEmail         = "email"
	fieldTypeFile          = "file"
	fieldTypeHidden        = "hidden"
	fieldTypeImage         = "image"
	fieldTypeMonth         = "month"
	fieldTypeNumber        = "number"
	fieldTypePassword      = "password"
	fieldTypeRadio         = "radio"
	fieldTypeRange         = "range"
	fieldTypeReset         = "reset"
	fieldTypeSearch        = "search"
	fieldTypeSubmit        = "submit"
	fieldTypeTel           = "tel"
	fieldTypeText          = "text"
	fieldTypeTime          = "time"
	fieldTypeUrl           = "url"
	fieldTypeWeek          = "week"
	
	fieldDataTypeBool   = "bool"
	fieldDataTypeFile   = "file"
	fieldDataTypeFloat  = "float"
	fieldDataTypeInt    = "int"
	fieldDataTypeString = "string"
)

func Add(name string) *FieldBuilder {
	return &FieldBuilder{
		name:       name,
		validators: make([]validator, 0),
	}
}

func (b *FieldBuilder) With(config FieldConfig, validators ...Validator) *FieldBuilder {
	switch config.value.(type) {
	case []string:
		createFieldType[string](b, config.fieldType, config.dataType, config.value.([]string)...)
	case []int:
		createFieldType[int](b, config.fieldType, config.dataType, config.value.([]int)...)
	case []float32:
		createFieldType[float32](b, config.fieldType, config.dataType, config.value.([]float32)...)
	case []float64:
		createFieldType[float64](b, config.fieldType, config.dataType, config.value.([]float64)...)
	case []bool:
		createFieldType[bool](b, config.fieldType, config.dataType, config.value.([]bool)...)
	case []Multipart:
		createFieldType[Multipart](b, config.fieldType, config.dataType, config.value.([]Multipart)...)
	}
	for _, v := range validators {
		b.validators = append(b.validators, v.(validator))
	}
	return b
}

func Button(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeButton,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Checkbox(value ...bool) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeCheckbox,
		dataType:  fieldDataTypeBool,
		value:     value,
	}
}

func Color(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeColor,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Date(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeDate,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func DateTimeLocal(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeDateTimeLocal,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Email(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeEmail,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func File(value ...Multipart) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeFile,
		dataType:  fieldDataTypeFile,
		value:     value,
	}
}

func Hidden(value ...any) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeHidden,
		dataType:  fieldDataTypeString,
		value: convertSlice[any, string](
			value, func(v any) string {
				return fmt.Sprintf("%v", v)
			},
		),
	}
}

func Image(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeImage,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func (b *FieldBuilder) Id(id string) *FieldBuilder {
	b.id = id
	return b
}

func Month(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeMonth,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func (b *FieldBuilder) Multiple(multiple ...bool) *FieldBuilder {
	b.multiple = true
	if len(multiple) > 0 {
		b.multiple = multiple[0]
	}
	return b
}

func Number[T constraints.Float | constraints.Integer](value ...T) FieldConfig {
	v := *new(T)
	switch any(v).(type) {
	case int:
		return FieldConfig{
			fieldType: fieldTypeNumber,
			dataType:  fieldDataTypeInt,
			value:     value,
		}
	}
	return FieldConfig{
		fieldType: fieldTypeNumber,
		dataType:  fieldDataTypeFloat,
		value:     value,
	}
}

func Password(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypePassword,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Radio(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeRadio,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Range(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeRange,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Reset(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeReset,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Search(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeSearch,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Submit(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeSubmit,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Tel(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeTel,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Text(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeText,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Time(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeTime,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Url(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeUrl,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Week(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeWeek,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func createFieldType[T any](b *FieldBuilder, fieldType, dataType string, values ...T) {
	b.fieldType = fieldType
	b.dataType = dataType
	b.multiple = len(values) > 1
	if b.multiple {
		b.value = values
	}
	if !b.multiple && len(values) > 0 {
		b.value = values[0]
	}
	if b.value == nil {
		if b.multiple {
			b.value = make([]T, 0)
		}
		if !b.multiple {
			b.value = *new(T)
		}
	}
}
