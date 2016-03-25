package blit
import (
	"testing"
)

func TestErrorPrintStackTrace(t *testing.T) {
	t.Log(dumpStack())
}

func TestErrorNew(t *testing.T) {
	_ = NewErrC("x", 501)
}
