package node

import (
	"testing"

	"github.com/c2stack/c2g/meta"
)

func TestCoerseValue(t *testing.T) {
	v, err := NewValue(meta.NewDataType(nil, "int32"), 35)
	if err != nil {
		t.Error(err)
	} else if v.Int != 35 {
		t.Error("Coersion error")
	}
}
