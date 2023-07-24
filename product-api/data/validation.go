package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError wraps the validators FieldError
// so we ain't exposing this out to code
type ValidationError struct {
	validator.FieldError
}

// Returns validation error message
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: %s Error: Field validation for %s failed on %s",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}


// Collection of ValidationError
type ValidationErrors []ValidationError


// Converts the errors slice into string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}


type Validation struct {
	validate *validator.Validate
}


// Returns new validation
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}	
}


// Validates SKU field
func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in the format abc-abc-abc
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	sku := re.FindAllString(fl.Field().String(), -1)
	return len(sku) == 1
}


// Validates the interface i
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	 var returnErrs []ValidationError
	 for _, err := range errs {
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	 }

	 return returnErrs
}

