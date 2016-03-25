package node
import (
	"testing"
)

func TestExtend(t *testing.T) {
	child := &MyNode{Label: "Bloop"}
	n := &MyNode{
		Label: "Blop",
		OnRead: func(FieldRequest) (*Value, error) {
			return &Value{Str:"Hello"}, nil
		},
		OnSelect: func(r ContainerRequest) (Node, error) {
			return child, nil
		},
	}
	x := Extend{
		Label: "Bleep",
		Node: n,
		OnRead: func(p Node, r FieldRequest) (*Value, error) {
			v, _ := p.Read(r)
			return &Value{Str:v.Str + " World"}, nil
		},
	}
	actualValue, _  := x.Read(FieldRequest{})
	if actualValue.Str != "Hello World" {
		t.Error(actualValue.Str)
	}
	if x.String() != "(Blop) <- Bleep" {
		t.Error(x.String())
	}
	actualChild, _ := x.Select(ContainerRequest{})
	if actualChild.String() != "Bloop" {
		t.Error(actualChild.String())
	}
}