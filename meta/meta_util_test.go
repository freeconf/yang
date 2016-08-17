package meta

import "testing"

func TestMetaNameToFieldName(t *testing.T) {
	var actual string
	tests := []struct {
		in  string
		out string
	}{
		{"X", "X"},
		{"x", "X"},
		{"abc", "Abc"},
		{"ABC", "ABC"},
		{"abCd", "AbCd"},
		{"one-two", "OneTwo"},
	}
	for _, test := range tests {
		if actual = MetaNameToFieldName(test.in); actual != test.out {
			t.Error(test.out, "!=", actual)
		}
	}
}

func TestMetaDeepCopy(t *testing.T) {
	p := &Container{Ident: "p"}
	c := &Container{Ident: "C"}
	p.AddMeta(c)
	c.AddMeta(&Uses{Ident: "g"})
	g := &Grouping{Ident: "g"}
	p.AddMeta(g)

	copy := DeepCopy(p).(*Container)
	if copy == p {
		t.Error("did not copy")
	}
	if ListLen(p) != ListLen(copy) {
		t.Error("not same size")
	}
}
