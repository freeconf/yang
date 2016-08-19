package browse

import (
	"bytes"
	"errors"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"strings"
	"testing"
)

func TestPipeLeaf(t *testing.T) {
	pull, push := NewPipe().PullPush()
	aValue := &node.Value{Str: "A"}
	aReq := node.FieldRequest{
		Meta:  &meta.Leaf{Ident: "a"},
	}
	bReq := node.FieldRequest{
		Meta: &meta.Leaf{Ident: "b"},
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
	m, err := yang.LoadModuleCustomImport(pipeTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []string{
		`{"c":"hello"}`,
		`{"a":{"b":{"x":"waldo"}}}`,
		`{"p":[{"k":"walter"}]}`,
		`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
	}
	for _, test := range tests {
		pipe := NewPipe()
		pull, push := pipe.PullPush()
		in := node.NewJsonReader(strings.NewReader(test)).Node()
		var actualBytes bytes.Buffer
		out := node.NewJsonWriter(&actualBytes).Node()
		go func() {
			sel := node.NewBrowser2(m, push).Root()
			pipe.Close(sel.InsertFrom(in).LastErr)
		}()
		if err := node.NewBrowser2(m, pull).Root().InsertInto(out).LastErr; err != nil {
			t.Error(err)
		}
		actual := actualBytes.String()
		if actual != test {
			t.Error("\nExpected:%s\n  Actual:%s", test, actual)
		}
	}
}

func TestPipeErrorHandling(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(pipeTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	pipe := NewPipe()
	pull, push := pipe.PullPush()
	hasProblems := &node.MyNode{
		OnSelect: func(node.ContainerRequest) (node.Node, error) {
			return nil, errors.New("planned error in select")
		},
		OnField: func(node.FieldRequest, *node.ValueHandle) error {
			return errors.New("planned error in read")
		},
	}
	go func() {
		sel := node.NewBrowser2(m, push).Root()
		pipe.Close(sel.InsertFrom(hasProblems).LastErr)
	}()
	err = node.NewBrowser2(m, pull).Root().InsertInto(&node.MyNode{}).LastErr
	if err == nil {
		t.Error("Expected error")
	}
}
