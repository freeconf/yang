package meta

import "testing"

func TestSet(t *testing.T) {
	m := NewModule("x")
	err := Set(m, nil)
	if err == nil {
		t.Error("expected err")
	}
}
