package meta

import (
	"fmt"
	"strings"

	"github.com/freeconf/yang/c2"
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
type SetSecondaryExtension struct {
	On         string
	Extensions []*Extension
}

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
