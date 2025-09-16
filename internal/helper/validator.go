package helper

import (
	"regexp"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/go-playground/validator/v10"
)

var iranPhoneRegex = regexp.MustCompile(`^\+989\d{9}$`)

func RegisterValidations(v *validator.Validate) error {
	v.RegisterStructValidation(validateCategoryQueryParams, entity.CategoryQueryParamRequest{})
	return v.RegisterValidation("irphone", validateIranPhone)
}

func validateCategoryQueryParams(sl validator.StructLevel) {
	req := sl.Current().Interface().(entity.CategoryQueryParamRequest)
	if req.Name == "" && req.Slug == "" {
		sl.ReportError(req.Name, "Name", "name", "required_without_slug", "either name or slug must be set")
		sl.ReportError(req.Slug, "Slug", "slug", "required_without_name", "either slug or name must be set")
	}
}

func validateIranPhone(fl validator.FieldLevel) bool {
	return iranPhoneRegex.MatchString(fl.Field().String())
}
