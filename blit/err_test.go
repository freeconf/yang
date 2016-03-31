package blit
import (
	"testing"
)

func TestErrorPrintStackTrace(t *testing.T) {
	t.Log(DumpStack())
}

func TestErrorNew(t *testing.T) {
	_ = NewErrC("x", 501)
}
