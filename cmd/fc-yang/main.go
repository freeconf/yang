package main

import (
	"log"
	"os"

	"github.com/freeconf/yang/cmd/fc-yang/doc"
	"github.com/freeconf/yang/cmd/fc-yang/get"
)

func main() {

	// pop out the main command and rewrite the args as if it wasn't there so each
	// sub command can pretend as if was called directly
	if len(os.Args) <= 1 {
		log.Fatal("Usage: [doc, get] ...")
	}
	cmd := os.Args[1]
	if len(os.Args) == 2 {
		os.Args = []string{os.Args[0]}
	} else {
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	}
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
