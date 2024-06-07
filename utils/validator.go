package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) string {
	// Define a slice to collect error messages.
	var errorMessages []string

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, err.Field()+": "+err.Error())
	}

	// Join error messages with a comma and a space.
	return strings.Join(errorMessages, ", ")
}
