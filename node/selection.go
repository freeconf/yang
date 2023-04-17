package node

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"context"

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
//	 Example:
//	    var err error
//	    root := browser.Root()
//	    jay := root.Find("birds=bluejay")
//	    myCheckErr(dim.LastErr)
//
//	    // write
//	    err = jay.UpsertFrom(nodeutil.ReadJSON(`{"dimensions":55}`)).LastErr
//
//	    // read
//	    err = jay.UpsertInto(someOtherNode)
//
//	    // action
//	    _, err = jay.Find("fly").Action(nil)
//
//	    // subscribe
//	    reportRareBird := func(msg node.Selection) {
//		        fmt.Println(nodeutil.ReadJSON(msg))
//	    }
//	    unsubscribe, err := root.Find("rareSighting").Notify(reportRareBird)
//	    // unsubscribe
//	    unsubscribe()
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

var ErrNilSelection = errors.New("selection is nil")

func (sel Selection) Meta() meta.Definition {
	return sel.Path.Meta
}

// This selection points nowhere and must have been returned from a function that didn't find
// another selection
func (sel Selection) IsNil() bool {
	return sel.Path == nil
}

// Create a new independant selection with a different browser from this point in the tree based on a whole
// new data node
func (sel Selection) Split(node Node) Selection {
	if sel.IsNil() {
		return Selection{
			LastErr: errors.New("selection is nil"),
			Context: sel.Context,
		}
	}
	fork := sel
	fork.Parent = nil
	fork.Browser = NewBrowser(meta.RootModule(sel.Path.Meta), node)
	fork.Constraints = &Constraints{}
	fork.Node = node
	return fork
}

// If this is a selection in a list, this is the key value of that list item.
func (sel Selection) Key() []val.Value {
	return sel.Path.Key
}

func (sel Selection) selekt(r *ChildRequest) Selection {
	// check pre-constraints
	if proceed, constraintErr := sel.Constraints.CheckContainerPreConstraints(r); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: sel.Context,
		}
	}

	// select node
	var child Selection
	childNode, err := sel.Node.Child(*r)
	if err != nil {
		child = Selection{
			LastErr: err,
			Context: sel.Context,
		}
	} else if childNode == nil {
		child = Selection{}
	} else {
		child = Selection{
			Browser:     sel.Browser,
			Parent:      &sel,
			Path:        &Path{Parent: sel.Path, Meta: r.Meta},
			Node:        childNode,
			Constraints: sel.Constraints,
			Context:     sel.Context,
		}
		child.Context = childNode.Context(child)
		child.Context = sel.Constraints.ContextConstraint(child)
	}

	// check post-constraints
	if proceed, constraintErr := sel.Constraints.CheckContainerPostConstraints(*r, child); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: sel.Context,
		}
	}

	return child
}

type ListItem struct {
	Selection Selection
	Key       []val.Value
	req       ListRequest
}

// If at list, this will be iterator into first item in list
func (sel Selection) First() ListItem {
	item := ListItem{
		req: ListRequest{
			Request: Request{
				Selection: sel,
				Path:      sel.Path,
				Base:      sel.Path,
			},
			First: true,
			Meta:  sel.Meta().(*meta.List),
		},
	}
	return item.Next()
}

// iterating a list, get next item in list
func (li ListItem) Next() ListItem {
	li.Selection, li.Key = li.req.Selection.selectVisibleListItem(&li.req)
	li.req.IncrementRow()
	return li
}

func (sel Selection) selectVisibleListItem(r *ListRequest) (Selection, []val.Value) {
	for {
		sel, visible, key := sel.selectListItem(r)
		if visible || sel.IsNil() {
			return sel, key
		}
		r.IncrementRow()
	}
}

func (sel Selection) selectListItem(r *ListRequest) (Selection, bool, []val.Value) {
	// check pre-constraints
	if proceed, constraintErr := sel.Constraints.CheckListPreConstraints(r); !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: sel.Context,
		}, true, nil
	}

	// select node
	var child Selection
	childNode, key, err := sel.Node.Next(*r)
	if err != nil {
		child = Selection{
			LastErr: err,
			Context: sel.Context,
		}
	} else if childNode == nil {
		child = Selection{}
	} else {
		var parentPath *Path
		if sel.Parent != nil {
			parentPath = sel.Parent.Path
		}
		child = Selection{
			Browser: sel.Browser,
			Parent:  &sel,
			Node:    childNode,
			// NOTE: Path.parent is lists parentPath, not self.path
			Path:        &Path{Parent: parentPath, Meta: sel.Path.Meta, Key: key},
			InsideList:  true,
			Constraints: sel.Constraints,
			Context:     sel.Context,
		}
		child.Context = childNode.Context(child)
		child.Context = sel.Constraints.ContextConstraint(child)
	}

	// check post-constraints
	proceed, visible, constraintErr := sel.Constraints.CheckListPostConstraints(*r, child, r.Selection.Path.Key)
	if !proceed || constraintErr != nil {
		return Selection{
			LastErr: constraintErr,
			Context: sel.Context,
		}, visible, nil
	}
	return child, visible, key
}

func (sel Selection) Peek(consumer interface{}) interface{} {
	if sel.LastErr != nil {
		panic(sel.LastErr)
	}
	if sel.IsNil() {
		return nil
	}
	return sel.Node.Peek(sel, consumer)
}

// Apply constraints in the form of url parameters.
// Original selector and constraints will remain unaltered
// Example:
//
//	   sel2 = sel.Constrain("content=config&depth=4")
//	sel will not have content or depth constraints applies, but sel 2 will
func (sel Selection) Constrain(params string) Selection {
	if sel.LastErr != nil {
		return sel
	}
	if dummy, err := url.Parse("bogus?" + params); err != nil {
		sel.LastErr = err
		return sel
	} else {
		buildConstraints(&sel, dummy.Query())
		sel.Context = sel.Constraints.ContextConstraint(sel)
	}
	return sel
}

var errMaxDepthZeroNotAllowed = errors.New("depth zero is not allowed")

func buildConstraints(sel *Selection, params map[string][]string) {
	constraints := NewConstraints(sel.Constraints)
	maxDepth := MaxDepth{MaxDepth: 64}
	if n, found := findIntParam(params, "depth"); found {
		if n == 0 {
			sel.LastErr = errMaxDepthZeroNotAllowed
		} else {
			maxDepth.MaxDepth = n
		}
	}
	constraints.AddConstraint("depth", 10, 50, maxDepth)
	if p, found := params["fc.range"]; found {
		if listSelector, selectorErr := NewListRange(p[0]); selectorErr != nil {
			sel.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("fc.range", 20, 50, listSelector)
		}
	}
	if p, found := params["fields"]; found {
		if listSelector, selectorErr := NewFieldsMatcher(p[0]); selectorErr != nil {
			sel.LastErr = selectorErr
			return
		} else {
			constraints.AddConstraint("fields", 10, 50, listSelector)
		}
	}
	if p, found := params["fc.xfields"]; found {
		if listSelector, selectorErr := NewExcludeFieldsMatcher(p[0]); selectorErr != nil {
			sel.LastErr = selectorErr
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
		if c, err := NewContentConstraint(sel.Path, p[0]); err != nil {
			sel.LastErr = err
		} else {
			constraints.AddConstraint("content", 10, 70, c)
		}
	}

	if p, found := params["with-defaults"]; found {
		if c, err := NewWithDefaultsConstraint(p[0]); err != nil {
			sel.LastErr = err
		} else {
			constraints.AddConstraint("with-defaults", 50, 70, c)
		}
	}
	if p, found := params["filter"]; found {
		if c, err := NewFilterConstraint(p[0]); err != nil {
			sel.LastErr = err
		} else {
			constraints.AddConstraint("filter", 10, 50, c)
		}
	}
	if p, found := params["where"]; found {
		if c, err := NewWhere(p[0]); err != nil {
			sel.LastErr = err
		} else {
			constraints.AddConstraint("where", 10, 50, c)
		}
	}

	sel.Constraints = constraints
}

func (sel Selection) beginEdit(r NodeRequest, bubble bool) error {
	r.Selection = sel
	if sel.IsNil() {
		return errors.New("selection is nil")
	}
	if err := sel.Browser.Triggers.beginEdit(r); err != nil {
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

func (sel Selection) endEdit(r NodeRequest, bubble bool) error {
	r.Selection = sel
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
	if err := sel.Browser.Triggers.endEdit(r); err != nil {
		return err
	}
	return nil
}

func (sel Selection) Delete() (err error) {

	// allow children to recieve indication their parent is being deleted by
	// sending node request w/delete=true
	if err := sel.beginEdit(NodeRequest{Source: sel, Delete: true, EditRoot: true}, true); err != nil {
		return err
	}

	if sel.InsideList {
		r := ListRequest{
			Request: Request{
				Selection: *sel.Parent,
			},
			Meta:   sel.Meta().(*meta.List),
			Delete: true,
			Key:    sel.Key(),
		}
		if _, _, err := r.Selection.Node.Next(r); err != nil {
			return err
		}
	} else {
		r := ChildRequest{
			Request: Request{
				Selection: *sel.Parent,
			},
			Meta:   sel.Meta().(meta.HasDataDefinitions),
			Delete: true,
		}
		if _, err := r.Selection.Node.Child(r); err != nil {
			return err
		}
	}

	if err := sel.endEdit(NodeRequest{Source: sel, Delete: true, EditRoot: true}, true); err != nil {
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
func (sel Selection) InsertInto(toNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path}
		sel.LastErr = e.edit(sel, sel.Split(toNode), editInsert)
	}
	return sel
}

// InsertFrom Copy given node into current node.  If there are any existing containers of list
// items then this will fail by design.
func (sel Selection) InsertFrom(fromNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path}
		sel.LastErr = e.edit(sel.Split(fromNode), sel, editInsert)
	}
	return sel
}

// UpsertInto Merge current node into given node.  If there are any existing containers of list
// items then data will be merged.
func (sel Selection) UpsertInto(toNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path}
		sel.LastErr = e.edit(sel, sel.Split(toNode), editUpsert)
	}
	return sel
}

// Merge given node into current node.  If there are any existing containers of list
// items then data will be merged.
func (sel Selection) UpsertFrom(fromNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path}
		sel.LastErr = e.edit(sel.Split(fromNode), sel, editUpsert)
	}
	return sel
}

// UpsertIntoSetDefaults is like UpsertInto but top container will have defaults set from YANG
func (sel Selection) UpsertIntoSetDefaults(toNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path, useDefault: true}
		sel.LastErr = e.edit(sel, sel.Split(toNode), editUpsert)
	}
	return sel
}

// UpsertFromSetDefauls is like UpsertFrom but top container will have defaults set from YANG
func (sel Selection) UpsertFromSetDefaults(fromNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path, useDefault: true}
		sel.LastErr = e.edit(sel.Split(fromNode), sel, editUpsert)
	}
	return sel
}

// Copy current node into given node.  There must be matching containers of list
// items or this will fail by design.
func (sel Selection) UpdateInto(toNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path}
		sel.LastErr = e.edit(sel, sel.Split(toNode), editUpdate)
	}
	return sel
}

func (sel Selection) ReplaceFrom(fromNode Node) error {
	parent := sel.Parent
	if err := sel.Delete(); err != nil {
		return err
	}
	return parent.InsertFrom(fromNode).LastErr
}

// Copy given node into current node.  There must be matching containers of list
// items or this will fail by design.
func (sel Selection) UpdateFrom(fromNode Node) Selection {
	if sel.LastErr == nil {
		e := editor{basePath: sel.Path}
		sel.LastErr = e.edit(sel.Split(fromNode), sel, editUpdate)
	}
	return sel
}

// ClearField write nil/empty value to field.
func (sel Selection) ClearField(m meta.Leafable) error {
	if sel.LastErr != nil {
		return sel.LastErr
	}
	r := FieldRequest{
		Request: Request{
			Selection: sel,
		},
		Write: true,
		Clear: true,
		Meta:  m,
	}
	return sel.set(&r, &ValueHandle{})
}

// Notifications let's caller subscribe to a node.  Node must be a 'notification' node.
func (sel Selection) Notifications(stream NotifyStream) (NotifyCloser, error) {
	if sel.LastErr != nil {
		return nil, sel.LastErr
	}
	if sel.IsNil() {
		return nil, ErrNilSelection
	}
	r := NotifyRequest{
		Request: Request{
			Selection: sel,
		},
		Meta:   sel.Meta().(*meta.Notification),
		Stream: checkStreamConstraints(sel.Constraints, stream),
	}
	return sel.Node.Notify(r)
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
func (sel Selection) Action(input Node) Selection {
	if sel.LastErr != nil {
		return sel
	}
	r := ActionRequest{
		Request: Request{
			Selection: sel,
		},
		Meta: sel.Meta().(*meta.Rpc),
	}

	if input != nil {
		r.Input = Selection{
			Browser:     sel.Browser,
			Parent:      &sel,
			Path:        &Path{Parent: sel.Path, Meta: r.Meta.Input()},
			Node:        input,
			Constraints: sel.Constraints,
			Context:     sel.Context,
		}
	}

	if proceed, constraintErr := sel.Constraints.CheckActionPreConstraints(&r); !proceed || constraintErr != nil {
		sel.LastErr = constraintErr
		return sel
	}

	rpcOutput, rerr := sel.Node.Action(r)
	if rerr != nil {
		sel.LastErr = rerr
		return sel
	}

	var output Selection
	if rpcOutput != nil {
		output = Selection{
			Browser:     sel.Browser,
			Parent:      &sel,
			Path:        &Path{Parent: sel.Path, Meta: r.Meta.Output()},
			Node:        rpcOutput,
			Constraints: sel.Constraints,
			Context:     sel.Context,
		}
	}

	if proceed, constraintErr := sel.Constraints.CheckActionPostConstraints(r); !proceed || constraintErr != nil {
		sel.LastErr = constraintErr
		return sel
	}

	return output
}

// When you've selected a leaf field, this will coerse the data into correct value type
// then set. Error if coersion is not successful
func (sel Selection) SetValue(value interface{}) error {
	if sel.LastErr != nil {
		return sel.LastErr
	}
	if !meta.IsLeaf(sel.Path.Meta) {
		return fmt.Errorf("%s is not a leaf", sel.Path.Meta.Ident())
	}
	m := sel.Path.Meta.(meta.Leafable)
	v, e := NewValue(m.Type(), value)
	if e != nil {
		return e
	}
	return sel.Set(v)
}

// When you've selected a leaf field, this will set the value.
// Value must be in correct type according to YANG
func (sel Selection) Set(v val.Value) error {
	if !meta.IsLeaf(sel.Path.Meta) {
		return fmt.Errorf("%s is not a leaf", sel.Path.Meta.Ident())
	}
	m := sel.Path.Meta.(meta.Leafable)
	r := FieldRequest{
		Request: Request{
			Selection: sel,
		},
		Write: true,
		Meta:  m,
	}
	return sel.set(&r, &ValueHandle{Val: v})
}

func (sel Selection) set(r *FieldRequest, hnd *ValueHandle) error {
	r.Write = true

	if proceed, constraintErr := sel.Constraints.CheckFieldPreConstraints(r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	if err := sel.Node.Field(*r, hnd); err != nil {
		return err
	}

	if proceed, constraintErr := sel.Constraints.CheckFieldPostConstraints(*r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	return nil
}

// GetValue let's you get the leaf value as a Value instance. Returns null if value is null
func (sel Selection) Get() (val.Value, error) {
	if sel.LastErr != nil {
		return nil, sel.LastErr
	}
	if !meta.IsLeaf(sel.Path.Meta) {
		return nil, fmt.Errorf("%s is not a leaf", sel.Path.Meta.Ident())
	}
	m := sel.Path.Meta.(meta.Leafable)
	r := FieldRequest{
		Request: Request{
			Selection: sel,
		},
		Meta: m,
	}

	r.Write = false
	var hnd ValueHandle
	err := sel.get(&r, &hnd, true)
	return hnd.Val, err
}

func (sel Selection) get(r *FieldRequest, hnd *ValueHandle, useDefault bool) error {
	if proceed, constraintErr := sel.Constraints.CheckFieldPreConstraints(r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}
	if err := sel.Node.Field(*r, hnd); err != nil {
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

	if proceed, constraintErr := sel.Constraints.CheckFieldPostConstraints(*r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	return nil
}
