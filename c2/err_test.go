package c2

import (
	"testing"
)

func TestErrorNew(t *testing.T) {
	_ = NewErrC("x", 501)
}
