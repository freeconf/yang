package c2

import (
	"reflect"
	"fmt"
)

// Mostly useful in unit tests
func CheckEqual(a interface{}, b interface{}) (error) {
	if ! reflect.DeepEqual(a, b) {
		return NewErr(fmt.Sprintf("\nExpected:'%v'\n  Actual:'%v'", a, b))
	}
	return nil
}