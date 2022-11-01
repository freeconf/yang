package fc

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

// AssertEqual emits testing error if a and b are not equal. Returns true if
// equal
func AssertEqual(t Tester, a interface{}, b interface{}, msgs ...string) bool {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		err := fmt.Errorf("\nExpected:'%v'\n  Actual:'%v' %s", a, b, strings.Join(msgs, " "))
		t.Error(err)
		return false
	}
	return true
}

// DiffBytes will compare two byte arrays and emit formatted difference
// useful in "Golden File Testing", return true if no differences
func DiffBytes(t Tester, a []byte, b []byte) bool {
	t.Helper()
	f, fErr := ioutil.TempFile(os.TempDir(), "tst")
	if fErr != nil {
		panic(fErr.Error())
	}
	defer os.Remove(f.Name())
	if _, wErr := f.Write(a); wErr != nil {
		panic(wErr)
	}
	return Diff(t, a, f.Name())
}

// Diff will compare byte array to file and emit formatted difference
// useful in "Golden File Testing", return true if no differences
func Diff(t Tester, a []byte, b string) bool {
	t.Helper()
	f, fErr := ioutil.TempFile(os.TempDir(), "tst")
	if fErr != nil {
		panic(fErr.Error())
	}
	defer os.Remove(f.Name())
	if _, wErr := f.Write(a); wErr != nil {
		panic(wErr)
	}
	return DiffFiles(t, f.Name(), b)
}

// DiffFiles will compare two files and emit formatted difference
// useful in "Golden File Testing", return true if no differences
func DiffFiles(t Tester, a string, b string) bool {
	t.Helper()
	for _, f := range []string{a, b} {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			panic(f + " does not exist")
		}
	}
	cmd := exec.Command("diff", "-U", "3", b, a)
	var outBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Run()
	if !cmd.ProcessState.Success() {
		t.Error(errors.New(outBuff.String()))
		return false
	}
	return true
}

// Gold compares bytes to a the contents of a file on disk UNLESS update flag
// is passed, then it replaces contents of file on disk. This testing strategy
// if known as "gold files" and can be found in many projects including the Go SDK
func Gold(t Tester, update bool, actual []byte, gfile string) bool {
	t.Helper()
	if update {
		if err := ioutil.WriteFile(gfile, actual, 0666); err != nil {
			panic(err)
		}
	} else {
		return Diff(t, actual, gfile)
	}
	return true
}

// Tester is internal. Here just to decouple test utilities from testing.T package.
type Tester interface {
	Helper()
	Error(args ...interface{})
}
