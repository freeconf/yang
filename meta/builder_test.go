package meta

import "testing"

func TestBuilder(t *testing.T) {
	b := &Builder{}
	b.Must(new(Leaf), "x")
	if b.LastErr != nil {
		t.Error(b.LastErr)
	}
}
