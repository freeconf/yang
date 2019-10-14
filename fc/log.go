package fc

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Err is for errors only.  An error is somethhing that deserves immediate attention
var Err *log.Logger

// Info is Warnings, or general information.
var Info *log.Logger

// Debug is discarded unless -debug flag is passed on CLI.
var Debug *log.Logger

func init() {
	Err = log.New(os.Stderr, "ERRO ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	Debug = log.New(ioutil.Discard, "DEBG ", log.Lshortfile)
}

var debugLogEnabled bool

// DebugLog turns on debug level logging
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

// DebugLogEnabled returns whether debug level logging is on or not
func DebugLogEnabled() bool {
	return debugLogEnabled
}
