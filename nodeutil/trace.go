package nodeutil

import (
	"fmt"
	"io"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

const tpadding = "                                                                                       "

type trace struct {
	Out io.Writer
}

// Trace will record all calls and data into a target node and recursively wrap each level to effectively
// trace all operations on a node and it's children
func Trace(target node.Node, out io.Writer) node.Node {
	if target == nil {
		fmt.Fprintln(out, "node: nil")
		return nil
	}
	return trace{
		Out: out,
	}.Node(0, target)
}

func (t trace) trace(level int, key string, v interface{}) {
	_, err := fmt.Fprintf(t.Out, "%s%s: %v\n", tpadding[:(4*level)], key, v)
	if err != nil {
		panic(err)
	}
}

func (t trace) chkerr(level int, err error) error {
	if err != nil {
		t.trace(level, "err", err)
	}
	return err
}

func (t trace) traceOnTrue(level int, key string, flag bool) {
	if flag {
		t.trace(level, key, true)
	}
}

func (t trace) traceVals(level int, key string, vals []val.Value) {
	for i, v := range vals {
		t.traceVal(level, fmt.Sprintf("%s[%d]", key, i), v)
	}
}

func (t trace) traceVal(level int, key string, val val.Value) {
	if val == nil {
		t.trace(level, key, nil)
	} else {
		t.trace(level, key, fmt.Sprintf("%s(%s)", val.Format(), val))
	}
}

func (t trace) ident(p *node.Path) string {
	if p.Key != nil {
		var strs []string
		for _, k := range p.Key {
			strs = append(strs, k.String())
		}
		return fmt.Sprintf("%s=%s", p.Meta.Ident(), strings.Join(strs, ","))
	}
	return p.Meta.Ident()
}

func (t trace) Node(level int, target node.Node) node.Node {
	n := &Basic{}
	n.OnPeek = target.Peek
	n.OnAction = target.Action
	n.OnNotify = target.Notify
	n.OnChoose = func(sel node.Selection, choice *meta.Choice) (choosen *meta.ChoiceCase, err error) {
		t.trace(level, "choose", choice.Ident())
		choosen, err = target.Choose(sel, choice)
		if choosen != nil {
			t.trace(level+1, "choosen", choosen.Ident())
		} else {
			t.trace(level+1, "choosen", "nil")
		}
		t.chkerr(level+1, err)
		return choosen, err
	}
	n.OnChild = func(r node.ChildRequest) (child node.Node, err error) {
		if r.New {
			t.trace(level, "child.new", t.ident(r.Path))
		} else if r.Delete {
			t.trace(level, "child.delete", t.ident(r.Path))
		} else {
			t.trace(level, "child.read", t.ident(r.Path))
		}
		child, err = target.Child(r)
		t.trace(level+1, "found", child != nil)
		t.chkerr(level+1, err)
		if child == nil {
			return nil, err
		}
		return t.Node(level+1, child), nil
	}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			t.trace(level, "field.write", r.Meta.Ident())
			t.traceVal(level+1, "val", hnd.Val)
			err = t.chkerr(level+1, target.Field(r, hnd))
		} else {
			t.trace(level, "field.read", r.Meta.Ident())
			err = t.chkerr(level+1, target.Field(r, hnd))
			t.traceVal(level+1, "val", hnd.Val)
		}
		return err
	}
	n.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		if r.New {
			t.trace(level, fmt.Sprintf("next.new[%d]", r.Row), t.ident(r.Path))
		} else if r.Delete {
			t.trace(level, fmt.Sprintf("next.delete[%d]", r.Row), t.ident(r.Path))
		} else {
			t.trace(level, fmt.Sprintf("next.read[%d]", r.Row), t.ident(r.Path))
		}
		if r.Key != nil {
			t.traceVals(level+1, "request.key", r.Key)
		}
		next, key, err = target.Next(r)
		t.chkerr(level+1, err)
		t.trace(level+1, "found", next != nil)
		if !r.New {
			t.traceVals(level+1, "response.key", key)
		}
		if next != nil {
			return t.Node(level+1, next), key, err
		}
		return nil, key, err
	}
	n.OnBeginEdit = func(r node.NodeRequest) (err error) {
		t.trace(level, "edit.begin", t.ident(r.Selection.Path))
		t.traceOnTrue(level+1, "new", r.New)
		t.traceOnTrue(level+1, "delete", r.Delete)
		return t.chkerr(level, target.BeginEdit(r))
	}
	n.OnEndEdit = func(r node.NodeRequest) (err error) {
		t.trace(level, "edit.end", t.ident(r.Selection.Path))
		t.traceOnTrue(level+1, "new", r.New)
		t.traceOnTrue(level+1, "delete", r.Delete)
		return t.chkerr(level, target.EndEdit(r))
	}
	return n
}
