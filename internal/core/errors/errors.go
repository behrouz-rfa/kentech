package errs

import (
	"fmt"
	"net/http"
)

// HTTPError is an error that will be rendered to the client.
type HTTPError struct {
	Code              int               `json:"-"`
	Message           string            `json:"original_message,omitempty"`
	Details           string            `json:"details,omitempty"`
	ExtraDetails      string            `json:"extraDetails,omitempty"`
	ReferenceId       string            `json:"reference_id,omitempty"`
	TranslationKey    string            `json:"-"`
	TranslatedMessage string            `json:"message,omitempty"`
	ValidationErrors  []ValidationError `json:"validation_errors"`
}

// New initializes new HTTPError
func New(err error, code int, validationErrs ...ValidationError) *HTTPError {
	if code <= 0 {
		code = http.StatusInternalServerError
	}

	e := &HTTPError{
		Code: code,
	}
	if err != nil {
		e.Message = err.Error()
	}

	if len(validationErrs) > 0 {
		e.ValidationErrors = validationErrs
	}

	return e
}

// Detail sets HTTPError Details
func (e *HTTPError) Detail(details string) *HTTPError {
	e.Details = details
	return e
}

// ReferenceID sets HTTPError ReferenceId
func (e *HTTPError) ReferenceID(referenceID string) *HTTPError {
	e.ReferenceId = referenceID
	return e
}

// Key sets HTTPError TranslationKey
func (e *HTTPError) Key(translationKey string) *HTTPError {
	e.TranslationKey = translationKey
	return e
}

// Error implements the error interface
func (e *HTTPError) Error() string {
	return fmt.Sprintf("handler: %v: %v %v", e.Code, e.Message, e.Details)
}

// Validations appends validation errors to HTTPError.ValidationErrors
func (e *HTTPError) Validations(validations ...ValidationError) *HTTPError {
	e.ValidationErrors = append(e.ValidationErrors, validations...)
	return e
}
