package errs

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func NewValidationError(field, error string) ValidationError {
	return ValidationError{
		Field: field,
		Error: error,
	}
}
