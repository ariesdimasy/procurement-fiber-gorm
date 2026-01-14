package validations

import "fmt"

// ...existing code...

type FormatValidationError struct {
	Field string
	Err   error
}

func (e FormatValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("field %s: %v", e.Field, e.Err)
	}
	return fmt.Sprintf("invalid format for field %s", e.Field)
}

func (e FormatValidationError) Unwrap() error { return e.Err }
