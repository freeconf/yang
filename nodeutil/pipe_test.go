package nodeutil

import (
	"errors"
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/val"
)

func TestPipeLeaf(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("m", nil)
	pull, push := NewPipe().PullPush()
	aValue := val.String("A")
	aReq := node.FieldRequest{
		Meta: b.Leaf(m, "a"),
	}
	bReq := node.FieldRequest{
		Meta: b.Leaf(m, "b"),
	}
	go func() {
		aReq.Write = true
		push.Field(aReq, &node.ValueHandle{Val: aValue})
	}()
	var actualB, actualA node.ValueHandle
	errB := pull.Field(bReq, &actualB)
	if errB != nil {
		t.Error(errB)
	}
	if actualB.Val != nil {
		t.Error("B shouldn't exist")
	}
	aReq.Write = false
	errA := pull.Field(aReq, &actualA)
	if errA != nil {
		t.Error(errA)
	}
	if actualA.Val == nil {
		t.Error("A should exist")
	}
}

var pipeTestModule = `
module m {
	namespace "";
	prefix "";
	revision 0;
	leaf c {
		type string;
	}
	container a {
		container b {
			leaf x {
				type string;
			}
		}
	}
	list p {
		key "k";
		leaf k {
			type string;
		}
		container q {
			leaf s {
				type string;
			}
		}
		list r {
			leaf z {
				type int32;
			}
		}
	}
}
`

func TestPipeFull(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, pipeTestModule)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  `{"c":"hello"}`,
			out: `{"c":"hello"}`,
		},
		{
			in:  `{"a":{"b":{"x":"waldo"}}}`,
			out: `{"a":{"b":{"x":"waldo"}}}`,
		},
		{
			in:  `{"p":[{"k":"walter"}]}`,
			out: `{"p":[{"k":"walter"}]}`,
		},
		{
			in:  `{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
			out: `{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
		},
	}
	for _, test := range tests {
		pipe := NewPipe()
		pull, push := pipe.PullPush()

		go func() {
			sel := node.NewBrowser(m, push).Root()
			pipe.Close(sel.InsertFrom(ReadJSON(test.in)))
		}()
		actual, err := WriteJSON(node.NewBrowser(m, pull).Root())
		if err != nil {
			t.Error(err)
		} else if actual != test.out {
			t.Errorf("\nExpected:%s\n  Actual:%s", test, actual)
		}
	}
}

func TestPipeErrorHandling(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, pipeTestModule)
	if err != nil {
		t.Fatal(err)
	}
	pipe := NewPipe()
	pull, push := pipe.PullPush()
	hasProblems := &Basic{
		OnChild: func(node.ChildRequest) (node.Node, error) {
			return nil, errors.New("planned error in select")
		},
		OnField: func(node.FieldRequest, *node.ValueHandle) error {
			return errors.New("planned error in read")
		},
	}
	go func() {
		sel := node.NewBrowser(m, push).Root()
		pipe.Close(sel.InsertFrom(hasProblems))
	}()
	err = node.NewBrowser(m, pull).Root().InsertInto(&Basic{})
	if err == nil {
		t.Error("Expected error")
	}
}
