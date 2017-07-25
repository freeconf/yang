package c2

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
)

// Mostly useful in unit tests
type Testing interface {
	Error(err ...interface{})
}

func CheckEqual(a interface{}, b interface{}) error {
	if !reflect.DeepEqual(a, b) {
		return NewErr(fmt.Sprintf("\nExpected:'%v'\n  Actual:'%v'", a, b))
	}
	return nil
}

func Equals(t Testing, a interface{}, b interface{}) bool {
	if !reflect.DeepEqual(a, b) {
		t.Error(errors.New(fmt.Sprintf("\nExpected:'%v'\n  Actual:'%v'", a, b)))
		return false
	}
	return true
}

func Diff(a []byte, b []byte) error {
	f, fErr := ioutil.TempFile(os.TempDir(), "tst")
	if fErr != nil {
		panic(fErr.Error())
	}
	defer os.Remove(f.Name())
	if _, wErr := f.Write(a); wErr != nil {
		panic(wErr)
	}
	return Diff2(f.Name(), b)
}

func Diff2(a string, b []byte) error {
	f, fErr := ioutil.TempFile(os.TempDir(), "tst")
	if fErr != nil {
		panic(fErr.Error())
	}
	defer os.Remove(f.Name())
	if _, wErr := f.Write(b); wErr != nil {
		panic(wErr)
	}
	return DiffFiles(a, f.Name())
}

func DiffFiles(a string, b string) error {
	cmd := exec.Command("diff", "-U", "3", a, b)
	var outBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Run()
	if !cmd.ProcessState.Success() {
		return errors.New(outBuff.String())
	}
	return nil
}
