package errs

import "github.com/stackus/errors"

var ErrBadRequest = ThrowBadRequestError(errors.ErrBadRequest)
var ErrInternalServerError = ThrowInternalServerError(errors.ErrInternalServerError)
var ErrUnauthorized = ThrowUnauthorizedError(errors.ErrUnauthorized)
var ErrForbidden = ThrowForbiddenError(errors.ErrUnauthorized)
var ErrNotFound = ThrowNotFoundError(errors.ErrNotFound)
