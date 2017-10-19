package yang

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/c2stack/c2g/c2"
)

var updateFlag = flag.Bool("update", false, "update golden files instead of verifying against them")

func TestLexExamples(t *testing.T) {
	tests := []string{
		"foo",
		"rtstone",
		"turing-machine",
	}
	for _, test := range tests {
		y, _ := ioutil.ReadFile("./testdata/" + test + ".yang")
		l := lex(string(y), nil)
		var actual bytes.Buffer
		for {
			token, err := l.nextToken()
			if err != nil {
				t.Errorf(err.Error())
			} else if token.typ == ParseEof {
				break
			}
			actual.WriteString(fmt.Sprintf("%s %s\n", l.keyword(token.typ), token.String()))
		}
		c2.Gold(t, *updateFlag, actual.Bytes(), "gold/"+test+".lex")
	}
}
