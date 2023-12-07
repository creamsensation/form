package form

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	t.Run(
		"required", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text(), Validate.Required()),
					Add("quantity").With(Number[int](), Validate.Required()),
					Add("amount").With(Number[float64](), Validate.Required()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, validatorErrorRequired, form.Email.Errors[0])
			assert.Equal(t, validatorErrorRequired, form.Quantity.Errors[0])
			assert.Equal(t, validatorErrorRequired, form.Amount.Errors[0])
		},
	)
	t.Run(
		"email valid", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Email("test@test.cz"), Validate.Email()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, 0, len(form.Email.Errors))
		},
	)
	t.Run(
		"email invalid", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text("test"), Validate.Email()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, validatorErrorEmail, form.Email.Errors[0])
		},
	)
	t.Run(
		"string min", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text("test"), Validate.Min(5)),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, validatorErrorStringMin, form.Email.Errors[0])
		},
	)
	t.Run(
		"string max", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text("test"), Validate.Max(3)),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, validatorErrorStringMax, form.Email.Errors[0])
		},
	)
	t.Run(
		"string slice", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("roles").Multiple().With(Text(), Validate.Required()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, 1, len(form.Roles.Errors))
			assert.Equal(t, validatorErrorRequired, form.Roles.Errors[0])
		},
	)
}
