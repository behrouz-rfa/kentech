package errs

import (
	"errors"
	"net/http"
)

// ThrowBadRequestError returns HTTPError with err and http.StatusBadRequest
func ThrowBadRequestError(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Bad Request")
	}
	return New(err, http.StatusBadRequest)
}

// ThrowInternalServerError returns HTTPError with err and http.StatusInternalServerError
func ThrowInternalServerError(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Internal Server Error")
	}
	return New(err, http.StatusInternalServerError)
}

// ThrowNotFoundError returns HTTPError with err and http.StatusNotFound
func ThrowNotFoundError(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Not Found")
	}
	return New(err, http.StatusNotFound)
}

// ThrowUnauthorizedError returns HTTPError with err and http.StatusUnauthorized
func ThrowUnauthorizedError(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Unauthorized")
	}
	return New(err, http.StatusUnauthorized)
}

// ThrowServiceUnavailableError returns HTTPError with err and http.StatusServiceUnavailable
func ThrowServiceUnavailableError(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Service Unavailable")
	}
	return New(err, http.StatusServiceUnavailable)
}

// ThrowForbiddenError returns HTTPError with err and http.StatusForbidden
func ThrowForbiddenError(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Forbidden")
	}
	return New(err, http.StatusForbidden)
}

// ThrowUnprocessableEntity returns HTTPError with err and http.StatusUnprocessableEntity
func ThrowUnprocessableEntity(err error) *HTTPError {
	if err == nil || err.Error() == "" {
		err = errors.New("Unprocessable Content")
	}
	return New(err, http.StatusUnprocessableEntity)
}

// ThrowValidationErrors returns HTTPError with validations and http.StatusBadRequest
func ThrowValidationErrors(errs ...ValidationError) *HTTPError {
	e := ThrowUnprocessableEntity(nil)
	return e.Validations(errs...)
}
