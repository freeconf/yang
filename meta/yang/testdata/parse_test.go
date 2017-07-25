package testdata

import (
	"flag"
	"log"
	"path/filepath"
	"testing"

	"fmt"
	"os"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

var update = flag.Bool("update", false, "update .golden files")

func Test_Files(t *testing.T) {
	ypath := &meta.FileStreamSource{Root: "."}
	files, err := filepath.Glob("*.yang")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		mod := file[:len(file)-5]
		m, err := yang.LoadModule(ypath, mod)
		if err != nil {
			t.Errorf("file %s, err: %s", file, err.Error())
			continue
		}
		b := node.SelectModule(m, false)
		var outFile string
		if *update {
			outFile = fmt.Sprintf("golden/%s.json", mod)
		} else {
			outFile = fmt.Sprintf("var/%s.json", mod)
		}
		out, err := os.Create(outFile)
		if err != nil {
			t.Fatal(err)
		}
		if err := b.Root().InsertInto(node.NewJsonPretty(out).Node()).LastErr; err != nil {
			t.Error(mod, err)
			continue
		}
		out.Close()
		if !*update {
			if err := c2.DiffFiles(fmt.Sprintf("./golden/%s.json", mod), outFile); err != nil {
				t.Error(err)
			}
		}
	}
}
