package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Error list.
var (
	ErrInvalidDBFormat      = errors.New("invalid db address")
	ErrNotFoundBoilerplate  = errors.New("not found boilerplate")
	ErrNotFoundProduct     = errors.New("not found product")
	ErrInvalidRequestFormat = errors.New("invalid request format")
	ErrInternalDB           = errors.New("internal database error")
	ErrInternalElastic      = errors.New("internal elastic error")
	ErrInternalCache        = errors.New("internal cache error")
	ErrInternalServer       = errors.New("internal server error")
)

// ErrRequiredField is error for missing field.
func ErrRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

// ErrGTField is error for greater than field.
func ErrGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

// ErrGTEField is error for greater than or equal field.
func ErrGTEField(str, value string) error {
	return fmt.Errorf("field %s must be greater than or equal %s", str, value)
}

// ErrLTField is error for lower than field.
func ErrLTField(str, value string) error {
	return fmt.Errorf("field %s must be lower than %s", str, value)
}

// ErrLTEField is error for lower than or equal field.
func ErrLTEField(str, value string) error {
	return fmt.Errorf("field %s must be lower than or equal %s", str, value)
}

// ErrLenField is error for length field.
func ErrLenField(str, value string) error {
	return fmt.Errorf("field %s length must be %s", str, value)
}

// ErrISO3166Alpha2Field is error for ISO 3166-1 alpha-2 field.
func ErrISO3166Alpha2Field(str string) error {
	return fmt.Errorf("field %s must be in ISO 3166-1 alpha-2 format", str)
}

// ErrEmailField is error for email field.
func ErrEmailField(str string) error {
	return fmt.Errorf("field %s must be in email format", str)
}

// ErrURLField is error for url field.
func ErrURLField(str string) error {
	return fmt.Errorf("field %s must be in URL format", str)
}

// ErrInvalidFormatField is error for invalid format field.
func ErrInvalidFormatField(str string) error {
	return fmt.Errorf("invalid format field %s", str)
}

// ErrOneOfField is error for oneof field.
func ErrOneOfField(str, value string) error {
	return fmt.Errorf("field %s must be one of %s", str, strings.Join(strings.Split(value, " "), "/"))
}

// ErrNumericField is error for numeric field.
func ErrNumericField(str string) error {
	return fmt.Errorf("field %s must contain number only", str)
}

// ErrAlphaField is error for alpha field.
func ErrAlphaField(str string) error {
	return fmt.Errorf("field %s must contain letter only", str)
}

// ErrMinPrice is error for min price.
func ErrMinPrice(price float64) error {
	return fmt.Errorf("minimum price is %d", int(price))
}

// ErrMaxPrice is error for max price.
func ErrMaxPrice(price float64) error {
	return fmt.Errorf("maximum price is %d", int(price))
}

// ErrMinAmount is error for min amount.
func ErrMinAmount(amount float64) error {
	return fmt.Errorf("minimum amount is %d", int(amount))
}

// ErrMaxAmount is error for max amount.
func ErrMaxAmount(amount float64) error {
	return fmt.Errorf("maximum amount is %d", int(amount))
}

// ErrDatetimeField is error for datetime field.
func ErrDatetimeField(str string) error {
	return fmt.Errorf("%s must be in date format (dd/mm/yyyy)", str)
}
