package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(v any) error {
	return validate.Struct(v)
}
