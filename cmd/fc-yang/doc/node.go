package doc

import (
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/val"
)

func api(doc *doc) node.Node {
	return &nodeutil.Extend{
		Base: defNode(doc.Module),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "def":
				return defsNode(doc.DataDefs), nil
			}
			return nil, nil
		},
	}
}

func defsNode(defs []*def) node.Node {
	return &nodeutil.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			key := r.Key
			var d *def
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

func defNode(d *def) node.Node {
	return &nodeutil.Extend{
		Base: nodeutil.ReflectChild(d),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "parent":
				if d.Parent != nil {
					hnd.Val = val.String(meta.SchemaPath(d.Parent.Meta))
				}
			case "type":
				if d.ScalarType != "" {
					hnd.Val = val.String(d.ScalarType)
				}
			case "title":
				hnd.Val = val.String(d.Meta.Ident())
			case "description":
				hnd.Val = sval(d.Meta.(meta.Describable).Description())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "field":
				if len(d.Fields) > 0 {
					return defsNode(d.Fields), nil
				}
			case "input":
				if d.Input != nil {
					return defsNode(d.Input.Expand()), nil
				}
			case "output":
				if d.Output != nil {
					return defsNode(d.Output.Expand()), nil
				}
			case "action":
				if len(d.Actions) > 0 {
					return defsNode(d.Actions), nil
				}
			case "event":
				if len(d.Events) > 0 {
					return defsNode(d.Events), nil
				}
			}
			return nil, nil
		},
	}
}

func sval(s string) val.Value {
	if s != "" {
		return val.String(s)
	}
	return nil
}
