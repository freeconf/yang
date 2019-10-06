package doc

import (
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/val"
)

func Api(doc *Doc) node.Node {
	return &nodes.Extend{
		Base: nodes.ReflectChild(doc),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "def":
				if doc.Defs != nil {
					return defsNode(doc.Defs), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
	}
}

func defsNode(defs []*DocDef) node.Node {
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			key := r.Key
			var d *DocDef
			if key != nil {
				for _, candidate := range defs {
					if candidate.Meta.Ident() == key[0].String() {
						d = candidate
						break
					}
				}
			} else if r.Row < len(defs) {
				d = defs[r.Row]
				var err error
				key, err = node.NewValues(r.Meta.KeyMeta(), d.Meta.Ident())
				if err != nil {
					return nil, nil, err
				}
			}
			if d != nil {
				return defNode(d), key, nil
			}
			return nil, nil, nil
		},
	}
}

func metaNode(m meta.Definition) node.Node {
	return &nodes.Basic{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "title":
				hnd.Val = val.String(m.Ident())
			case "description":
				hnd.Val = sval(m.(meta.Describable).Description())
			}
			return nil
		},
	}
}

func defNode(def *DocDef) node.Node {
	return &nodes.Extend{
		Base: metaNode(def.Meta),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "parent":
				if def.Parent != nil {
					hnd.Val = val.String(meta.GetPath(def.Parent.Meta))
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "field":
				if len(def.Fields) > 0 {
					return fieldsNode(def.Fields), nil
				}
			case "action":
				if len(def.Actions) > 0 {
					return actionsNode(def.Actions), nil
				}
			case "event":
				if len(def.Events) > 0 {
					return eventsNode(def.Events), nil
				}
			}
			return nil, nil
		},
	}
}

func actionsNode(actions []*DocAction) node.Node {
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			key := r.Key
			var a *DocAction
			if key != nil {
				for _, candidate := range actions {
					if candidate.Meta.Ident() == key[0].String() {
						a = candidate
						break
					}
				}
			} else if r.Row < len(actions) {
				a = actions[r.Row]
				var err error
				key, err = node.NewValues(r.Meta.KeyMeta(), a.Meta.Ident())
				if err != nil {
					return nil, nil, err
				}
			}
			if a != nil {
				return actionNode(a), key, nil
			}
			return nil, nil, nil
		},
	}
}

func actionNode(a *DocAction) node.Node {
	return &nodes.Extend{
		Base: defNode(a.Def),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "input":
				if len(a.InputFields) > 0 {
					return fieldsNode(a.InputFields), nil
				}
			case "output":
				if len(a.OutputFields) > 0 {
					return fieldsNode(a.OutputFields), nil
				}
			}
			return nil, nil
		},
	}
}

func eventsNode(events []*DocEvent) node.Node {
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			key := r.Key
			var a *DocEvent
			if key != nil {
				for _, candidate := range events {
					if candidate.Meta.Ident() == key[0].String() {
						a = candidate
						break
					}
				}
			} else if r.Row < len(events) {
				a = events[r.Row]
				var err error
				key, err = node.NewValues(r.Meta.KeyMeta(), a.Meta.Ident())
				if err != nil {
					return nil, nil, err
				}
			}
			if a != nil {
				return eventNode(a), key, nil
			}
			return nil, nil, nil
		},
	}
}

func eventNode(a *DocEvent) node.Node {
	return defNode(a.Def)
}

func fieldsNode(defs []*DocField) node.Node {
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			key := r.Key
			var f *DocField
			if key != nil {
				for _, candidate := range defs {
					if candidate.Meta.Ident() == key[0].String() {
						f = candidate
						break
					}
				}
			} else if r.Row < len(defs) {
				f = defs[r.Row]
				var err error
				key, err = node.NewValues(r.Meta.KeyMeta(), f.Meta.Ident())
				if err != nil {
					return nil, nil, err
				}
			}
			if f != nil {
				return fieldNode(f), key, nil
			}
			return nil, nil, nil
		},
	}
}

func sval(s string) val.Value {
	if s != "" {
		return val.String(s)
	}
	return nil
}

func fieldNode(f *DocField) node.Node {
	n := nodes.ReflectChild(f)
	return &nodes.Extend{
		Base: metaNode(f.Meta),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "type", "details", "level":
				return n.Field(r, hnd)
			default:
				return p.Field(r, hnd)
			}
		},
	}
}
