package node

import (
	"fmt"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

type editStrategy int

var strategyNotImplemented = fmt.Errorf("strategy %w.", fc.NotImplementedError)

const (
	editUpsert editStrategy = iota + 1
	editInsert
	editUpdate
)

type editor struct {
	basePath   *Path
	useDefault bool
}

func (self editor) edit(from Selection, to Selection, s editStrategy) (err error) {
	if err := self.enter(from, to, false, s, true, true); err != nil {
		return err
	}
	return nil
}

func (self editor) enter(from Selection, to Selection, new bool, strategy editStrategy, root bool, bubble bool) error {
	if err := to.beginEdit(NodeRequest{New: new, Source: to, EditRoot: root}, bubble); err != nil {
		return err
	}
	if meta.IsList(from.Meta()) && !from.InsideList {
		if err := self.list(from, to, from.Meta().(*meta.List), new, strategy); err != nil {
			return err
		}
	} else if meta.IsLeaf(from.Meta()) {
		if err := self.leaf(from, to, from.Meta().(meta.Leafable), new, strategy); err != nil {
			return err
		}
	} else {
		ml := newContainerMetaList(from)
		m := ml.nextMeta()
		//fmt.Printf("Begin %s\n", meta.SchemaPath(from.Meta()))
		for m != nil {
			var err error
			if meta.IsLeaf(m) {
				err = self.leaf(from, to, m.(meta.Leafable), new, strategy)
			} else {
				err = self.node(from, to, m.(meta.HasDataDefinitions), new, strategy)
			}
			if err != nil {
				return err
			}
			m = ml.nextMeta()
		}
		//fmt.Printf("Ended %s\n", meta.SchemaPath(from.Meta()))
	}
	if err := to.endEdit(NodeRequest{New: new, Source: to, EditRoot: root}, bubble); err != nil {
		return err
	}
	return nil
}

func (self editor) leaf(from Selection, to Selection, m meta.Leafable, new bool, strategy editStrategy) error {
	r := FieldRequest{
		Request: Request{
			Selection: from,
			Path:      &Path{Parent: from.Path, Meta: m},
			Base:      self.basePath,
		},
		Meta: m,
	}
	useDefault := (strategy != editUpdate && new) || self.useDefault
	var hnd ValueHandle
	if err := from.get(&r, &hnd, useDefault); err != nil {
		return err
	}

	if hnd.Val != nil {
		// If there is a different choice selected, need to clear it
		// first if in upsert mode
		if strategy == editUpsert {
			if err := self.clearOnDifferentChoiceCase(to, m); err != nil {
				return err
			}
		}

		r.Selection = to
		if err := to.set(&r, &hnd); err != nil {
			return err
		}
	}
	return nil
}

func (self editor) clearOnDifferentChoiceCase(existing Selection, want meta.Meta) error {
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
	return self.clearChoiceCase(existing, existingCase)
}

func (self editor) clearChoiceCase(sel Selection, c *meta.ChoiceCase) error {
	i := newChoiceCaseIterator(sel, c)
	m := i.nextMeta()
	for m != nil {
		if meta.IsLeaf(m) {
			if err := sel.ClearField(m.(meta.Leafable)); err != nil {
				return err
			}
		} else {
			sub := sel.Find(m.(meta.Identifiable).Ident())
			if !sub.IsNil() {
				if err := sub.Delete(); err != nil {
					return err
				}
			}
		}
		m = i.nextMeta()
	}
	return nil
}

func (self editor) node(from Selection, to Selection, m meta.HasDataDefinitions, new bool, strategy editStrategy) error {
	var newChild bool
	fromRequest := ChildRequest{
		Request: Request{
			Selection: from,
			Path:      &Path{Parent: from.Path, Meta: m},
			Base:      self.basePath,
		},
		Meta: m,
	}
	fromChild := from.selekt(&fromRequest)
	if fromChild.LastErr != nil || fromChild.IsNil() {
		return fromChild.LastErr
	}
	toRequest := ChildRequest{
		Request: Request{
			Selection: to,
			Path:      fromRequest.Path,
			Base:      self.basePath,
		},
		From: fromChild,
		Meta: m,
	}
	toRequest.New = false
	toRequest.Selection = to
	toRequest.From = fromChild

	toChild := to.selekt(&toRequest)
	if toChild.LastErr != nil {
		return toChild.LastErr
	}
	toRequest.New = true
	switch strategy {
	case editInsert:
		if !toChild.IsNil() {
			return fmt.Errorf("%w. item '%s' found in '%s'.  ", fc.ConflictError, m.Ident(), fromRequest.Path)
		}
		if toChild = to.selekt(&toRequest); toChild.LastErr != nil {
			return toChild.LastErr
		}
		newChild = true
	case editUpsert:

		// If there is a different choice selected, need to clear it
		// first if in upsert mode
		if err := self.clearOnDifferentChoiceCase(to, m); err != nil {
			return err
		}

		if toChild.IsNil() {
			if toChild = to.selekt(&toRequest); toChild.LastErr != nil {
				return toChild.LastErr
			}
			newChild = true
		}
	case editUpdate:
		if toChild.IsNil() {
			return fmt.Errorf("%w. cannot update '%s' not found in '%s' container destination node ",
				fc.NotFoundError, m.Ident(), fromRequest.Path)
		}
	default:
		return strategyNotImplemented
	}

	if toChild.IsNil() {
		return fmt.Errorf("'%s' could not create '%s' container node ", toRequest.Path, m.Ident())
	}
	if err := self.enter(fromChild, toChild, newChild, strategy, false, false); err != nil {
		return err
	}
	return nil
}

func (self editor) list(from Selection, to Selection, m *meta.List, new bool, strategy editStrategy) error {
	p := *from.Path
	fromRequest := &ListRequest{
		Request: Request{
			Selection: from,
			Path:      &p,
			Base:      self.basePath,
		},
		First: true,
		Meta:  m,
	}
	fromChild, key := from.selectVisibleListItem(fromRequest)
	if fromChild.LastErr != nil {
		return fromChild.LastErr
	} else if fromChild.IsNil() {
		return nil
	}
	p.Key = key
	toRequest := ListRequest{
		Request: Request{
			Selection: to,
			Path:      &p,
			Base:      self.basePath,
		},
		First: true,
		Meta:  m,
	}
	empty := Selection{}
	var toChild Selection
	for !fromChild.IsNil() {
		var newItem bool
		toChild = empty

		toRequest.First = true
		toRequest.SetRow(fromRequest.Row64)
		toRequest.Selection = to
		toRequest.From = fromChild
		toRequest.Key = key
		p.Key = key
		if len(key) > 0 {
			toRequest.New = false
			if toChild, _, _ = to.selectListItem(&toRequest); toChild.LastErr != nil {
				return toChild.LastErr
			}
		}
		toRequest.New = true
		switch strategy {
		case editUpdate:
			if toChild.IsNil() {
				return fmt.Errorf("%w, '%v' not found in '%s' list node ",
					fc.NotFoundError, key, to.Path)
			}
		case editUpsert:
			if toChild.IsNil() {
				toChild, _, _ = to.selectListItem(&toRequest)
				newItem = true
			}
		case editInsert:
			if !toChild.IsNil() {
				return fmt.Errorf("%w, duplicate item found with same key in list %s",
					fc.ConflictError, to.Path)
			}
			toChild, _, _ = to.selectListItem(&toRequest)
			newItem = true
		default:
			return strategyNotImplemented
		}

		if toChild.LastErr != nil {
			return toChild.LastErr
		} else if toChild.IsNil() {
			return fmt.Errorf("could not create destination list node %s", to.Path)
		}
		toChild.Path.Key = key
		if err := self.enter(fromChild, toChild, newItem, editUpsert, false, false); err != nil {
			return err
		}

		fromRequest.IncrementRow()
		if fromChild, key = from.selectVisibleListItem(fromRequest); fromChild.LastErr != nil {
			return fromChild.LastErr
		}
	}
	return nil
}
