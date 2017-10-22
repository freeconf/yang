package yang_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
)

var updateFlag = flag.Bool("update", false, "update golden files instead of verifying against them")

func TestLexSamples(t *testing.T) {
	for _, test := range yangTestFiles {
		y, _ := ioutil.ReadFile("./testdata" + test.dir + "/" + test.fname + ".yang")
		var actual bytes.Buffer
		yang.LexDump(string(y), &actual)
		c2.Gold(t, *updateFlag, actual.Bytes(), "gold/"+test.dir+"/"+test.fname+".lex")
	}
}
