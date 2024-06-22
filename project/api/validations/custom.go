package validations

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil
}

func IsAdult(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(int)
	fmt.Printf("value: %v\n", value)
	if !ok || value < 18 {
		errors.New("invalid value")
		return false
	}
	return true
}
