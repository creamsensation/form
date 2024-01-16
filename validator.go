package form

import (
	"net/http"
	"regexp"
)

type Validator interface{}

type Validators struct{}

type validator struct {
	validatorType int
	value         any
	pattern       string
}

const (
	validatorEmail = "[-A-Za-z0-9!#$%&'*+/=?^_`{|}~]+(?:\\.[-A-Za-z0-9!#$%&'*+/=?^_`{|}~]+)*@(?:[A-Za-z0-9](?:[-A-Za-z0-9]*[A-Za-z0-9])?\\.)+[A-Za-z0-9](?:[-A-Za-z0-9]*[A-Za-z0-9])?"
)

const (
	validatorTypeRequired = iota
	validatorTypeMin
	validatorTypeMax
	validatorTypeEmail
	validatorTypeCustom
)

func CreateValidator[T any](pattern string) func(value ...T) Validator {
	return func(value ...T) Validator {
		v := *new(T)
		if len(value) > 0 {
			v = value[0]
		}
		return validator{
			validatorType: validatorTypeCustom,
			pattern:       pattern,
			value:         v,
		}
	}
}

var Validate = Validators{}

func (v Validators) Required() Validator {
	return validator{
		validatorType: validatorTypeRequired,
	}
}

func (v Validators) Email() Validator {
	return validator{
		validatorType: validatorTypeEmail,
		pattern:       validatorEmail,
	}
}

func (v Validators) Min(value int) Validator {
	return validator{
		validatorType: validatorTypeMin,
		value:         value,
	}
}

func (v Validators) Max(value int) Validator {
	return validator{
		validatorType: validatorTypeMax,
		value:         value,
	}
}

func validateField(fb *FieldBuilder, req *http.Request) []error {
	errors := make([]error, 0)
	if req != nil && req.Method == http.MethodGet {
		return errors
	}
	for _, v := range fb.validators {
		switch v.validatorType {
		case validatorTypeRequired:
			errors = append(errors, validateRequired(fb)...)
		case validatorTypeMin:
			errors = append(errors, validateMin(fb, v)...)
		case validatorTypeMax:
			errors = append(errors, validateMax(fb, v)...)
		case validatorTypeEmail:
			errors = append(errors, validateEmail(fb, v)...)
		case validatorTypeCustom:
			errors = append(errors, validateCustom(fb, v)...)
		}
	}
	return errors
}

func validateRequired(fb *FieldBuilder) []error {
	errors := make([]error, 0)
	switch fv := fb.value.(type) {
	
	case []string:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
		if len(fv) > 0 {
			for _, item := range fv {
				if len(item) == 0 {
					errors = append(errors, fb.validatorError[ErrorRequired])
					break
				}
			}
		}
	case []int:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
		if len(fv) > 0 {
			for _, item := range fv {
				if item < 1 {
					errors = append(errors, fb.validatorError[ErrorRequired])
					break
				}
			}
		}
	case []float64:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
		if len(fv) > 0 {
			for _, item := range fv {
				if item < 0.01 {
					errors = append(errors, fb.validatorError[ErrorRequired])
					break
				}
			}
		}
	case []float32:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
		if len(fv) > 0 {
			for _, item := range fv {
				if item < 0.01 {
					errors = append(errors, fb.validatorError[ErrorRequired])
					break
				}
			}
		}
	case []bool:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
		if len(fv) > 0 {
			for _, item := range fv {
				if !item {
					errors = append(errors, fb.validatorError[ErrorRequired])
					break
				}
			}
		}
	case []Multipart:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
		if len(fv) > 0 {
			for _, item := range fv {
				if len(item.Data) == 0 {
					errors = append(errors, fb.validatorError[ErrorRequired])
					break
				}
			}
		}
	
	case string:
		if len(fv) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
	case int:
		if fv < 1 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
	case float64:
		if fv < 0.01 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
	case float32:
		if fv < 0.01 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
	case bool:
		if !fv {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
	
	case Multipart:
		if len(fv.Data) == 0 {
			errors = append(errors, fb.validatorError[ErrorRequired])
		}
	}
	return errors
}

func validateMin(fb *FieldBuilder, v validator) []error {
	errors := make([]error, 0)
	vv := v.value.(int)
	switch fv := fb.value.(type) {
	case []string:
		for _, item := range fv {
			if len(item) < vv {
				errors = append(errors, fb.validatorError[ErrorStringMin])
				break
			}
		}
	case []int:
		for _, item := range fv {
			if item < vv {
				errors = append(errors, fb.validatorError[ErrorNumberMin])
				break
			}
		}
	case []float32:
		for _, item := range fv {
			if item < float32(vv) {
				errors = append(errors, fb.validatorError[ErrorNumberMin])
				break
			}
		}
	case []float64:
		for _, item := range fv {
			if item < float64(vv) {
				errors = append(errors, fb.validatorError[ErrorNumberMin])
				break
			}
		}
	
	case string:
		if len(fv) < vv {
			errors = append(errors, fb.validatorError[ErrorStringMin])
		}
	case int:
		if fv < vv {
			errors = append(errors, fb.validatorError[ErrorNumberMin])
		}
	case float32:
		if fv < float32(vv) {
			errors = append(errors, fb.validatorError[ErrorNumberMin])
		}
	case float64:
		if fv < float64(vv) {
			errors = append(errors, fb.validatorError[ErrorNumberMin])
		}
	}
	return errors
}

func validateMax(fb *FieldBuilder, v validator) []error {
	errors := make([]error, 0)
	vv := v.value.(int)
	switch fv := fb.value.(type) {
	case []string:
		for _, item := range fv {
			if len(item) > vv {
				errors = append(errors, fb.validatorError[ErrorStringMax])
				break
			}
		}
	case []int:
		for _, item := range fv {
			if item > vv {
				errors = append(errors, fb.validatorError[ErrorNumberMax])
				break
			}
		}
	case []float32:
		for _, item := range fv {
			if item > float32(vv) {
				errors = append(errors, fb.validatorError[ErrorNumberMax])
				break
			}
		}
	case []float64:
		for _, item := range fv {
			if item > float64(vv) {
				errors = append(errors, fb.validatorError[ErrorNumberMax])
				break
			}
		}
	
	case string:
		if len(fv) > vv {
			errors = append(errors, fb.validatorError[ErrorStringMax])
		}
	case int:
		if fv > vv {
			errors = append(errors, fb.validatorError[ErrorNumberMax])
		}
	case float32:
		if fv > float32(vv) {
			errors = append(errors, fb.validatorError[ErrorNumberMax])
		}
	case float64:
		if fv > float64(vv) {
			errors = append(errors, fb.validatorError[ErrorNumberMax])
		}
	}
	return errors
}

func validateEmail(fb *FieldBuilder, v validator) []error {
	errors := make([]error, 0)
	switch fv := fb.value.(type) {
	case []string:
		for _, item := range fv {
			ok, err := regexp.MatchString(v.pattern, item)
			if err != nil {
				errors = append(errors, err)
			}
			if !ok {
				errors = append(errors, fb.validatorError[ErrorEmail])
			}
		}
	case string:
		if len(fv) > 0 {
			ok, err := regexp.MatchString(v.pattern, fv)
			if err != nil {
				errors = append(errors, err)
			}
			if !ok {
				errors = append(errors, fb.validatorError[ErrorEmail])
			}
		}
	}
	return errors
}

func validateCustom(fb *FieldBuilder, v validator) []error {
	errors := make([]error, 0)
	switch fv := fb.value.(type) {
	case []string:
		for _, item := range fv {
			ok, err := regexp.MatchString(v.pattern, item)
			if err != nil {
				errors = append(errors, err)
			}
			if !ok {
				errors = append(errors, fb.validatorError[ErrorInvalid])
			}
		}
	case string:
		if len(fv) > 0 {
			ok, err := regexp.MatchString(v.pattern, fv)
			if err != nil {
				errors = append(errors, err)
			}
			if !ok {
				errors = append(errors, fb.validatorError[ErrorInvalid])
			}
		}
	}
	return errors
}
