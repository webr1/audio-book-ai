package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validate *validator.Validate
}

func New() *CustomValidator {
	return &CustomValidator{validate: validator.New()}
}

func (v *CustomValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
