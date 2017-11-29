package meta

import "testing"

func TestSet(t *testing.T) {
	m := NewModule("x", nil)
	err := Set(m, nil)
	if err == nil {
		t.Error("expected err")
	}
}
