package node_test

import (
	"encoding/json"
	"testing"

	"github.com/freeconf/c2g/node"
)

func TestMapValue(t *testing.T) {
	var err error
	dataJson := `{"a":{"b":{"x":"waldo"}},"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`
	var data map[string]interface{}
	if err = json.Unmarshal([]byte(dataJson), &data); err != nil {
		t.Error(err)
	}
	if node.MapValue(data, "a.b.x") != "waldo" {
		t.Error("can't find waldo")
	}
	if node.MapValue(data, "p.1.k") != "waldo" {
		t.Error("can't find waldo")
	}
}
