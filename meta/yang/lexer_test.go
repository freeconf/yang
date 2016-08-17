package yang

import (
	"fmt"
	"testing"

	"github.com/dhubler/c2g/meta"
)

func TestTokenStr(t *testing.T) {
	actual := Token{ParseEof, "EOF"}.String()
	if actual != "\"EOF\"" {
		t.Error(actual)
	}
}

func TestPosition(t *testing.T) {
	l := lex("x", nil)
	c1 := l.next()
	if c1 != 'x' {
		t.Errorf("next: unexpected rune %d", c1)
	}
	if l.pos != 1 {
		t.Errorf("next: unexpected position %d", l.pos)
	}
	if l.start != 0 {
		t.Errorf("next: unexpected start %d", l.start)
	}
	l.backup()
	if l.pos != 0 {
		t.Errorf("backup: unexpected position %d", l.pos)
	}
	l.peek()
	if l.pos != 0 {
		t.Errorf("peek: unexpected position %d", l.pos)
	}
	l.next()
	c2 := l.next()
	if c2 != eof {
		t.Errorf("next: expected eof %d", c2)
	}
}

func TestNextWithWhitespace(t *testing.T) {
	l := lex("  zzz \t  ggg", nil)
	l.acceptWS()
	c1 := l.next()
	if c1 != 'z' {
		t.Errorf("next: unexpected rune %d", c1)
	}

	l = lex("  /* this is a comment */ aaa", nil)
	l.acceptWS()
	c1 = l.next()
	if c1 != 'a' {
		t.Errorf("did not ignore comment")
	}

}

func TestAccept(t *testing.T) {
	l := lex("xyz", nil)
	p0 := l.pos
	l.acceptRun(0, "abc")
	if p0 != l.pos {
		t.Error("Shouldn't find abc")
	}
	l.acceptRun(0, "yx")
	if p0 != l.pos-2 {
		t.Errorf("Should only advance 2 not %d", l.pos-p0)
	}
	l = lex("   \t\t  x  \n\n  ", nil)
	l.acceptWS()
	c0 := l.next()
	if c0 != 'x' {
		t.Errorf("ignore ws landed on %d", c0)
	}
}

func TestBegin(t *testing.T) {
	l := lex(" module foo {", nil)
	next := lexBegin(l)
	if next == nil {
		t.Error("expected lexModule")
	}
	if l.head != 3 {
		LogTokens(l)
		t.Errorf("expected 3 module tokens but got %d", l.head)
	}
}

func TestNextToken(t *testing.T) {
	l := lex(" ", nil)
	token, _ := l.nextToken()
	if token.typ != ParseEof {
		t.Errorf("unexpected token(%d) %s", token.typ, token)
	}
}

func TestMaxElements(t *testing.T) {
	l := lex("max-elements 100;", nil)
	if !l.acceptToken(kywd_max_elements) {
		t.Errorf("unexpected max-elements")
	}
	if !l.acceptInteger(token_int) {
		t.Errorf("expected int")
	}
	l.popToken()
	token := l.popToken()
	if token.val != "100" {
		t.Errorf("expected 100, got '%s'", token.val)
	}
}

func TestAlphaNumeric(t *testing.T) {
	l := lex("aaa zzz", nil)
	if !l.acceptAlphaNumeric(0) {
		t.Errorf("unexpected alphanumeric")
	}
	token := l.popToken()
	if token.val != "aaa" {
		t.Errorf("expected 'aaa' but got '%s'", token.val)
	}
}

func TestString(t *testing.T) {
	expected := "\"string here\""
	l := lex(expected, nil)
	if !l.acceptString(0) {
		t.Errorf("unexpected alphanumeric")
	}
	token := l.popToken()
	if token.val != expected {
		t.Errorf("expected '%s' but got '%s'", expected, token.val)
	}
}

func TestAcceptRun(t *testing.T) {
	l := lex("aaabbbzzz", nil)
	if !l.acceptRun(0, "abc") {
		t.Errorf("unexpected alphanumeric")
	}
	token := l.popToken()
	if token.val != "aaabbb" {
		t.Errorf("expected ")
	}

	l = lex("2015-06-03 {", nil)
	if !l.acceptRun(0, "0123456789-") {
		t.Errorf("unexpected alphanumeric")
	}
	token = l.popToken()
	if token.val != "2015-06-03" {
		t.Errorf("expected ")
	}
}

func TestModule(t *testing.T) {
	l := lex("module foo { } ", nil)
	expecteds := [...]int{kywd_module, token_ident, token_curly_open, token_curly_close}
	for _, expected := range expecteds {
		token, err := l.nextToken()
		if err != nil {
			t.Errorf(err.Error())
		}
		if token.typ != expected {
			t.Errorf("expected %d but got %d, %s", expected, token.typ, token.String())
		}
	}
}

func TestChoice(t *testing.T) {
	l := lex("choice foo { } ", nil)
	expecteds := [...]int{kywd_choice, token_ident, token_curly_open, token_curly_close}
	for _, expected := range expecteds {
		token, err := l.nextToken()
		if err != nil {
			t.Errorf(err.Error())
		}
		if token.typ != expected {
			t.Errorf("expected %d but got %d, %s", expected, token.typ, token.String())
		}
	}
}

func TestBasicYang(t *testing.T) {
	yang := `
module foo {
	namespace "x";
	prefix "z";
	revision 2015-06-03 {
	  description "d1";
	}
	container q {
	  description "d2";
	  leaf state {
	    config "false";
	    description "d3";
	  }
	  leaf-list ids {
	    type int;
	  }
	}
}`
	expecteds := [...]int{
		kywd_module, token_ident, token_curly_open,
		kywd_namespace, token_string, token_semi,
		kywd_prefix, token_string, token_semi,
		kywd_revision, token_rev_ident, token_curly_open,
		kywd_description, token_string, token_semi,
		token_curly_close,
		kywd_container, token_ident, token_curly_open,
		kywd_description, token_string, token_semi,
		kywd_leaf, token_ident, token_curly_open,
		kywd_config, token_string, token_semi,
		kywd_description, token_string, token_semi,
		token_curly_close,
		kywd_leaf_list, token_ident, token_curly_open,
		kywd_type, token_ident, token_semi,
		token_curly_close,
		token_curly_close,
		token_curly_close,
	}

	l := lex(yang, nil)
	for _, expected := range expecteds {
		token, err := l.nextToken()
		if err != nil {
			t.Errorf(err.Error())
		}
		if token.typ != expected {
			t.Fatalf("expected %s but got %s - %s", l.keyword(expected), l.keyword(token.typ), token.String())
		}
	}
}

func TestStack(t *testing.T) {
	stack := newDefStack(10)
	expected := &meta.Module{Ident: "x"}
	stack.Push(expected)
	actual, ok := stack.Pop().(*meta.Module)
	if !ok {
		t.Fail()
	}
	if actual.Ident != expected.Ident {
		t.Fail()
	}
}

func LogTokens(l *lexer) {
	fmt.Printf("Tokens %s\n", l.tokens[l.tail:l.head])
}
