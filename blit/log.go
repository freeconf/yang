package blit

import (
	"log"
	"os"
)

// Don't send enything to this logger unless it's an error that deserves immediate attention
var Err *log.Logger

// Warnings, or general information.
var Info *log.Logger

var Debug *log.Logger

func init() {
	Err = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	Debug = log.New(os.Stdout, "", log.Lshortfile)
}
