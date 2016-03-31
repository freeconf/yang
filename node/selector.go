package node

import (
	"github.com/c2g/c2"
	"net/url"
	"github.com/c2g/meta"
	"strconv"
	"strings"
)

type Selector struct {
	Context     *Context
	Selection   *Selection
	Constraints *Constraints
	LastErr     error
}

func (self Selector) Find(path string) Selector {
	p := path
	for strings.HasPrefix(p, "../") {
		if self.Selection.parent != nil {
			self.Selection = self.Selection.parent
			p = p[3:]
		} else {
			self.LastErr = c2.NewErrC("No parent path to resolve "+p, 404)
			return self
		}
	}
	var u *url.URL
	u, self.LastErr = url.Parse(p)
	if self.LastErr != nil {
		return self
	}
	return self.FindUrl(u)
}

func (self Selector) FindUrl(url *url.URL) Selector {
	if self.LastErr != nil || url == nil {
		return self
	}
	var targetSlice PathSlice
	targetSlice, self.LastErr = ParseUrlPath(url, self.Selection.Meta())
	if self.LastErr != nil {
		return self
	}
	if len(url.Query()) > 0 {
		buildConstraints(&self, url.Query())
		if self.LastErr != nil {
			return self
		}
	}
	findController := &FindTarget{
		Path : targetSlice,
		Constraints: self.Constraints,
	}
	if self.LastErr = self.Selection.Walk(self.Context, findController); self.LastErr == nil {
		self.Selection = findController.Target
	}
	return self
}

func (self Selector) Constrain(params string) Selector {
	if self.LastErr != nil {
		return self
	}
	if dummy, err := url.Parse("bogus?" + params); err != nil {
		self.LastErr = err
		return self
	} else {
		buildConstraints(&self, dummy.Query())
	}
	return self
}

func buildConstraints(self *Selector, params map[string][]string) {
	self.Constraints = NewConstraints(self.Constraints)
	if _, auto := params["autocreate"]; auto {
		self.Constraints.AddConstraint("autocreate", 50, 50, AutoCreate{})
	}
	depth := self.Selection.path.Len()
	maxDepth := MaxDepth{InitialDepth: depth, MaxDepth: 32}
	if n, found := findIntParam(params, "depth"); found {
		maxDepth.MaxDepth = n
	}
	self.Constraints.AddConstraint("depth", 10, 50, maxDepth)
	if p, found := params["c2-range"]; found {
		if listSelector, selectorErr := NewListRange(self.Selection.path, p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			self.Constraints.AddConstraint("c2-range", 20, 50, listSelector)
		}
	}
	if p, found := params["fields"]; found {
		if listSelector, selectorErr := NewFieldsMatcher(self.Selection.path, p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			self.Constraints.AddConstraint("c2-range", 10, 50, listSelector)
		}
	}
	maxNode := MaxNode{Max: 10000}
	if n, found := findIntParam(params, "c2-max-node-count"); found {
		maxNode.Max = n
	}
	self.Constraints.AddConstraint("c2-max-node-count", 10, 60, maxNode)
}

func findIntParam(params map[string][]string, param string) (int, bool) {
	if v, found := params[param]; found {
		if n, err := strconv.Atoi(v[0]); err == nil {
			return n, true
		}
	}
	return 0, false
}

func (self Selector) InsertInto(toNode Node) Selector {
	return self.edit(false, toNode, INSERT)
}

func (self Selector) InsertFrom(fromNode Node) Selector {
	return self.edit(true, fromNode, INSERT)
}

func (self Selector) UpsertInto(toNode Node) Selector {
	return self.edit(false, toNode, UPSERT)
}

func (self Selector) UpsertFrom(toNode Node) Selector {
	return self.edit(true, toNode, UPSERT)
}

func (self Selector) UpdateInto(toNode Node) Selector {
	return self.edit(false, toNode, UPDATE)
}

func (self Selector) UpdateFrom(toNode Node) Selector {
	return self.edit(true, toNode, UPDATE)
}

func (self Selector) edit(pull bool, n Node, strategy Strategy) Selector {
	if self.LastErr != nil {
		return self
	}
	if self.Selection == nil {
		self.LastErr = c2.NewErrC("No selection", 404)
	}
	var e *Editor
	if pull {
		e = &Editor{
			from: self.Selection.Fork(n),
			to:   self.Selection,
		}
	} else {
		e = &Editor{
			from: self.Selection,
			to:   self.Selection.Fork(n),
		}

	}
	cntlr := &ControlledWalk{
		Constraints: self.Constraints,
	}
	self.LastErr = e.Edit(self.Context, strategy, cntlr)
	return self
}

func (self Selector) Action(input Node) Selector {
	if self.LastErr != nil {
		return self
	}
	r := ActionRequest{
		Request: Request{
			Context:   self.Context,
			Selection: self.Selection,
			Constraints: self.Constraints,
		},
		Meta: self.Selection.Meta().(*meta.Rpc),
	}
	r.Input = Select(r.Meta.Input, input)
	rpcOutput, rerr := self.Selection.node.Action(r)
	if rerr != nil {
		self.LastErr = rerr
		return self
	}
	if rpcOutput != nil {
		self.Selection = Select(r.Meta.Output, rpcOutput)
	} else {
		// legit - rpc has no output
		self.Selection = nil
	}
	return self
}

func (self Selector) Set(ident string, value interface{}) error {
	if self.LastErr != nil {
		return self.LastErr
	}
	n := self.Selection.node
	if cw, ok := n.(ChangeAwareNode); ok {
		n = cw.Changes()
	}
	pos := meta.FindByIdent2(self.Selection.path.meta, ident)
	if pos == nil {
		return c2.NewErrC("property not found "+ident, 404)
	}
	m := pos.(meta.HasDataType)
	v, e := SetValue(m.GetDataType(), value)
	if e != nil {
		return e
	}
	r := FieldRequest{
		Request: Request{
			Context:   self.Context,
			Selection: self.Selection,
			Constraints: self.Constraints,
		},
		Meta: m,
	}
	return n.Write(r, v)
}

func (self Selector) Get(ident string) (interface{}, error) {
	if self.LastErr != nil {
		return nil, self.LastErr
	}
	v, e := self.GetValue(ident)
	if e != nil {
		return nil, e
	}
	return v.Value(), nil
}

func (self Selector) GetValue(ident string) (*Value, error) {
	if self.LastErr != nil {
		return nil, self.LastErr
	}
	pos := meta.FindByIdent2(self.Selection.path.meta, ident)
	if pos == nil {
		return nil, c2.NewErrC("property not found "+ident, 404)
	}
	if !meta.IsLeaf(pos) {
		return nil, c2.NewErrC("property is not a leaf "+ident, 400)
	}
	r := FieldRequest{
		Request: Request{
			Context:   self.Context,
			Selection: self.Selection,
			Constraints: self.Constraints,
		},
		Meta: pos.(meta.HasDataType),
	}
	v, err := self.Selection.node.Read(r)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (self Selector) Select(n Node) Selector {
	self.Selection = self.Selection.Fork(n)
	return self
}
