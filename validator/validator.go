package validator

import (
	"fmt"
	"strings"
)

// Validator holds the validation errors.
type Validator struct {
	Errors map[string]string // Map of validation errors
}

// New creates a new instance of the Validator struct.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid returns true if there are no errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds a new error to the map.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// NotEmpty checks if a string is empty.
func NotEmpty(v *Validator, field, value string) {
	if strings.TrimSpace(value) == "" {
		v.AddError(field, "must not be empty")
	}
}

// GreaterThenEquals checks if a value is greater then or equal to a minimum value.
func GreaterThenEquals(v *Validator, field string, value int, min int) {
	if value < min {
		v.AddError(field, fmt.Sprintf("must be greater then or equal to %d", min))
	}
}

// LowerThenEquals checks if a value is lower then or equal to a maximum value.
func LowerThenEquals(v *Validator, field string, value int, max int) {
	if value > max {
		v.AddError(field, fmt.Sprintf("must be lower then or equal to %d", max))
	}
}
