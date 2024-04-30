package http

import (
	"errors"
	errs "github.com/behrouz-rfa/kentech/internal/core/errors"
	"net/http"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	model.ErrInternal:                   http.StatusInternalServerError,
	model.ErrDataNotFound:               http.StatusNotFound,
	model.ErrConflictingData:            http.StatusConflict,
	model.ErrInvalidCredentials:         http.StatusUnauthorized,
	model.ErrUnauthorized:               http.StatusUnauthorized,
	model.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	model.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	model.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	model.ErrInvalidToken:               http.StatusUnauthorized,
	model.ErrExpiredToken:               http.StatusUnauthorized,
	model.ErrForbidden:                  http.StatusForbidden,
	model.ErrNoUpdatedData:              http.StatusBadRequest,
	model.ErrInsufficientStock:          http.StatusBadRequest,
	model.ErrInsufficientPayment:        http.StatusBadRequest,
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func handleError(ctx *gin.Context, err error) {

	e, ok := err.(*errs.HTTPError)
	if ok {
		ctx.AbortWithStatusJSON(e.Code, newErrorResponse([]string{e.Details}))
		return
	}

	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

// handleAbort sends an error response and aborts the request with the specified status code and error message
func handleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.AbortWithStatusJSON(statusCode, errRsp)
}

// parseError parses error messages from the error object and returns a slice of error messages
func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

// errorResponse represents an error response body format
type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// newErrorResponse is a helper function to create an error response body
func NewErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

// response represents a response body format
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// meta represents metadata for a paginated response
type meta struct {
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip uint64) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

// userResponse represents a user response body
type userResponse struct {
	*model.User
}

// newUserResponse is a helper function to create a response body for handling user data
func newUserResponse(user *model.User) userResponse {
	return userResponse{
		User: user,
	}
}

// userResponse represents a user response body
type filmResponse struct {
	*model.Film
}

// newUserResponse is a helper function to create a response body for handling user data
func newFilmResponse(film *model.Film) filmResponse {
	return filmResponse{
		film,
	}
}
