package node

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"net/url"
	"strconv"
	"strings"
)

// Used to operate on a Selection including finding other selection.  Unlike Selection
// a Selector can have constraints for various needs
type Selector struct {
	Selection   *Selection
	constraints *Constraints
	handler     *ConstraintHandler
	LastErr     error
}

// Handler let's you alter what happens when a contraints finds an error
func (self Selector) Handler() *ConstraintHandler {
	return self.handler
}

// Constraints hold list of things to check when walking or editing a node.
func (self Selector) Constraints() *Constraints {
	return self.constraints
}

// Find navigates to another selector.  Constraints will be same. Cannot include
// url parameters, for that use FindUrl
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

// FindUrl navigates to another selection with possible constraints as url parameters.  Constraints
// are added to any existing contraints.  Original selector and constraints will remain unaltered
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
		Path: targetSlice,
		WalkConstraints: self.constraints,
		WalkConstraintsHandler: self.handler,
	}
	if self.LastErr = self.Selection.Walk(findController); self.LastErr == nil {
		self.Selection = findController.Target
	}
	return self
}

// Apply constraints in the form of url parameters.
// Original selector and constraints will remain unaltered
// Example:
//     sel2 = sel.Constrain("content=config&depth=4")
//  sel will not have content or depth constraints applies, but sel 2 will
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
	constraints := NewConstraints(self.constraints)
	if _, auto := params["autocreate"]; auto {
		constraints.AddConstraint("autocreate", 50, 50, AutoCreate{})
	}
	depth := self.Selection.path.Len()
	maxDepth := MaxDepth{InitialDepth: depth, MaxDepth: 32}
	if n, found := findIntParam(params, "depth"); found {
		maxDepth.MaxDepth = n
	}
	constraints.AddConstraint("depth", 10, 50, maxDepth)
	if p, found := params["c2-range"]; found {
		if listSelector, selectorErr := NewListRange(self.Selection.path, p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("c2-range", 20, 50, listSelector)
		}
	}
	if p, found := params["fields"]; found {
		if listSelector, selectorErr := NewFieldsMatcher(self.Selection.path, p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("fields", 10, 50, listSelector)
		}
	}
	maxNode := MaxNode{Max: 10000}
	if n, found := findIntParam(params, "c2-max-node-count"); found {
		maxNode.Max = n
	}
	constraints.AddConstraint("c2-max-node-count", 10, 60, maxNode)

	if p, found := params["content"]; found {
		if c, err := NewContentConstraint(self.Selection.path, p[0]); err != nil {
			self.LastErr = err
		} else {
			constraints.AddConstraint("content", 10, 70, c)
		}
	}

	self.constraints = constraints
}

func findIntParam(params map[string][]string, param string) (int, bool) {
	if v, found := params[param]; found {
		if n, err := strconv.Atoi(v[0]); err == nil {
			return n, true
		}
	}
	return 0, false
}

// Copy current node into given node.  If there are any existing containers of list
// items then this will fail by design.
func (self Selector) InsertInto(toNode Node) Selector {
	return self.edit(false, toNode, INSERT)
}

// Copy given node into current node.  If there are any existing containers of list
// items then this will fail by design.
func (self Selector) InsertFrom(fromNode Node) Selector {
	return self.edit(true, fromNode, INSERT)
}

// Merge current node into given node.  If there are any existing containers of list
// items then data will be merged.
func (self Selector) UpsertInto(toNode Node) Selector {
	return self.edit(false, toNode, UPSERT)
}

// Merge given node into current node.  If there are any existing containers of list
// items then data will be merged.
func (self Selector) UpsertFrom(toNode Node) Selector {
	return self.edit(true, toNode, UPSERT)
}

// Copy current node into given node.  There must be matching containers of list
// items or this will fail by design.
func (self Selector) UpdateInto(toNode Node) Selector {
	return self.edit(false, toNode, UPDATE)
}

// Copy given node into current node.  There must be matching containers of list
// items or this will fail by design.
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
		Constraints: self.constraints,
		Handler: self.handler,
	}
	self.LastErr = e.Edit(strategy, cntlr)
	return self
}

// Notifications let's caller subscribe to a node.  Node must be a 'notification' node.
func (self Selector) Notifications(stream NotifyStream) (NotifyCloser, Selector) {
	if self.LastErr != nil {
		return nil, self
	}
	r := NotifyRequest{
		Request: Request{
			Selection: self.Selection,
		},
		Meta: self.Selection.Meta().(*meta.Notification),
		Stream: stream,
	}
	var closer NotifyCloser
	closer, self.LastErr = self.Selection.node.Notify(r)
	return closer, self
}

// Action let's to call a procedure potentially passing on data and potentially recieving
// data back.
func (self Selector) Action(input Node) Selector {
	if self.LastErr != nil {
		return self
	}
	r := ActionRequest{
		Request: Request{
			Selection: self.Selection,
		},
		Meta: self.Selection.Meta().(*meta.Rpc),
	}
	r.Input = self.Selection.SelectChild(r.Meta.Input, input)

	if self.constraints != nil {
		r.Constraints = self.constraints
		r.ConstraintsHandler = self.handler
		if proceed, constraintErr := self.constraints.CheckActionPreConstraints(&r); !proceed || constraintErr != nil {
			self.LastErr = constraintErr
			return self
		}
	}

	rpcOutput, rerr := self.Selection.node.Action(r)
	if rerr != nil {
		self.LastErr = rerr
		return self
	}
	if rpcOutput != nil {
		self.Selection = self.Selection.SelectChild(r.Meta.Output, rpcOutput)
	} else {
		// legit - rpc has no output
		self.Selection = nil
	}

	if self.constraints != nil {
		r.Constraints = self.constraints
		r.ConstraintsHandler = self.handler
		if proceed, constraintErr := self.constraints.CheckActionPostConstraints(r); !proceed || constraintErr != nil {
			self.LastErr = constraintErr
			return self
		}
	}

	return self
}

// Set let's you set a leaf value on a container or list item.
func (self Selector) Set(ident string, value interface{}) error {
	if self.LastErr != nil {
		return self.LastErr
	}
	n := self.Selection.node
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
			Selection: self.Selection,
		},
		Write: true,
		Meta: m,
	}
	return n.Field(r, &ValueHandle{Val:v})
}

// Get let's you get a leaf value from a container or list item
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

// GetValue let's you get the leaf value as a Value instance.  Returns null if value is null
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
			Selection: self.Selection,
		},
		Meta: pos.(meta.HasDataType),
	}
	var hnd ValueHandle
	err := self.Selection.node.Field(r, &hnd)
	if err != nil {
		return nil, err
	}
	if hnd.Val != nil {
		hnd.Val.Type = r.Meta.GetDataType()
	} else {
		if r.Meta.GetDataType().HasDefault() {
			hnd.Val = &Value{Type:r.Meta.GetDataType()}
			hnd.Val.CoerseStrValue(r.Meta.GetDataType().Default())
		}
	}
	return hnd.Val, nil
}

// Divert let's you clone the selection state but alter the data source.
func (self Selector) Divert(n Node) Selector {
	self.Selection = self.Selection.Fork(n)
	return self
}
