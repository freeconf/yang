package c2

import (
	"bufio"
	"bytes"
	"runtime"
	"strconv"
)

type HttpError interface {
	error
	Stack() string
	HttpCode() int
}

func IsNotFoundErr(err error) bool {
	if herr, isHErr := err.(HttpError); isHErr {
		return herr.HttpCode() == 404
	}
	return false
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

func NewErr(msg string) error {
	return &codedErrorString{
		s: msg,
		c: 500,
	}
}

func NewErrC(msg string, code int) error {
	return &codedErrorString{
		s: msg,
		c: code,
	}
}

func trim(s string, max int) string {
	if len(s) > max {
		return "..." + s[len(s)-(max+3):]
	}
	return s
}

func DumpStack() string {
	var buff bytes.Buffer
	w := bufio.NewWriter(&buff)
	var stack [25]uintptr
	len := runtime.Callers(2, stack[:])
	for i := 1; i < len; i++ {
		f := runtime.FuncForPC(stack[i])
		w.WriteRune(' ')
		w.WriteString(f.Name())
		w.WriteRune(' ')
		file, lineno := f.FileLine(stack[i-1])
		w.WriteString(trim(file, 20))
		w.WriteRune(':')
		w.WriteString(strconv.Itoa(lineno))
		w.WriteString("\n")
	}
	w.Flush()
	return buff.String()
}
