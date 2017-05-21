package xpath

import "testing"

func Test_LexBasics(t *testing.T) {
	l := lex("a")
	if !l.acceptAlphaNumeric(0) {
		t.Error("expected alphanumeric")
	}
	l = lex("/")
	if l.acceptAlphaNumeric(0) {
		t.Error("unexpected alphanumeric")
	}
	l = lex("1")
	if !l.acceptNumeric(0) {
		t.Error("unexpected numeric")
	}
	l = lex("1.2")
	if !l.acceptNumeric(0) {
		t.Error("unexpected numeric")
	}
}

func Test_LexExamples(t *testing.T) {
	tests := []struct {
		path     string
		expected []int
	}{
		{
			"a/b",
			[]int{token_name, kywd_slash, token_name},
		},
		{
			"a<1",
			[]int{token_name, token_operator, token_number},
		},
		{
			"a='b'",
			[]int{token_name, token_operator, token_literal},
		},
		{
			"a/b<1",
			[]int{token_name, kywd_slash, token_name, token_operator, token_number},
		},
	}
	for _, test := range tests {
		l := lex(test.path)
		for _, expected := range test.expected {
			token, err := l.nextToken()
			if err != nil {
				t.Errorf(err.Error())
			}
			if token.typ != expected {
				t.Errorf("%s - expected %s got %s - %v", test.path, expected, token.typ, token)
			}
		}
	}
}
