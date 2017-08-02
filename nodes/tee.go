package nodes

import (
	"fmt"

	"context"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

// when writing values, splits output into two nodes.
// when reading, reads from secondary only when primary returns nil
type Tee struct {
	A node.Node
	B node.Node
}

func (self Tee) String() string {
	return fmt.Sprintf("Tee(%s,%s)", self.A.String(), self.B.String())
}

func (self Tee) Child(r node.ChildRequest) (node.Node, error) {
	var err error
	var child Tee
	if child.A, err = self.A.Child(r); err != nil {
		return nil, err
	}
	if child.B, err = self.B.Child(r); err != nil {
		return nil, err
	}
	if child.A != nil && child.B != nil {
		return child, nil
	}
	return nil, nil
}

func (self Tee) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	var err error
	var next Tee
	key := r.Key
	if next.A, key, err = self.A.Next(r); err != nil {
		return nil, nil, err
	}
	if next.B, _, err = self.B.Next(r); err != nil {
		return nil, nil, err
	}
	if next.A != nil && next.B != nil {
		return next, key, nil
	}
	return nil, nil, nil
}

func (self Tee) Field(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
	if r.Write {
		if err = self.A.Field(r, hnd); err == nil {
			err = self.B.Field(r, hnd)
		}
	} else {
		// merging results, prefer first
		if err = self.A.Field(r, hnd); err == nil && hnd.Val == nil {
			err = self.B.Field(r, hnd)
		}
	}
	return
}

func (self Tee) Choose(sel node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	return self.A.Choose(sel, choice)
}

func (self Tee) Action(r node.ActionRequest) (output node.Node, err error) {
	return self.A.Action(r)
}

func (self Tee) Notify(r node.NotifyRequest) (closer node.NotifyCloser, err error) {
	return self.A.Notify(r)
}

func (self Tee) Peek(sel node.Selection, consumer interface{}) interface{} {
	if v := self.A.Peek(sel, consumer); v != nil {
		return v
	}
	return self.B.Peek(sel, consumer)
}

func (self Tee) BeginEdit(r node.NodeRequest) (err error) {
	if err = self.A.BeginEdit(r); err == nil {
		err = self.B.BeginEdit(r)
	}
	return
}

func (self Tee) EndEdit(r node.NodeRequest) (err error) {
	if err = self.A.EndEdit(r); err == nil {
		err = self.B.EndEdit(r)
	}
	return
}

func (self Tee) Delete(r node.NodeRequest) (err error) {
	if err = self.A.Delete(r); err == nil {
		err = self.B.Delete(r)
	}
	return
}

func (self Tee) Context(s node.Selection) context.Context {
	s.Context = self.A.Context(s)
	return self.B.Context(s)
}
