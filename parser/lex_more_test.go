package parser

import (
	"bytes"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/freeconf/yang/fc"
)

var updateFlag = flag.Bool("update", false, "update golden files instead of verifying against them")

func TestLexSamples(t *testing.T) {
	for _, test := range yangTestFiles {
		testId := test.dir + "/" + test.fname
		t.Log(testId)
		y, _ := ioutil.ReadFile("./testdata" + testId + ".yang")
		var actual bytes.Buffer
		if err := lexDump(string(y), &actual); err != nil {
			t.Error(err)
		} else {
			fc.Gold(t, *updateFlag, actual.Bytes(), "./testdata"+test.dir+"/gold/"+test.fname+".lex")
		}
	}
}
