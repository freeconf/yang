package node

import (
	"fmt"
	"github.com/dhubler/c2g/meta"
	"github.com/dhubler/c2g/c2"
)

type Strategy int
const (
	UPSERT Strategy = iota + 1
	INSERT
	UPDATE
)

type Editor struct{
	from *Selection
	to *Selection
}

func (self *Selection) Delete() (err error) {
	if err = self.Fire(START_TREE_EDIT.New(self)); err == nil {
		if err = self.Fire(DELETE.New(self)); err != nil {
			return err
		}
		if (self.insideList) {
			if err = self.Parent().Fire(REMOVE_LIST_ITEM.New(self)); err != nil {
				return err
			}
		} else {
			if err = self.Parent().Fire(REMOVE_CONTAINER.New(self)); err != nil {
				return err
			}
		}
		err = self.Fire(END_TREE_EDIT.New(self))
	}
	return
}

func (e *Editor) Edit(strategy Strategy, controller WalkController) (err error) {
	var n Node
	if meta.IsList(e.from.path.meta) && !e.from.insideList {
		n, err = e.list(e.from, e.to, false, strategy)
	} else {
		n, err = e.container(e.from, e.to, false, strategy)
	}
	if err != nil {
		return err
	}
	// we could fork "from" or "to", shouldn't matter
	s := e.from.Fork(n)
	if err = e.to.Fire(START_TREE_EDIT.New(e.to)); err == nil {
		if err = s.Walk(controller); err == nil {
			if err = e.to.Fire(LEAVE_EDIT.New(e.to)); err == nil {
				err = e.to.Fire(END_TREE_EDIT.New(e.to))
			}
		}
	}
	return
}

func (e *Editor) list(from *Selection, to *Selection, new bool, strategy Strategy) (Node, error) {
	s := &MyNode{Label: fmt.Sprint("Edit list ", from.node.String(), "=>", to.node.String())}
	s.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
		var created bool
		var fromNextNode Node
		fromRequest := r
		fromRequest.Selection = from
		fromRequest.New = false
		fromRequest.From = to
		fromNextNode, key, err = from.node.Next(fromRequest)
		if err != nil || fromNextNode == nil {
			return
		}
		fromChild := from.SelectListItem(fromNextNode, key)

		toRequest := r
		toRequest.First = true
		toRequest.Selection = to
		toRequest.From = fromChild
		var toNextNode Node
		if len(key) > 0 {
			toRequest.Key = key
			toRequest.New = false
			if toNextNode, _, err = to.node.Next(toRequest); err != nil {
				return
			}
		}
		toRequest.New = true
		switch strategy {
		case UPDATE:
			if toNextNode == nil {
				msg := fmt.Sprintf("'%v' not found in '%s' list node ", key, r.Selection.String())
				return nil, nil, c2.NewErrC(msg, 404)
			}
		case UPSERT:
			if toNextNode == nil {
				if toNextNode, _, err = to.node.Next(toRequest); err != nil {
					return
				}
				created = true
			}
		case INSERT:
			if toNextNode != nil {
				msg := fmt.Sprint("Duplicate item found with same key in list ", r.Selection.String())
				return nil, nil, c2.NewErrC(msg, 409)
			}
			if toNextNode, _, err = to.node.Next(toRequest); err != nil {
				return
			}
			created = true
		default:
			return nil, nil, c2.NewErrC("Stratgey not implmented", 501)
		}
		if err != nil {
			return
		} else  if toNextNode == nil {
			return nil, nil, c2.NewErr("Could not create destination list node " + to.String())
		}
		toChild := to.SelectListItem(toNextNode, key)
		next, err = e.container(fromChild, toChild, created, UPSERT)
		return
	}
	s.OnEvent = func(sel *Selection, event Event) (err error) {
		return e.handleEvent(sel, from, to, new, event)
	}
	return s, nil
}

func (e *Editor) container(from *Selection, to *Selection, new bool, strategy Strategy) (Node, error) {
	s := &MyNode{Label: fmt.Sprint("Edit container ", from.node.String(), "=>", to.node.String())}
	s.OnChoose = func(sel *Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
		return from.node.Choose(from, choice)
	}
	s.OnSelect = func(r ContainerRequest) (Node, error) {
		var created bool
		var err error
		var fromChildNode Node
		fromRequest := r
		fromRequest.New = false
		fromRequest.Selection = from
		fromChildNode, err = from.node.Select(fromRequest)
		if err != nil || fromChildNode == nil {
			return nil, err
		}
		fromChild := from.SelectChild(r.Meta, fromChildNode)

		var toChildNode Node
		toRequest := r
		toRequest.New = false
		toRequest.Selection = to
		toRequest.From = fromChild
		toChildNode, err = to.node.Select(toRequest)
		if err != nil {
			return nil, err
		}
		isList := meta.IsList(r.Meta)
		toRequest.New = true

		switch strategy {
		case INSERT:
			if toChildNode != nil {
				msg := fmt.Sprintf("Duplicate item '%s' found in '%s' ", r.Meta.GetIdent(), r.Selection.String())
				return nil, c2.NewErrC(msg, 409)
			}
			if toChildNode, err = to.node.Select(toRequest); err != nil {
				return nil, err
			}
			created = true
		case UPSERT:
			if toChildNode == nil {
				if toChildNode, err = to.node.Select(toRequest); err != nil {
					return nil, err
				}
				created = true
			}
		case UPDATE:
			if toChildNode == nil {
				msg := fmt.Sprintf("cannot update '%s' not found in '%s' container destination node ",
					r.Meta.GetIdent(), r.Selection.String())
				return nil, c2.NewErrC(msg, 404)
			}
		default:
			return nil, c2.NewErrC("Stratgey not implemented", 501)
		}

		if err != nil {
			return nil, err
		} else if toChildNode == nil {
			msg := fmt.Sprintf("'%s' could not create '%s' container node ", to.String(), r.Meta.GetIdent())
			return nil, c2.NewErr(msg)
		}
		// we always switch to upsert strategy because if there were any conflicts, it would have been
		// discovered in top-most level.
		toChild := to.SelectChild(r.Meta, toChildNode)
		if isList {
			return e.list(fromChild, toChild, created, UPSERT)
		}
		return e.container(fromChild, toChild, created, UPSERT)
	}
	s.OnEvent = func(sel *Selection, event Event) (err error) {
		return e.handleEvent(sel, from, to, new, event)
	}
	s.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if err = from.node.Field(r, hnd); err != nil {
			return
		}
		if hnd.Val == nil && strategy != UPDATE {
			if r.Meta.GetDataType().HasDefault() {
				hnd.Val = &Value{Type:r.Meta.GetDataType()}
				hnd.Val.CoerseStrValue(r.Meta.GetDataType().Default())
			}
		}
		if hnd.Val != nil {
			hnd.Val.Type = r.Meta.GetDataType()
			r.Write = true
			if err = to.node.Field(r, hnd); err != nil {
				return
			}
		}
		return
	}

	return s, nil
}

func (e *Editor) handleEvent(sel *Selection, from *Selection, to *Selection, new bool, event Event) (err error) {
	if event.Type == LEAVE {
		if new {
			if err = to.Fire(NEW.New(to)); err != nil {
				return
			}
		}
		if err = to.Fire(LEAVE_EDIT.New(to)); err != nil {
			return
		}
	}

	if err = to.Fire(event); err != nil {
		return
	}
	if err = from.Fire(event); err != nil {
		return
	}
	return
}

func (e *Editor) loadKey(selection *Selection, explictKey []*Value) ([]*Value, error) {
	if len(explictKey) > 0 {
		return explictKey, nil
	}
	return selection.path.key, nil
}

