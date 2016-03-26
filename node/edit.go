package node

import (
	"fmt"
	"meta"
	"blit"
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
	if err = self.Fire(START_TREE_EDIT.New()); err == nil {
		if err = self.Fire(DELETE.New()); err == nil {
			err = self.Fire(END_TREE_EDIT.New())
		}
	}
	return
}

func (e *Editor) Edit(context *Context, strategy Strategy, controller WalkController) (err error) {
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
	if err = e.to.Fire(START_TREE_EDIT.New()); err == nil {
		if err = s.Walk(context, controller); err == nil {
			if err = e.to.Fire(LEAVE_EDIT.New()); err == nil {
				err = e.to.Fire(END_TREE_EDIT.New())
			}
		}
	}
	return
}

func (e *Editor) list(from *Selection, to *Selection, new bool, strategy Strategy) (Node, error) {
	s := &MyNode{Label: fmt.Sprint("Edit list ", from.node.String(), "=>", to.node.String())}
	// List Edit - See "List Edit State Machine" diagram for additional documentation
	s.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
		var created bool
		var fromNextNode Node
		fromRequest := r
		fromRequest.Selection = from
		fromRequest.New = false
		fromNextNode, key, err = from.node.Next(fromRequest)
		if err != nil || fromNextNode == nil {
			return
		}

		toRequest := r
		toRequest.First = true
		toRequest.Selection = to
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
				return nil, nil, blit.NewErrC(msg, 404)
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
				return nil, nil, blit.NewErrC(msg, 409)
			}
			if toNextNode, _, err = to.node.Next(toRequest); err != nil {
				return
			}
			created = true
		default:
			return nil, nil, blit.NewErrC("Stratgey not implmented", 501)
		}
		if err != nil {
			return
		} else  if toNextNode == nil {
			return nil, nil, blit.NewErr("Could not create destination list node " + to.String())
		}
		fromChild := from.SelectListItem(fromNextNode, key)
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

		var toChildNode Node
		toRequest := r
		toRequest.New = false
		toRequest.Selection = to
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
				return nil, blit.NewErrC(msg, 409)
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
				return nil, blit.NewErrC(msg, 404)
			}
		default:
			return nil, blit.NewErrC("Stratgey not implemented", 501)
		}

		if err != nil {
			return nil, err
		} else if toChildNode == nil {
			msg := fmt.Sprintf("'%s' could not create '%s' container node ", to.String(), r.Meta.GetIdent())
			return nil, blit.NewErr(msg)
		}
		// we always switch to upsert strategy because if there were any conflicts, it would have been
		// discovered in top-most level.
		fromChild := from.SelectChild(r.Meta, fromChildNode)
		toChild := to.SelectChild(r.Meta, toChildNode)
		if isList {
			return e.list(fromChild, toChild, created, UPSERT)
		}
		return e.container(fromChild, toChild, created, UPSERT)
	}
	s.OnEvent = func(sel *Selection, event Event) (err error) {
		return e.handleEvent(sel, from, to, new, event)
	}
	s.OnRead = func(r FieldRequest) (v *Value, err error) {
		if v, err = from.node.Read(r); err != nil {
			return
		}
		if v == nil && strategy != UPDATE {
			if r.Meta.GetDataType().HasDefault() {
				v = &Value{Type:r.Meta.GetDataType()}
				v.CoerseStrValue(r.Meta.GetDataType().Default())
			}
		}
		if v != nil {
			v.Type = r.Meta.GetDataType()
			if err = to.node.Write(r, v); err != nil {
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
			if err = to.Fire(NEW.New()); err != nil {
				return
			}
		}
		if err = to.Fire(LEAVE_EDIT.New()); err != nil {
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

