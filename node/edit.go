package node

import (
	"fmt"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

type editStrategy int

var strategyNotImplemented = fmt.Errorf("strategy %w.", fc.NotImplementedError)

const (
	editUpsert editStrategy = iota + 1
	editInsert
	editUpdate
	editReplace
)

type editor struct {
	basePath   *Path
	useDefault bool
}

func (e editor) edit(from *Selection, to *Selection, s editStrategy) (err error) {
	if err := e.enter(from, to, false, s, true, true); err != nil {
		return err
	}
	to.Release()
	return nil
}

func (e editor) enter(from *Selection, to *Selection, new bool, strategy editStrategy, root bool, bubble bool) (err error) {
	if err = to.beginEdit(NodeRequest{New: new, Source: to, EditRoot: root}, bubble); err != nil {
		return
	}
	defer func() {
		if endErr := to.endEdit(NodeRequest{New: new, Source: to, EditRoot: root}, bubble); endErr != nil {
			err = fmt.Errorf("error during endEdit: %v, previous error: %w", endErr, err)
		}
	}()
	if meta.IsList(from.Meta()) && !from.InsideList {
		if err := e.list(from, to, from.Meta().(*meta.List), new, strategy); err != nil {
			return err
		}
	} else if meta.IsLeaf(from.Meta()) {
		if err := e.leaf(from, to, from.Meta().(meta.Leafable), new, strategy); err != nil {
			return err
		}
	} else {
		ml := newContainerMetaList(from)
		m := ml.nextMeta()
		for m != nil {
			var err error
			if meta.IsLeaf(m) {
				err = e.leaf(from, to, m.(meta.Leafable), new, strategy)
			} else {
				err = e.node(from, to, m.(meta.HasDataDefinitions), new, strategy)
			}
			if err != nil {
				return err
			}
			m = ml.nextMeta()
		}
	}
	return nil
}

func (e editor) leaf(from *Selection, to *Selection, m meta.Leafable, new bool, strategy editStrategy) error {
	r := FieldRequest{
		Request: Request{
			Selection: from,
			Path:      &Path{Parent: from.Path, Meta: m},
			Base:      e.basePath,
		},
		Meta: m,
	}
	useDefault := (strategy != editUpdate && new) || e.useDefault
	var hnd ValueHandle
	if err := from.get(&r, &hnd, useDefault); err != nil {
		return err
	}

	r.Selection = to
	r.From = from
	if hnd.Val != nil {
		// If there is a different choice selected, need to clear it
		// first if in upsert mode
		if strategy == editUpsert || strategy == editReplace {
			if err := e.clearOnDifferentChoiceCase(to, m); err != nil {
				return err
			}
		}
		if err := to.set(&r, &hnd); err != nil {
			return err
		}
	} else if strategy == editReplace {
		r.Clear = true
		if err := to.set(&r, &ValueHandle{}); err != nil {
			return err
		}
	}
	return nil
}

func (e editor) clearOnDifferentChoiceCase(existing *Selection, want meta.Meta) error {
	wantCase, valid := want.Parent().(*meta.ChoiceCase)
	if !valid {
		return nil
	}
	choice := wantCase.Parent().(*meta.Choice)
	existingCase, err := existing.Node.Choose(existing, choice)
	if err != nil {
		// we're eating the error here because destination may not implement choose because
		// it's a write-only implementation. clearing the old value is a courtesy anyway so
		// proceed with edit as planned.
		return nil
	}
	if existingCase == wantCase || existingCase == nil {
		return nil
	}
	return e.clearChoiceCase(existing, existingCase)
}

func (e editor) clearChoiceCase(sel *Selection, c *meta.ChoiceCase) error {
	i := newChoiceCaseIterator(sel, c)
	m := i.nextMeta()
	for m != nil {
		if meta.IsLeaf(m) {
			if err := sel.ClearField(m.(meta.Leafable)); err != nil {
				return err
			}
		} else {
			sub, err := sel.Find(m.(meta.Identifiable).Ident())
			if err != nil {
				return err
			}
			if sub != nil {
				if err := sub.Delete(); err != nil {
					return err
				}
			}
		}
		m = i.nextMeta()
	}
	return nil
}

func (e editor) node(from *Selection, to *Selection, m meta.HasDataDefinitions, new bool, strategy editStrategy) error {
	var newChild bool
	var err error
	// this ensures that even on panic we release any selections created in this func and it's loop
	var fromChild *Selection
	var toChild *Selection
	releaseFromChild := func() {
		if fromChild != nil {
			fromChild.Release()
			fromChild = nil
		}
	}
	defer releaseFromChild()
	releaseToChild := func() {
		if toChild != nil {
			toChild.Release()
			toChild = nil
		}
	}
	defer releaseToChild()

	fromRequest := ChildRequest{
		Request: Request{
			Selection: from,
			Path:      &Path{Parent: from.Path, Meta: m},
			Base:      e.basePath,
		},
		Meta: m,
	}
	fromChild, err = from.selekt(&fromRequest)
	if err != nil {
		return err
	}
	if fromChild == nil && strategy != editReplace {
		return nil
	}
	toRequest := ChildRequest{
		Request: Request{
			Selection: to,
			Path:      fromRequest.Path,
			Base:      e.basePath,
		},
		Meta: m,
	}
	toRequest.New = false
	toRequest.Selection = to

	toChild, err = to.selekt(&toRequest)
	if err != nil {
		return err
	}

	switch strategy {
	case editInsert:
		if toChild != nil {
			return fmt.Errorf("%w. item '%s' found in '%s'.  ", fc.ConflictError, m.Ident(), fromRequest.Path)
		}
		toRequest.New = true
		if toChild, err = to.selekt(&toRequest); err != nil {
			return err
		}
		newChild = true
	case editReplace:
		if toChild != nil {
			if err := toChild.delete(); err != nil {
				return err
			}
		}
		if fromChild == nil {
			return nil
		}
		toRequest.New = true
		if toChild, err = to.selekt(&toRequest); err != nil {
			return err
		}
		newChild = true
	case editUpsert:

		// If there is a different choice selected, need to clear it
		// first if in upsert mode
		if err := e.clearOnDifferentChoiceCase(to, m); err != nil {
			return err
		}

		if toChild == nil {
			toRequest.New = true
			if toChild, err = to.selekt(&toRequest); err != nil {
				return err
			}
			newChild = true
		}
	case editUpdate:
		if toChild == nil {
			return fmt.Errorf("%w. cannot update '%s' not found in '%s' container destination node ",
				fc.NotFoundError, m.Ident(), fromRequest.Path)
		}
	default:
		return strategyNotImplemented
	}

	if toChild == nil {
		return fmt.Errorf("'%s' could not create '%s' container node ", toRequest.Path, m.Ident())
	}
	if err := e.enter(fromChild, toChild, newChild, strategy, false, false); err != nil {
		return err
	}
	return nil
}

func (e editor) list(from *Selection, to *Selection, m *meta.List, new bool, strategy editStrategy) error {
	// this ensures that even on panic we release any selections created in this func and it's loop
	var fromChild *Selection
	var toChild *Selection
	releaseFromChild := func() {
		if fromChild != nil {
			fromChild.Release()
			fromChild = nil
		}
	}
	defer releaseFromChild()
	releaseToChild := func() {
		if toChild != nil {
			toChild.Release()
			toChild = nil
		}
	}
	defer releaseToChild()

	p := *from.Path
	fromRequest := &ListRequest{
		Request: Request{
			Selection: from,
			Path:      &p,
			Base:      e.basePath,
		},
		First: true,
		Meta:  m,
	}
	var key []val.Value
	var err error
	fromChild, key, err = from.selectVisibleListItem(fromRequest)
	if err != nil {
		return err
	}
	p.Key = key
	toRequest := ListRequest{
		Request: Request{
			Selection: to,
			Path:      &p,
			Base:      e.basePath,
		},
		First: true,
		Meta:  m,
	}

	for fromChild != nil {
		var newItem bool
		toChild = nil

		toRequest.First = true
		toRequest.SetRow(fromRequest.Row64)
		toRequest.Selection = to
		toRequest.From = fromChild
		toRequest.Key = key
		p.Key = key
		if len(key) > 0 {
			toRequest.New = false
			if toChild, _, _, err = to.selectListItem(&toRequest); err != nil {
				return err
			}
		}
		toRequest.New = true
		switch strategy {
		case editUpdate:
			if toChild == nil {
				return fmt.Errorf("%w, '%v' not found in '%s' list node ",
					fc.NotFoundError, key, to.Path)
			}
		case editUpsert:
			if toChild == nil {
				if toChild, _, _, err = to.selectListItem(&toRequest); err != nil {
					return err
				}
				newItem = true
			}
		case editInsert:
			if toChild != nil {
				return fmt.Errorf("%w, duplicate item found with same key in list %s",
					fc.ConflictError, to.Path)
			}
			if toChild, _, _, err = to.selectListItem(&toRequest); err != nil {
				return err
			}
			newItem = true
		case editReplace:
			if toChild, _, _, err = to.selectListItem(&toRequest); err != nil {
				return err
			}
			newItem = true
		default:
			return strategyNotImplemented
		}

		if toChild == nil {
			return fmt.Errorf("could not create destination list node %s", to.Path)
		}
		toChild.Path.Key = key
		if err = e.enter(fromChild, toChild, newItem, editUpsert, false, false); err != nil {
			return err
		}

		releaseToChild()
		releaseFromChild()

		fromRequest.IncrementRow()
		if fromChild, key, err = from.selectVisibleListItem(fromRequest); err != nil {
			return err
		}
	}
	if strategy == editReplace {
		// Iterate through "to" list, look if such element is in "from" list, delete otherwise
		p = *to.Path
		toRequest = ListRequest{
			Request: Request{
				Selection: to,
				Path:      &p,
				Base:      e.basePath,
			},
			First: true,
			Meta:  m,
		}
		fromRequest = &ListRequest{
			Request: Request{
				Selection: from,
				Path:      &p,
				Base:      e.basePath,
			},
			First: true,
			Meta:  m,
		}
		for {
			if toChild, _, key, err = to.selectListItem(&toRequest); err != nil {
				return err
			}
			if toChild == nil {
				break
			}

			fromRequest.First = true
			fromRequest.SetRow(toRequest.Row64)
			fromRequest.Selection = to
			fromRequest.From = toChild
			fromRequest.Key = key
			p.Key = key
			if fromChild, _, _, err = from.selectListItem(fromRequest); err != nil {
				return err
			}
			if fromChild == nil {
				fmt.Printf("|||||delete fromChild=%v\n", toChild.Path)
				if err = toChild.delete(); err != nil {
					return err
				}
				toChild = nil
			}

			releaseToChild()
			releaseFromChild()

			toRequest.IncrementRow()
		}
	}
	return nil
}
