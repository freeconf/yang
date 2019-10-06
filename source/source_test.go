package source

import (
	"testing"
)

func TestDirNotFound(t *testing.T) {
	m := Dir(".")
	s, err := m("bogus", ".txt")
	if s != nil {
		t.Error("expected no stream")
	}
	if err != nil {
		t.Error("expected no err")
	}
}
