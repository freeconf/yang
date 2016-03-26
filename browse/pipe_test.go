package browse

import (
	"testing"
	"github.com/blitter/node"
	"github.com/blitter/meta"
	"github.com/blitter/meta/yang"
	"strings"
	"bytes"
	"errors"
)

func TestPipeLeaf(t *testing.T) {
	pull, push := NewPipe().PullPush()
	aValue := &node.Value{Str:"A"}
	aReq := node.FieldRequest{
		Meta: &meta.Leaf{Ident:"a"},
	}
	bReq := node.FieldRequest{
		Meta: &meta.Leaf{Ident:"b"},
	}
	go func() {
		push.Write(aReq, aValue)
	}()
	actualB, errB := pull.Read(bReq)
	if errB != nil {
		t.Error(errB)
	}
	if actualB != nil {
		t.Error("B shouldn't exist")
	}
	actualA, errA := pull.Read(aReq)
	if errA != nil {
		t.Error(errA)
	}
	if actualA == nil {
		t.Error("A should exist")
	}
}

var pipeTestModule =  `
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
	c := node.NewContext()
	for _, test := range tests {
		pipe := NewPipe()
		pull, push := pipe.PullPush()
		in := node.NewJsonReader(strings.NewReader(test)).Node()
		var actualBytes bytes.Buffer
		out := node.NewJsonWriter(&actualBytes).Node()
		go func() {
			pipe.Close(c.Select(m, push).InsertFrom(in).LastErr)
		}()
		if err := c.Select(m, pull).InsertInto(out).LastErr; err != nil {
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
	c := node.NewContext()
	hasProblems := &node.MyNode{
		OnSelect:func(node.ContainerRequest) (node.Node, error) {
			return nil, errors.New("planned error in select")
		},
		OnRead:func(node.FieldRequest) (*node.Value, error) {
			return nil, errors.New("planned error in read")
		},
	}
	go func() {
		pipe.Close(c.Select(m, push).InsertFrom(hasProblems).LastErr)
	}()
	err = c.Select(m, pull).InsertInto(&node.MyNode{}).LastErr
	if err == nil {
		t.Error("Expected error")
	}
}