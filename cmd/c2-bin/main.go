package main

import (
	"os"
	"github.com/c2g/browse"
	"github.com/c2g/node"
	"flag"
	"io"
)


var inFilePtr = flag.String("in", "", "Binary snapshot.  Optionally send stdin")

func main() {
	flag.Parse()

	var err error
	var inStream io.ReadCloser
	if inFilePtr != nil {
		if inStream, err = os.Open(*inFilePtr); err != nil {
			panic(err)
		}
	} else {
		inStream = os.Stdin
	}
	in := browse.NewBinaryReader(inStream).Node()
	c := node.NewContext()
	s, err := browse.RestoreSelection(c, in)
	if err != nil {
		panic(err)
	}
	err = c.Selector(s).InsertInto(node.NewJsonWriter(os.Stdout).Node()).LastErr
	if err != nil {
		panic(err)
	}
}
