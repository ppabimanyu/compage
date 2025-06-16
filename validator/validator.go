package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Validator struct {
	v *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("name")
	})
	return &Validator{v: v}
}

func (v *Validator) Struct(s interface{}) map[string]string {
	err := v.v.Struct(s)
	if err != nil {
		return v._formatValidationError(err)
	}
	return nil
}

func (v *Validator) Var(field interface{}, tag string) map[string]string {
	err := v.v.Var(field, tag)
	if err != nil {
		return v._formatValidationError(err)
	}
	return nil
}

func (v *Validator) _formatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errors[err.Field()] = fmt.Sprintf("%s is required", err.Field())
		case "email":
			errors[err.Field()] = fmt.Sprintf("%s is not a valid email", err.Field())
		case "min":
			errors[err.Field()] = fmt.Sprintf("%s must be at least %s", err.Field(), err.Param())
		case "max":
			errors[err.Field()] = fmt.Sprintf("%s must be at most %s", err.Field(), err.Param())
		case "len":
			errors[err.Field()] = fmt.Sprintf("%s must be %s characters long", err.Field(), err.Param())
		case "gte":
			errors[err.Field()] = fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
		case "gt":
			errors[err.Field()] = fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
		case "lte":
			errors[err.Field()] = fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
		case "lt":
			errors[err.Field()] = fmt.Sprintf("%s must be less than %s", err.Field(), err.Param())
		case "numeric":
			errors[err.Field()] = fmt.Sprintf("%s must be numeric", err.Field())
		case "number":
			errors[err.Field()] = fmt.Sprintf("%s must be a number", err.Field())
		case "phone":
			errors[err.Field()] = fmt.Sprintf("%s invalid phone number", err.Field())
		default:
			errors[err.Field()] = fmt.Sprintf("%s is not valid", err.Field())
		}
	}
	return errors
}
