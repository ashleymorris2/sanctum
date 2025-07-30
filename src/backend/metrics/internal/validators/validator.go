package validators

import "github.com/go-playground/validator/v10"

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	return rv.validator.Struct(i)
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{
		validator: validator.New(),
	}
}
