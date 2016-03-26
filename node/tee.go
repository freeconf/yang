package node

import (
	"github.com/blitter/meta"
	"fmt"
)


// when writing values, splits output into two nodes.
// when reading, reads from secondary only when primary returns nil
type Tee struct {
	A Node
	B Node
}

func (self Tee) String() string {
	return fmt.Sprintf("Tee(%s,%s)", self.A.String(), self.B.String())
}

func (self Tee) Select(r ContainerRequest) (Node, error) {
	var err error
	var child Tee
	if child.A, err = self.A.Select(r); err != nil {
		return nil, err
	}
	if child.B, err = self.B.Select(r); err != nil {
		return nil, err
	}
	if child.A != nil && child.B != nil {
		return child, nil
	}
	return nil, nil
}

func (self Tee)  Next(r ListRequest) (Node, []*Value, error) {
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

func (self Tee) Read(r FieldRequest) (*Value, error) {
	// merging results, prefer first
	if v, err := self.A.Read(r); err != nil {
		return nil, err
	} else if v != nil {
		return v, nil
	}
	return self.B.Read(r)
}

func (self Tee) Write(r FieldRequest, val *Value) (err error) {
	if err = self.A.Write(r, val); err == nil {
		err = self.B.Write(r, val)
	}
	return
}

func (self Tee) Choose(sel *Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	return self.A.Choose(sel, choice)
}

func (self Tee) Event(sel *Selection, e Event) (err error) {
	if err = self.A.Event(sel, e); err == nil {
		err = self.B.Event(sel, e)
	}
	return
}

func (self Tee) Action(r ActionRequest) (output Node, err error) {
	return self.A.Action(r)
}

func (self Tee) Peek(sel *Selection, peekId string) interface{} {
	if v := self.A.Peek(sel, peekId); v != nil {
		return v
	}
	return self.B.Peek(sel, peekId)
}
