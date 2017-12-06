package meta

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/freeconf/c2g/c2"
)

type SetDescription string
type SetConfig bool
type SetMandatory bool
type SetUnbounded bool
type SetLenRange Range
type SetValueRange Range
type SetReference string
type SetPrefix string
type SetNamespace string
type SetOrganization string
type SetContact string
type SetKey string
type SetPath string
type SetPattern string
type SetDefault struct {
	Value interface{}
}
type SetMaxElements int
type SetMinElements int
type SetEnumValue int
type SetLength string
type SetBase string
type SetYangVersion string
type SetYinElement bool
type SetUnits string
type SetFractionDigits int

func Set(parent interface{}, prop interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c2.NewErr(fmt.Sprintf("failed to set %T.%T : %s", parent, prop, r))
		}
	}()
	parent.(buildable).add(prop)
	return
}

func Validate(m Meta) error {
	schemaPool := make(schemaPool)
	if err := m.(resolver).resolve(schemaPool); err != nil {
		return err
	}
	return m.(compilable).compile()
}

func (encoded SetKey) decode() []string {
	return strings.Split(string(encoded), " ")
}

type Range struct {
	Min    int
	Max    int
	HasMax bool
}

func (r Range) Empty() bool {
	return r.Min == 0 && r.Max == 0
}

func NewRange(encoded string) (r Range, err error) {
	segments := strings.Split(string(encoded), "..")
	if len(segments) == 2 {
		if r.Min, err = strconv.Atoi(segments[0]); err != nil {
			return
		}
		if r.Max, err = strconv.Atoi(segments[1]); err != nil {
			return
		}
		r.HasMax = true
	} else {
		if r.Max, err = strconv.Atoi(segments[0]); err != nil {
			return
		}
		r.HasMax = true
	}
	return
}

// // DecodeLengths will decode min and max lengths formated according to
// // RFC.  Example: 1..10 where 1 is min and 30 is max.
// func (encoded SetEncodedLength) decode(c buildable) {
// 	/* TODO: Support multiple lengths using "|" */
// 	segments := strings.Split(string(encoded), "..")
// 	if len(segments) == 2 {
// 		if minLength, err := strconv.Atoi(segments[0]); err != nil {
// 			panic(err.Error())
// 		} else {
// 			c.add(SetMinLength(minLength))
// 		}
// 		if maxLength, err := strconv.Atoi(segments[1]); err != nil {
// 			panic(err.Error())
// 		} else {
// 			c.add(SetMaxLength(maxLength))
// 		}
// 	} else {
// 		if maxLength, err := strconv.Atoi(segments[0]); err != nil {
// 			panic(err.Error())
// 		} else {
// 			c.add(SetMaxLength(maxLength))
// 		}
// 	}
// }
