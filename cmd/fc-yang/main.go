package main

import (
	"log"
	"os"

	"github.com/freeconf/yang/cmd/fc-yang/doc"
	"github.com/freeconf/yang/cmd/fc-yang/get"
)

// fc-yang is your one stop command utility for all things yang.  It's a bit
// overloaded for sure but from an end user deployment perspective, they just
// have one executable to manage which makes things, in the end, easier.  This
// follows in the evolution of go's "go" command that went thru same path.
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: [doc, get] ...")
	}
	cmd := os.Args[1]

	// pop out the main command and rewrite the args as if it wasn't there so each
	// sub command can pretend as if was called directly
	if len(os.Args) == 2 {
		os.Args = []string{os.Args[0]}
	} else {
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	}

	// dispatch to appropriate command
	switch cmd {
	case "doc":
		doc.Run()
	case "get":
		get.Run()
	default:
		log.Fatalf("'%s' is not a recognized command", os.Args[1])
	}
	os.Exit(0)
}
