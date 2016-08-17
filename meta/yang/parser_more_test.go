package yang

import "testing"

func TestSimpleParse(t *testing.T) {
	//yyDebug = 4
	l := lex(TestDataSimpleYang, nil)
	err_code := yyParse(l)
	if err_code != 0 {
		t.Error(l.lastError)
	}
}

func TestStoneParse(t *testing.T) {
	//yyDebug = 4
	l := lex(TestDataRomancingTheStone, nil)
	err_code := yyParse(l)
	if err_code != 0 {
		t.Error(l.lastError)
	}
}
