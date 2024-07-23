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
	parent *Selection

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
}

var ErrNilSelection = errors.New("selection is nil")

func (sel *Selection) Meta() meta.Definition {
	return sel.Path.Meta
}

// Create a new independant selection with a different browser from this point in the tree based on a whole
// new data node
func (sel *Selection) Split(node Node) *Selection {
	if sel == nil {
		panic("selection is nil")
	}
	fork := *sel
	fork.parent = nil
	fork.Browser = NewBrowser(meta.RootModule(sel.Path.Meta), node)
	fork.Constraints = &Constraints{}
	fork.Node = node
	return &fork
}

func (sel *Selection) Parent() *Selection {
	return sel.parent
}

// If this is a selection in a list, this is the key value of that list item.
func (sel *Selection) Key() []val.Value {
	return sel.Path.Key
}

func (sel *Selection) makeCopy() (*Selection, error) {
	copy := *sel
	if sel.parent != nil {
		var err error
		if copy.parent, err = sel.parent.makeCopy(); err != nil {
			return nil, err
		}
	}
	copy.Context = copy.Node.Context(&copy)
	return &copy, nil
}

func (sel *Selection) selekt(r *ChildRequest) (*Selection, error) {
	// check pre-constraints
	if proceed, constraintErr := sel.Constraints.CheckContainerPreConstraints(r); !proceed || constraintErr != nil {
		return nil, constraintErr
	}

	// select node
	var child *Selection
	childNode, err := sel.Node.Child(*r)
	if err != nil || childNode == nil {
		return nil, err
	}
	child = &Selection{
		Browser:     sel.Browser,
		parent:      sel,
		Path:        &Path{Parent: sel.Path, Meta: r.Meta},
		Node:        childNode,
		Constraints: sel.Constraints,
		Context:     sel.Context,
	}
	child.Context = childNode.Context(child)
	child.Context = sel.Constraints.ContextConstraint(child)

	// check post-constraints
	if proceed, constraintErr := sel.Constraints.CheckContainerPostConstraints(*r, child); !proceed || constraintErr != nil {
		return nil, constraintErr
	}

	return child, nil
}

type ListItem struct {
	Selection *Selection
	Key       []val.Value
	req       ListRequest
}

// If at list, this will be iterator into first item in list
func (sel *Selection) First() (ListItem, error) {
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
func (li ListItem) Next() (ListItem, error) {
	var err error
	li.Selection, li.Key, err = li.req.Selection.selectVisibleListItem(&li.req)
	li.req.IncrementRow()
	return li, err
}

func (sel *Selection) selectVisibleListItem(r *ListRequest) (*Selection, []val.Value, error) {
	for {
		sel, visible, key, err := sel.selectListItem(r)
		if visible || err != nil || sel == nil {
			return sel, key, err
		}
		r.IncrementRow()
	}
}

func (sel *Selection) selectListItem(r *ListRequest) (*Selection, bool, []val.Value, error) {
	// check pre-constraints
	if proceed, constraintErr := sel.Constraints.CheckListPreConstraints(r); !proceed || constraintErr != nil {
		return nil, true, nil, constraintErr
	}

	// select node
	childNode, key, err := sel.Node.Next(*r)
	if err != nil || childNode == nil {
		return nil, true, nil, err
	}
	// no need to trust implementation to return the key we passed to them
	if key == nil {
		key = r.Key
	}
	var parentPath *Path
	if sel.parent != nil {
		parentPath = sel.parent.Path
	}
	child := &Selection{
		Browser: sel.Browser,
		parent:  sel,
		Node:    childNode,
		// NOTE: Path.parent is lists parentPath, not self.path
		Path:        &Path{Parent: parentPath, Meta: sel.Path.Meta, Key: key},
		InsideList:  true,
		Constraints: sel.Constraints,
		Context:     sel.Context,
	}
	child.Context = childNode.Context(child)
	child.Context = sel.Constraints.ContextConstraint(child)

	// check post-constraints
	proceed, visible, constraintErr := sel.Constraints.CheckListPostConstraints(*r, child, r.Selection.Path.Key)
	if !proceed || constraintErr != nil {
		return nil, visible, key, constraintErr
	}
	return child, visible, key, nil
}

func (sel *Selection) Release() {
	if sel.Node != nil {
		sel.Node.Release(sel)
	}
	if sel.parent != nil {
		sel.parent.Release()
	}
}

func (sel *Selection) Peek(consumer interface{}) interface{} {
	return sel.Node.Peek(sel, consumer)
}

// Apply constraints in the form of url parameters.
// Original selector and constraints will remain unaltered
// Example:
//
//	   sel2 = sel.Constrain("content=config&depth=4")
//	sel will not have content or depth constraints applies, but sel 2 will
func (sel *Selection) Constrain(params string) (*Selection, error) {
	dummy, err := url.Parse("bogus?" + params)
	if err != nil {
		return nil, err
	}
	copy := *sel
	if err = BuildConstraints(&copy, dummy.Query()); err != nil {
		return nil, err
	}
	copy.Context = copy.Constraints.ContextConstraint(sel)
	return &copy, nil
}

var errMaxDepthZeroNotAllowed = errors.New("depth zero is not allowed")

func BuildConstraints(sel *Selection, params map[string][]string) error {
	if len(params) == 0 {
		return nil
	}
	constraints := NewConstraints(sel.Constraints)
	maxDepth := MaxDepth{MaxDepth: 64}
	if n, found := findIntParam(params, "depth"); found {
		if n == 0 {
			return errMaxDepthZeroNotAllowed
		} else {
			maxDepth.MaxDepth = n
		}
	}
	constraints.AddConstraint("depth", 10, 50, maxDepth)
	if p, found := params["fc.range"]; found {
		if listSelector, selectorErr := NewListRange(p[0]); selectorErr != nil {
			return selectorErr
		} else {
			constraints.AddConstraint("fc.range", 20, 50, listSelector)
		}
	}
	if p, found := params["fields"]; found {
		if listSelector, selectorErr := NewFieldsMatcher(p[0]); selectorErr != nil {
			return selectorErr
		} else {
			constraints.AddConstraint("fields", 10, 50, listSelector)
		}
	}
	if p, found := params["fc.xfields"]; found {
		if listSelector, selectorErr := NewExcludeFieldsMatcher(p[0]); selectorErr != nil {
			return selectorErr
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
			return err
		} else {
			constraints.AddConstraint("content", 10, 70, c)
		}
	}

	if p, found := params["with-defaults"]; found {
		if c, err := NewWithDefaultsConstraint(p[0]); err != nil {
			return err
		} else {
			constraints.AddConstraint("with-defaults", 50, 70, c)
		}
	}
	if p, found := params["filter"]; found {
		if c, err := NewFilterConstraint(p[0]); err != nil {
			return err
		} else {
			constraints.AddConstraint("filter", 10, 50, c)
		}
	}
	if p, found := params["where"]; found {
		if c, err := NewWhere(p[0]); err != nil {
			return err
		} else {
			constraints.AddConstraint("where", 10, 50, c)
		}
	}

	sel.Constraints = constraints
	return nil
}

func (sel *Selection) beginEdit(r NodeRequest, bubble bool) error {
	r.Selection = sel
	if err := sel.Browser.Triggers.beginEdit(r); err != nil {
		return err
	}
	for {
		if err := r.Selection.Node.BeginEdit(r); err != nil {
			return err
		}
		if r.Selection.parent == nil || !bubble {
			break
		}
		r.Selection = r.Selection.parent
		r.EditRoot = false
	}
	return nil
}

func (sel *Selection) endEdit(r NodeRequest, bubble bool) error {
	r.Selection = sel
	for {
		if err := r.Selection.Node.EndEdit(r); err != nil {
			return err
		}
		if r.Selection.parent == nil || !bubble {
			break
		}
		r.Selection = r.Selection.parent
		r.EditRoot = false
	}
	if err := sel.Browser.Triggers.endEdit(r); err != nil {
		return err
	}
	return nil
}

func (sel *Selection) Delete() (err error) {

	// allow children to recieve indication their parent is being deleted by
	// sending node request w/delete=true
	if err := sel.beginEdit(NodeRequest{Source: sel, Delete: true, EditRoot: true}, true); err != nil {
		return err
	}
	defer func() {
		if endErr := sel.endEdit(NodeRequest{Source: sel, Delete: true, EditRoot: true}, true); endErr != nil {
			err = fmt.Errorf("error during endEdit: %v, previous error: %w", endErr, err)
		}
	}()

	err = sel.delete()
	return
}

func (sel *Selection) delete() (err error) {
	if sel.InsideList {
		r := ListRequest{
			Request: Request{
				Selection: sel.parent,
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
				Selection: sel.parent,
			},
			Meta:   sel.Meta().(meta.HasDataDefinitions),
			Delete: true,
		}
		if _, err := r.Selection.Node.Child(r); err != nil {
			return err
		}
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
func (sel *Selection) InsertInto(toNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel, sel.Split(toNode), editInsert)
}

// InsertFrom Copy given node into current node.  If there are any existing containers of list
// items then this will fail by design.
func (sel *Selection) InsertFrom(fromNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel.Split(fromNode), sel, editInsert)
}

// UpsertInto Merge current node into given node.  If there are any existing containers of list
// items then data will be merged.
func (sel *Selection) UpsertInto(toNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel, sel.Split(toNode), editUpsert)
}

// Merge given node into current node.  If there are any existing containers of list
// items then data will be merged.
func (sel *Selection) UpsertFrom(fromNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel.Split(fromNode), sel, editUpsert)
}

// UpsertIntoSetDefaults is like UpsertInto but top container will have defaults set from YANG
func (sel *Selection) UpsertIntoSetDefaults(toNode Node) error {
	e := editor{basePath: sel.Path, useDefault: true}
	return e.edit(sel, sel.Split(toNode), editUpsert)
}

// UpsertFromSetDefauls is like UpsertFrom but top container will have defaults set from YANG
func (sel *Selection) UpsertFromSetDefaults(fromNode Node) error {
	e := editor{basePath: sel.Path, useDefault: true}
	return e.edit(sel.Split(fromNode), sel, editUpsert)
}

// Copy current node into given node.  There must be matching containers of list
// items or this will fail by design.
func (sel *Selection) UpdateInto(toNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel, sel.Split(toNode), editUpdate)
}

// Replace current tree with given one
func (sel *Selection) ReplaceFrom(fromNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel.Split(fromNode), sel, editReplace)
}

// Copy given node into current node.  There must be matching containers of list
// items or this will fail by design.
func (sel *Selection) UpdateFrom(fromNode Node) error {
	e := editor{basePath: sel.Path}
	return e.edit(sel.Split(fromNode), sel, editUpdate)
}

// ClearField write nil/empty value to field.
func (sel *Selection) ClearField(m meta.Leafable) error {
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
func (sel *Selection) Notifications(stream NotifyStream) (NotifyCloser, error) {
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
			n.Event.Node = ErrorNode{err}
			orig(n)
			return
		} else if !keep {
			return
		}
		orig(n)
	}
}

// Action let's to call a procedure potentially passing on data and potentially recieving
// data back.
func (sel *Selection) Action(input Node) (*Selection, error) {
	r := ActionRequest{
		Request: Request{
			Selection: sel,
		},
		Meta: sel.Meta().(*meta.Rpc),
	}

	if input != nil {
		r.Input = &Selection{
			Browser:     sel.Browser,
			parent:      sel,
			Path:        &Path{Parent: sel.Path, Meta: r.Meta.Input()},
			Node:        input,
			Constraints: sel.Constraints,
			Context:     sel.Context,
		}
	}

	if proceed, constraintErr := sel.Constraints.CheckActionPreConstraints(&r); !proceed || constraintErr != nil {
		return nil, constraintErr
	}

	rpcOutput, err := sel.Node.Action(r)
	if err != nil {
		return nil, err
	}

	var output *Selection
	if rpcOutput != nil {
		output = &Selection{
			Browser:     sel.Browser,
			parent:      sel,
			Path:        &Path{Parent: sel.Path, Meta: r.Meta.Output()},
			Node:        rpcOutput,
			Constraints: sel.Constraints,
			Context:     sel.Context,
		}
	}

	if proceed, constraintErr := sel.Constraints.CheckActionPostConstraints(r); !proceed || constraintErr != nil {
		return nil, constraintErr
	}

	return output, nil
}

// When you've selected a leaf field, this will coerse the data into correct value type
// then set. Error if coersion is not successful
func (sel *Selection) SetValue(value interface{}) error {
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
func (sel *Selection) Set(v val.Value) error {
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

func (sel *Selection) set(r *FieldRequest, hnd *ValueHandle) error {
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

// Get let's you get the leaf value as a Value instance. Returns null if value is null
// Returns error if path is not found.
func (sel *Selection) Get() (val.Value, error) {
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

// GetValue let's you get the leaf value at the specified path or ident. Returns null if
// value is null.  Returns error if path is not found.
func (sel *Selection) GetValue(pathOrIdent string) (val.Value, error) {
	s, err := sel.Find(pathOrIdent)
	if err != nil {
		return nil, err
	}
	return s.Get()
}

func (sel *Selection) get(r *FieldRequest, hnd *ValueHandle, useDefault bool) error {
	if proceed, constraintErr := sel.Constraints.CheckFieldPreConstraints(r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}
	if err := sel.Node.Field(*r, hnd); err != nil {
		return err
	}
	if hnd.Val == nil && useDefault {
		if r.Meta.HasDefault() {
			var err error
			if hnd.Val, err = NewValue(r.Meta.Type(), r.Meta.DefaultValue()); err != nil {
				return err
			}
		}
	}

	if proceed, constraintErr := sel.Constraints.CheckFieldPostConstraints(*r, hnd); !proceed || constraintErr != nil {
		return constraintErr
	}

	return nil
}
