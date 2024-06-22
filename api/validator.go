package api

import (
	"github.com/eizyc/simplebank/util"
	"github.com/go-playground/validator/v10"
)

// the type `validator.Func` expected by the validator package for custom validation functions.
// It takes a validator.FieldLevel parameter, which provides access to the field being validated and its metadata.
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// Attempt to cast the field's value to a string
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// If successful, call a utility function to check if the currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
