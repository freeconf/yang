package node_test

import (
	"testing"

	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/node"
)

func TestCoerseValue(t *testing.T) {
	dt := meta.NewDataType("int32")
	meta.Validate(dt)
	v, err := node.NewValue(dt, 35)
	if err != nil {
		t.Error(err)
	} else if v.Value().(int) != 35 {
		t.Error("Coersion error")
	}
}
