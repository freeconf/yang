package fc

import (
	"errors"
	"net/http"
	"os"
)

// NotFoundError is when a resource is not found.
// When used with RESTCONF context will result in 404 error
var NotFoundError = errors.New("not found")

// NotImplementedError is when something wasn't implemented and may be
// an optional part of the spec or a particular feature wasn't impemented
// When used with RESTCONF context, will result in 501 error.
var NotImplementedError = errors.New("not implemented")

// BadRequestError is when end user is attempting to perform an operation
// that is invalid.
// When used with RESTCONF context, will result in 400 error.
var BadRequestError = errors.New("bad request")

// ConflictError is when multiple attempts to do something cannot be completed
// When used with RESTCONF context, will result in 409 error.
var ConflictError = errors.New("conflict")

// UnauthorizedError when someone is attempting to do something they do not have access to
var UnauthorizedError = errors.New("not authorized")

// HttpableError will see if error be converted to one of non-500 errors
func HttpStatusCode(err error) int {
	if errors.Is(err, NotFoundError) {
		return http.StatusNotFound
	}
	if errors.Is(err, os.ErrNotExist) {
		return http.StatusNotFound
	}
	if errors.Is(err, NotImplementedError) {
		return http.StatusNotImplemented
	}
	if errors.Is(err, BadRequestError) {
		return http.StatusBadRequest
	}
	if errors.Is(err, ConflictError) {
		return http.StatusConflict
	}
	if errors.Is(err, ConflictError) {
		return http.StatusConflict
	}
	if errors.Is(err, UnauthorizedError) {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
