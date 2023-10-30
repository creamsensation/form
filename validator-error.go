package form

import "errors"

type validatorError struct {
	required  error
	stringMin error
	stringMax error
	numberMin error
	numberMax error
	email     error
	invalid   error
}

var (
	validatorErrorRequired  = errors.New("field is required")
	validatorErrorStringMin = errors.New("field value is too short")
	validatorErrorStringMax = errors.New("field value is too long")
	validatorErrorNumberMin = errors.New("field value is too low")
	validatorErrorNumberMax = errors.New("field value is too high")
	validatorErrorEmail     = errors.New("email is invalid")
	validatorErrorInvalid   = errors.New("field is invalid")
)

func createDefaultErrors() validatorError {
	return validatorError{
		required:  validatorErrorRequired,
		stringMin: validatorErrorStringMin,
		stringMax: validatorErrorStringMax,
		numberMin: validatorErrorNumberMin,
		numberMax: validatorErrorNumberMax,
		email:     validatorErrorEmail,
		invalid:   validatorErrorInvalid,
	}
}
