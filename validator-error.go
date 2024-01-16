package form

import "errors"

const (
	ErrorRequired  = "required"
	ErrorStringMin = "stringMin"
	ErrorStringMax = "stringMax"
	ErrorNumberMin = "numberMin"
	ErrorNumberMax = "numberMax"
	ErrorEmail     = "email"
	ErrorInvalid   = "invalid"
)

var (
	validatorErrorRequired  = errors.New("field is required")
	validatorErrorStringMin = errors.New("field value is too short")
	validatorErrorStringMax = errors.New("field value is too long")
	validatorErrorNumberMin = errors.New("field value is too low")
	validatorErrorNumberMax = errors.New("field value is too high")
	validatorErrorEmail     = errors.New("email is invalid")
	validatorErrorInvalid   = errors.New("field is invalid")
)

func createDefaultErrors() map[string]error {
	return map[string]error{
		ErrorRequired:  validatorErrorRequired,
		ErrorStringMin: validatorErrorStringMin,
		ErrorStringMax: validatorErrorStringMax,
		ErrorNumberMin: validatorErrorNumberMin,
		ErrorNumberMax: validatorErrorNumberMax,
		ErrorEmail:     validatorErrorEmail,
		ErrorInvalid:   validatorErrorInvalid,
	}
}
