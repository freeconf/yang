package yang

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSimpleLexExample(t *testing.T) {
	l := lex(TestDataSimpleYang, nil)
	tokens := list.New()
	for {
		token, err := l.nextToken()
		if err != nil {
			t.Errorf(err.Error())
		}
		if token.typ == ParseEof {
			break
		}
		tokens.PushBack(token)
	}
	if tokens.Len() != 308 {
		for e := tokens.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
		LogTokens(l)
		t.Fatalf("wrong num tokens %d", tokens.Len())
	}
}

func TestStoneLex(t *testing.T) {
	lex(TestDataRomancingTheStone, nil)
}
