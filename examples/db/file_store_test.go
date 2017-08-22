package db

import (
	"io/ioutil"
	"testing"

	"os"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/nodes"
)

func Test_FileStore(t *testing.T) {
	fs := FileStore{VarDir: "./var"}
	b, _ := nodes.BirdBrowser(`{"bird":[{
		"name" : "robin",
		"wingspan" : 10
	}]}`)
	err := fs.DbWrite("x", "m", b)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Open("./var/x:m.json")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	expected := `{
"bird":[
  {
    "name":"robin",
    "wingspan":10}]}`
	c2.AssertEqual(t, string(actual), expected)
}
