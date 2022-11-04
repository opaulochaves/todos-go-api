package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator func for creating a new validator for model fields.
func NewValidator() *validator.Validate {
  validate := validator.New()

  validate.RegisterCustomTypeFunc(ValidateUUID, uuid.UUID{})

  return validate
}

// ValidatorErrors func for showing validation error for each invalid field.
func ValidatorErrors(err error) map[string]string {
  fields := map[string]string{}

  for _, err := range err.(validator.ValidationErrors) {
    fields[err.Field()] = err.Error()
  }

  return fields
}

func ValidateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		return valuer.String()
	}
	return nil
}
