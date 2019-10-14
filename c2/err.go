package c2

import "net/http"

// NotFoundError is when a resource is not found.
// When used with RESTCONF context will result in 404 error
type NotFoundError string

// Error implements error interface
func (err NotFoundError) Error() string {
	return string(err)
}

// NotImplementedError is when something wasn't implemented and may be
// an optional part of the spec or a particular feature wasn't impemented
// When used with RESTCONF context, will result in 501 error.
type NotImplementedError string

// Error implements error interface
func (err NotImplementedError) Error() string {
	return string(err)
}

// BadRequestError is when end user is attempting to perform an operation
// that is invalid.
// When used with RESTCONF context, will result in 400 error.
type BadRequestError string

// Error implements error interface
func (err BadRequestError) Error() string {
	return string(err)
}

// ConflictError is when multiple attempts to do something cannot be completed
// When used with RESTCONF context, will result in 409 error.
type ConflictError string

// Error implements error interface
func (err ConflictError) Error() string {
	return string(err)
}

// HttpError is for one of the stadard http errors that do not require custom text
type HttpError int

// Error implements error interface
func (err HttpError) Error() string {
	return http.StatusText(int(err))
}

// HttpableError will see if error be converted to one of non-500 errors
func HttpableError(err error) (int, bool) {
	switch x := err.(type) {
	case NotFoundError:
		return 404, true
	case NotImplementedError:
		return 501, true
	case BadRequestError:
		return 400, true
	case ConflictError:
		return 409, true
	case HttpError:
		return int(x), true
	}
	return 0, false
}
