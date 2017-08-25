package node

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
