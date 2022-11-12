package node

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"context"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Selection give you access to all management operations on an manamagent API.
// It combines a single data node (node.Node) with a single model definition (meta.Meta) to
// represent a single location on a management API (including the root location).
//
// From here you can perform many operations from the underlying nodes including the
// following:
//
//  1. Find/navigate to any other point in the management API
//  2. Get data (i.e. export)
//  2. Write data (i.e perform an edit)
//  3. Run an Action/RPC
//  4. Subscribe to an event stream
//  5. Delete data
//
// You can chan
//
//  Example:
//     var err error
//     root := browser.Root()
//     jay := root.Find("birds=bluejay")
//     myCheckErr(dim.LastErr)
//
//     // write
//     err = jay.UpsertFrom(nodeutil.ReadJSON(`{"dimensions":55}`)).LastErr
//
//     // read
//     err = jay.UpsertInto(someOtherNode)
//
//     // action
//     _, err = jay.Find("fly").Action(nil)
//
//     // subscribe
//     reportRareBird := func(msg node.Selection) {
//	        fmt.Println(nodeutil.ReadJSON(msg))
//     }
//     unsubscribe, err := root.Find("rareSighting").Notify(reportRareBird)
//     // unsubscribe
//     unsubscribe()
//
type Selection struct {

	// Browser that this selection ultimately spawned from
	Browser *Browser

	// Direct parent selection that would have created this selection
	Parent *Selection

	// Underlying node that implements management functions
	Node Node

	// Meta path in YANG module tied to this node
	Path *Path

	// Potentialy stores external data made available to all requests
	Context context.Context

	// Useful when navigating lists, True if this selector is List node, False if
	// this is for an item in List node.
	InsideList bool

	// Constraints hold list of things to check when walking or editing a node.
	Constraints *Constraints

	// When a selection is returned, be sure to check if it resulted in an error. This
	// allows for chaining a few operations and checking the resulting error at the
	// end
	LastErr error
}

func (self Selection) Meta() meta.Definition {
	return self.Path.meta
}

// This selection points nowhere and must have been returned from a function that didn't find
// another selection
func (self Selection) IsNil() bool {
	return self.Path == nil
}

// Create a new independant selection with a different browser from this point in the tree based on a whole
// new data node
func (self Selection) Split(node Node) Selection {
	if self.IsNil() {
		return Selection{
			LastErr: errors.New("selection is nil"),
			Context: self.Context,
		}
	}
	fork := self
	fork.Parent = nil
	fork.Browser = NewBrowser(meta.RootModule(self.Path.meta), node)
	fork.Constraints = &Constraints{}
	fork.Node = node
	return fork
}

// If this is a selection in a list, this is the key value of that list item.
func (self Selection) Key() []val.Value {
	return self.Path.key
}

func (self Selection) selekt(r *ChildRequest) Selection {
	// check pre-constraints
	if proceed, constraintErr := self.Constraints.CheckContainerPreConstraints(r); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: self.Context,
		}
	}

	// select node
	var child Selection
	childNode, err := self.Node.Child(*r)
	if err != nil {
		child = Selection{
			LastErr: err,
			Context: self.Context,
		}
	} else if childNode == nil {
		child = Selection{}
	} else {
		child = Selection{
			Browser:     self.Browser,
			Parent:      &self,
			Path:        &Path{parent: self.Path, meta: r.Meta},
			Node:        childNode,
			Constraints: self.Constraints,
			Context:     self.Context,
		}
		child.Context = childNode.Context(child)
		child.Context = self.Constraints.ContextConstraint(child)
	}

	// check post-constraints
	if proceed, constraintErr := self.Constraints.CheckContainerPostConstraints(*r, child); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: self.Context,
		}
	}

	return child
}

type ListItem struct {
	Selection Selection
	Row       int64
	Key       []val.Value
	req       ListRequest
}

// If at list, this will be iterator into first item in list
func (self Selection) First() ListItem {
	item := ListItem{
		req: ListRequest{
			Request: Request{
				Selection: self,
				Path:      self.Path,
			},
			First: true,
			Meta:  self.Meta().(*meta.List),
		},
	}
	item.Selection, item.Key = self.selectListItem(&item.req)
	return item
}

// iterating a list, get next item in list
func (self ListItem) Next() ListItem {
	self.req.First = false
	self.req.IncrementRow()
	self.Selection, self.Key = self.req.Selection.selectListItem(&self.req)
	return self
}

func (self Selection) selectListItem(r *ListRequest) (Selection, []val.Value) {
	// check pre-constraints
	if proceed, constraintErr := self.Constraints.CheckListPreConstraints(r); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: self.Context,
		}, nil
	}

	// select node
	var child Selection
	childNode, key, err := self.Node.Next(*r)
	if err != nil {
		child = Selection{
			LastErr: err,
			Context: self.Context,
		}
	} else if childNode == nil {
		child = Selection{}
	} else {
		var parentPath *Path
		if self.Parent != nil {
			parentPath = self.Parent.Path
		}
		child = Selection{
			Browser: self.Browser,
			Parent:  &self,
			Node:    childNode,
			// NOTE: Path.parent is lists parentPath, not self.path
			Path:        &Path{parent: parentPath, meta: self.Path.meta, key: key},
			InsideList:  true,
			Constraints: self.Constraints,
			Context:     self.Context,
		}
		child.Context = childNode.Context(child)
		child.Context = self.Constraints.ContextConstraint(child)
	}

	// check post-constraints
	if proceed, constraintErr := self.Constraints.CheckListPostConstraints(*r, child, r.Selection.Path.key); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: self.Context,
		}, nil
	}

	return child, key
}

func (self Selection) Peek(consumer interface{}) interface{} {
	if self.LastErr != nil {
		panic(self.LastErr)
	}
	if self.IsNil() {
		return nil
	}
	return self.Node.Peek(self, consumer)
}

// Apply constraints in the form of url parameters.
// Original selector and constraints will remain unaltered
// Example:
//     sel2 = sel.Constrain("content=config&depth=4")
//  sel will not have content or depth constraints applies, but sel 2 will
func (self Selection) Constrain(params string) Selection {
	if self.LastErr != nil {
		return self
	}
	if dummy, err := url.Parse("bogus?" + params); err != nil {
		self.LastErr = err
		return self
	} else {
		buildConstraints(&self, dummy.Query())
		self.Context = self.Constraints.ContextConstraint(self)
	}
	return self
}

func buildConstraints(self *Selection, params map[string][]string) {
	constraints := NewConstraints(self.Constraints)
	maxDepth := MaxDepth{MaxDepth: 64}
	if n, found := findIntParam(params, "depth"); found {
		maxDepth.MaxDepth = n
	}
	constraints.AddConstraint("depth", 10, 50, maxDepth)
	if p, found := params["fc.range"]; found {
		if listSelector, selectorErr := NewListRange(p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("fc.range", 20, 50, listSelector)
		}
	}
	if p, found := params["fields"]; found {
		if listSelector, selectorErr := NewFieldsMatcher(p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("fields", 10, 50, listSelector)
		}
	}
	if p, found := params["fc.xfields"]; found {
		if listSelector, selectorErr := NewExcludeFieldsMatcher(p[0]); selectorErr != nil {
			self.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("fc.xfields", 10, 50, listSelector)
		}
	}
	maxNode := MaxNode{Max: 10000}
	if n, found := findIntParam(params, "fc.max-node-count"); found {
		maxNode.Max = n
	}
	constraints.AddConstraint("fc.max-node-count", 10, 60, maxNode)

	if p, found := params["content"]; found {
		if c, err := NewContentConstraint(self.Path, p[0]); err != nil {
			self.LastErr = err
		} else {
			constraints.AddConstraint("content", 10, 70, c)
		}
	}

	if p, found := params["with-defaults"]; found {
		if c, err := NewWithDefaultsConstraint(p[0]); err != nil {
			self.LastErr = err
		} else {
			constraints.AddConstraint("with-defaults", 50, 70, c)
		}
	}
	if p, found := params["filter"]; found {
		if c, err := NewFilterConstraint(p[0]); err != nil {
			self.LastErr = err
		} else {
			constraints.AddConstraint("filter", 10, 50, c)
		}
	}

	self.Constraints = constraints
}

func (self Selection) beginEdit(r NodeRequest, bubble bool) error {
	r.Selection = self
	if self.IsNil() {
		return errors.New("selection is nil")
	}
	if err := self.Browser.Triggers.beginEdit(r); err != nil {
		return err
	}
	for {
		if err := r.Selection.Node.BeginEdit(r); err != nil {
			return err
		}
		if r.Selection.Parent == nil || !bubble {
			break
		}
		r.Selection = *r.Selection.Parent
		r.EditRoot = false
	}
	return nil
}

func (self Selection) endEdit(r NodeRequest, bubble bool) error {
	r.Selection = self
	copy := r
	for {
		if err := copy.Selection.Node.EndEdit(copy); err != nil {
			return err
		}
		if copy.Selection.Parent == nil || !bubble {
			break
		}
		copy.Selection = *copy.Selection.Parent
		copy.EditRoot = false
	}
	if err := self.Browser.Triggers.endEdit(r); err != nil {
		return err
	}
	return nil
}

func (self Selection) Delete() (err error) {

	// allow children to recieve indication their parent is being deleted by
	// sending node request w/delete=true
	if err := self.beginEdit(NodeRequest{Source: self, Delete: true, EditRoot: true}, true); err != nil {
		return err
	}

	if self.InsideList {
		r := ListRequest{
			Request: Request{
				Selection: *self.Parent,
			},
			Meta:   self.Meta().(*meta.List),
			Delete: true,
			Key:    self.Key(),
		}
		if _, _, err := r.Selection.Node.Next(r); err != nil {
			return err
		}
	} else {
		r := ChildRequest{
			Request: Request{
				Selection: *self.Parent,
			},
			Meta:   self.Meta().(meta.HasDataDefinitions),
			Delete: true,
		}
		if _, err := r.Selection.Node.Child(r); err != nil {
			return err
		}
	}

	if err := self.endEdit(NodeRequest{Source: self, Delete: true, EditRoot: true}, true); err != nil {
		return err
	}
	return
}

func findIntParam(params map[string][]string, param string) (int, bool) {
	if v, found := params[param]; found {
		if n, err := strconv.Atoi(v[0]); err == nil {
			return n, true
		}
	}
	return 0, false
}

// InsertInto Copy current node into given node.  If there are any existing containers of list
// items then this will fail by design.
func (self Selection) InsertInto(toNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path}
		self.LastErr = e.edit(self, self.Split(toNode), editInsert)
	}
	return self
}

// InsertFrom Copy given node into current node.  If there are any existing containers of list
// items then this will fail by design.
func (self Selection) InsertFrom(fromNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path}
		self.LastErr = e.edit(self.Split(fromNode), self, editInsert)
	}
	return self
}

// UpsertInto Merge current node into given node.  If there are any existing containers of list
// items then data will be merged.
func (self Selection) UpsertInto(toNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path}
		self.LastErr = e.edit(self, self.Split(toNode), editUpsert)
	}
	return self
}

// Merge given node into current node.  If there are any existing containers of list
// items then data will be merged.
func (self Selection) UpsertFrom(fromNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path}
		self.LastErr = e.edit(self.Split(fromNode), self, editUpsert)
	}
	return self
}

// UpsertIntoSetDefaults is like UpsertInto but top container will have defaults set from YANG
func (self Selection) UpsertIntoSetDefaults(toNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path, useDefault: true}
		self.LastErr = e.edit(self, self.Split(toNode), editUpsert)
	}
	return self
}

// UpsertFromSetDefauls is like UpsertFrom but top container will have defaults set from YANG
func (self Selection) UpsertFromSetDefaults(fromNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path, useDefault: true}
		self.LastErr = e.edit(self.Split(fromNode), self, editUpsert)
	}
	return self
}

// Copy current node into given node.  There must be matching containers of list
// items or this will fail by design.
func (self Selection) UpdateInto(toNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path}
		self.LastErr = e.edit(self, self.Split(toNode), editUpdate)
	}
	return self
}

// Copy given node into current node.  There must be matching containers of list
// items or this will fail by design.
func (self Selection) UpdateFrom(fromNode Node) Selection {
	if self.LastErr == nil {
		e := editor{basePath: self.Path}
		self.LastErr = e.edit(self.Split(fromNode), self, editUpdate)
	}
	return self
}

// ClearField write nil/empty value to field.
func (self Selection) ClearField(m meta.Leafable) error {
	if self.LastErr != nil {
		return self.LastErr
	}
	r := FieldRequest{
		Request: Request{
			Selection: self,
		},
		Write: true,
		Clear: true,
		Meta:  m,
	}
	return self.SetValueHnd(&r, &ValueHandle{})
}

// Notifications let's caller subscribe to a node.  Node must be a 'notification' node.
func (self Selection) Notifications(stream NotifyStream) (NotifyCloser, error) {
	if self.LastErr != nil {
		return nil, self.LastErr
	}
	r := NotifyRequest{
		Request: Request{
			Selection: self,
		},
		Meta:   self.Meta().(*meta.Notification),
		Stream: checkStreamConstraints(self.Constraints, stream),
	}
	return self.Node.Notify(r)
}

func checkStreamConstraints(constraints *Constraints, orig NotifyStream) NotifyStream {
	if constraints == nil {
		return orig
	}
	return func(n Notification) {
		if keep, err := constraints.CheckNotifyFilterConstraints(n.Event); err != nil {
			n.Event = n.Event.Split(ErrorNode{Err: err})
			n.Event.LastErr = err
			//msg = Selection{LastErr: err} // msg.Split(ErrorNode{Err: err})
		} else if !keep {
			return
		}
		orig(n)
	}
}

// Action let's to call a procedure potentially passing on data and potentially recieving
// data back.
func (self Selection) Action(input Node) Selection {
	if self.LastErr != nil {
		return self
	}
	r := ActionRequest{
		Request: Request{
			Selection: self,
		},
		Meta: self.Meta().(*meta.Rpc),
	}

	if input != nil {
		r.Input = Selection{
			Browser:     self.Browser,
			Parent:      &self,
			Path:        &Path{parent: self.Path, meta: r.Meta.Input()},
			Node:        input,
			Constraints: self.Constraints,
			Context:     self.Context,
		}
	}

	if proceed, constraintErr := self.Constraints.CheckActionPreConstraints(&r); !proceed || constraintErr != nil {
		self.LastErr = constraintErr
		return self
	}

	rpcOutput, rerr := self.Node.Action(r)
	if rerr != nil {
		self.LastErr = rerr
		return self
	}

	var output Selection
	if rpcOutput != nil {
		output = Selection{
			Browser:     self.Browser,
			Parent:      &self,
			Path:        &Path{parent: self.Path, meta: r.Meta.Output()},
			Node:        rpcOutput,
			Constraints: self.Constraints,
			Context:     self.Context,
		}
	}

	if proceed, constraintErr := self.Constraints.CheckActionPostConstraints(r); !proceed || constraintErr != nil {
		self.LastErr = constraintErr
		return self
	}

	return output
}

// Set let's you set a leaf value on a container or list item.
func (self Selection) Set(ident string, value interface{}) error {
	if self.LastErr != nil {
		return self.LastErr
	}
	pos := meta.Find(self.Path.meta.(meta.HasDefinitions), ident)
	if pos == nil {
		return fmt.Errorf("%w. property not found %s", fc.NotFoundError, ident)
	}
	m := pos.(meta.Leafable)
	v, e := NewValue(m.Type(), value)
	if e != nil {
		return e
	}
	r := FieldRequest{
		Request: Request{
			Selection: self,
		},
		Write: true,
		Meta:  m,
	}
	return self.SetValueHnd(&r, &ValueHandle{Val: v})
}

func (self Selection) SetValueHnd(r *FieldRequest, hnd *ValueHandle) error {
	r.Write = true

	if proceed, constraintErr := self.Constraints.CheckFieldPreConstraints(r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	if err := self.Node.Field(*r, hnd); err != nil {
		return err
	}

	if proceed, constraintErr := self.Constraints.CheckFieldPostConstraints(*r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	return nil
}

// Get let's you get a leaf value from a container or list item
func (self Selection) Get(ident string) (interface{}, error) {
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
func (self Selection) GetValue(ident string) (val.Value, error) {

	if self.LastErr != nil {
		return nil, self.LastErr
	}
	pos := meta.Find(self.Path.meta.(meta.HasDefinitions), ident)
	if pos == nil {
		return nil, fmt.Errorf("%w. property not found %s", fc.NotFoundError, ident)
	}
	if !meta.IsLeaf(pos) {
		return nil, fmt.Errorf("%w. property is not a leaf %s", fc.NotFoundError, ident)
	}
	r := FieldRequest{
		Request: Request{
			Selection: self,
		},
		Meta: pos.(meta.Leafable),
	}

	r.Write = false
	var hnd ValueHandle
	err := self.GetValueHnd(&r, &hnd, true)
	return hnd.Val, err
}

func (self Selection) GetValueHnd(r *FieldRequest, hnd *ValueHandle, useDefault bool) error {
	if proceed, constraintErr := self.Constraints.CheckFieldPreConstraints(r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}
	if err := self.Node.Field(*r, hnd); err != nil {
		return err
	}
	if hnd.Val == nil && useDefault {
		if r.Meta.HasDefault() {
			var err error
			if hnd.Val, err = NewValue(r.Meta.Type(), r.Meta.Default()); err != nil {
				return err
			}
		}
	}

	if proceed, constraintErr := self.Constraints.CheckFieldPostConstraints(*r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	return nil
}
