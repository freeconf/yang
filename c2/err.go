package c2

// HttpError is for errors that want to influence the http error
// response.
type HttpError interface {
	error
	HttpCode() int
}

// HttpCode deterimines the HTTP code from an error
func HttpCode(err error) int {
	if herr, isHErr := err.(HttpError); isHErr {
		return herr.HttpCode()
	}
	return 500
}

type codedErrorString struct {
	s string
	c int
}

func (e *codedErrorString) Error() string {
	return e.s
}

// 404 - Not Found
// 500 - Internal error
// 501 - Not implemented
// 400 - Bad request (user error)
// 409 - Conflict - Like bad request, but based on existing state of data/system.
func (e *codedErrorString) HttpCode() int {
	return e.c
}

// NewErr general error
func NewErr(msg string) error {
	return &codedErrorString{
		s: msg,
		c: 500,
	}
}

// NewErrC that wants to influence the http error code in a response. Error
// may not may not be used in a http context but error code may still be
// useful to determine the nature of the error
func NewErrC(msg string, httpErrorCode int) error {
	return &codedErrorString{
		s: msg,
		c: httpErrorCode,
	}
}
