package c2

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Don't send enything to this logger unless it's an error that deserves immediate attention
var Err *log.Logger

// Warnings, or general information.
var Info *log.Logger

var Debug *log.Logger

func init() {
	Err = log.New(os.Stderr, "ERRO ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	Debug = log.New(ioutil.Discard, "DEBG ", log.Lshortfile)
}

var debugLogEnabled bool

func DebugLog(e bool) {
	if debugLogEnabled == e {
		return
	}
	debugLogEnabled = e
	var f io.Writer
	if debugLogEnabled {
		f = os.Stdout
	} else {
		f = ioutil.Discard
	}
	Debug = log.New(f, "DEBG ", log.Lshortfile)
}

func DebugLogEnabled() bool {
	return debugLogEnabled
}
