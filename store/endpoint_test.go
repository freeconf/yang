// +build ignore

package store

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func TestEndpointProxy(t *testing.T) {
	//	oper := node.NewJsonReader(strings.NewReader(`{"stats":{"crumbTrayCapacity":98}}`)).Node()
	//	src := &DummyProxySource{
	//		Module: c2.Yang("c2-toaster"),
	//		Node: oper,
	//	}
	//	store := node.NewBufferStore()
	//	store.Values["toast/darknessLevel"] = &node.Value{Int:99}
	//	registrar, err := NewDataProxy(src, store)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	var actual bytes.Buffer
	//	outNode := node.NewJsonWriter(&actual).Node()
	//	r := node.ListRequest{
	//		Target: node.NewPathSlice("toast", src.Module),
	//	}
	//	proxyNode, proxyNodeErr := registrar.Node(r.Target.Tail, "toast")
	//	if proxyNodeErr != nil {
	//		t.Fatal(proxyNodeErr)
	//	}
	//	sel := node.Select(src.Module, proxyNode)
	//	if err = sel.Find("toast").Push(outNode).Insert().LastErr; err != nil {
	//		t.Fatal(err)
	//	}
	//	expected := `{"darknessLevel":99,"stats":{"crumbTrayCapacity":98}}`
	//	if expected != actual.String() {
	//		t.Error("\nExpected:%s\n  Actual:%s", expected, actual.String())
	//	}
}

type DummyProxySource struct {
	Module *meta.Module
	Node   node.Node
}

func (self *DummyProxySource) Schema() (*meta.Module, error) {
	return self.Module, nil
}

func (self *DummyProxySource) OperationalNode(target string) (node.Node, error) {
	return self.Node, nil
}
