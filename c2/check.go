package c2

import (
	"reflect"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"errors"
	"bytes"
)

// Mostly useful in unit tests
func CheckEqual(a interface{}, b interface{}) (error) {
	if ! reflect.DeepEqual(a, b) {
		return NewErr(fmt.Sprintf("\nExpected:'%v'\n  Actual:'%v'", a, b))
	}
	return nil
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
	diffPath, noDiff := exec.LookPath("diff")
	if noDiff != nil {
		panic(noDiff.Error())
	}
	cmd := exec.Command(diffPath, "-U", "3", a, f.Name())
	var outBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Run()
	if ! cmd.ProcessState.Success() {
		//outData, _ := cmd.Output()
		return errors.New(outBuff.String())
	}
	return nil
}
