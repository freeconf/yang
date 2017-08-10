package node_test

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func TestCoerseValue(t *testing.T) {
	v, err := node.NewValue(meta.NewDataType(nil, "int32"), 35)
	if err != nil {
		t.Error(err)
	} else if v.Value().(int) != 35 {
		t.Error("Coersion error")
	}
}
