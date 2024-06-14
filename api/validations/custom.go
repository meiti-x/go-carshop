package validations

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func IsAdult(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(int)
	fmt.Printf("value: %v\n", value)
	if !ok || value < 18 {
		errors.New("invalid value")
		return false
	}
	return true
}
