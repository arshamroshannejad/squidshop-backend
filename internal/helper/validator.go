package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var iranPhoneRegex = regexp.MustCompile(`^\+989\d{9}$`)

func RegisterValidations(v *validator.Validate) error {
	return v.RegisterValidation("irphone", func(fl validator.FieldLevel) bool {
		return iranPhoneRegex.MatchString(fl.Field().String())
	})
}
