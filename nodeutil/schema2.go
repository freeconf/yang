package nodeutil

import (
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

type schema2 struct {
}

func Schema2(m *meta.Module) node.Node {
	api := schema2{}
	return api.manage(m)
}

func SchemaBrowser(fcYang *meta.Module, m *meta.Module) *node.Browser {
	return node.NewBrowser(fcYang, Schema2(m))
}

func (api schema2) manage(obj any) node.Node {
	return &Node{
		Object: obj,
		Options: NodeOptions{
			IgnoreEmpty:      true,
			TryPluralOnLists: true,
		},
		OnOptions: func(n *Node, m meta.Definition, opts NodeOptions) NodeOptions {
			switch m.Ident() {
			case "identity":
				opts.Ident = "identities"
			case "enumeration":
				opts.Ident = "enums"
			case "notify":
				opts.Ident = "notifications"
			case "dataDef":
				opts.Ident = "dataDefinitions"
			case "invert":
				opts.Ident = "inverted"
			case "id":
				if m.Parent().(meta.Definition).Ident() == "enumeration" {
					opts.Ident = "value"
				}
				opts.IgnoreEmpty = false
			case "position":
				opts.IgnoreEmpty = false
			case "defaults":
				opts.Ident = "default"
			}
			return opts
		},
		OnChoose: func(n *Node, sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
			switch choice.Ident() {
			case "body-stmt":
				return choice.Cases()[api.defType(n.Object.(meta.Meta))], nil
			}
			return n.DoChoose(sel, choice)
		},
		OnGetChild: func(n *Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "module", "container", "case", "leaf-list", "leaf", "anyxml", "anydata", "list", "choice", "pointer":
				return n, nil
			case "dataDef":
				if x, isChoice := n.Object.(*meta.Choice); isChoice {
					return n.NewList(x.Cases(), nil)
				}
				hasDefs := n.Object.(meta.HasDataDefinitions)
				if hasRecursiveChild(hasDefs) {
					defs := hasDefs.DataDefinitions()
					copy := make([]any, len(defs))
					for i := range defs {
						if defs[i].Parent() != hasDefs {
							copy[i] = &schemaPointer{parent: hasDefs, delegate: defs[i]}
						} else {
							copy[i] = defs[i]
						}
					}
					return n.NewList(copy, nil)
				}
			case "unique":
				if x, ok := n.Object.(*meta.List); ok {
					uniques := x.Unique()
					if len(uniques) > 0 {
						return api.uniques(uniques, 0), nil
					}
				}
			}
			return n.DoGetChild(r)
		},
		OnGetField: func(n *Node, r node.FieldRequest) (val.Value, error) {
			switch r.Meta.Ident() {
			case "status":
				// too lazy to fix fc-yang so eat here, these should not have status according to rfc
				_, isExtDefArg := n.Object.(*meta.ExtensionDefArg)
				_, isModule := n.Object.(*meta.Module)
				if isExtDefArg || isModule {
					return nil, nil
				}
			case "config":
				if x, valid := n.Object.(meta.HasConfig); valid {
					if !x.Config() {
						return val.Bool(x.Config()), nil
					}
				}
				return nil, nil
			case "unbounded":
				if x, valid := n.Object.(meta.HasListDetails); valid {
					if x.IsUnboundedSet() {
						return val.Bool(x.Unbounded()), nil
					}
					return nil, nil
				}
			case "derivedIds":
				identity := (n.Object).(*meta.Identity)
				ids := identity.DerivedDirectIds()
				if len(ids) == 0 {
					return nil, nil
				}
				return val.StringList(ids), nil
			case "key":
				l := n.Object.(*meta.List)
				keyMeta := l.KeyMeta()
				keys := make([]string, len(keyMeta))
				for i, k := range keyMeta {
					keys[i] = k.Ident()
				}
				return val.StringList(keys), nil
			case "when":
				if x, ok := n.Object.(meta.HasWhen); ok {
					if x.When() != nil {
						return sval(x.When().Expression()), nil
					}
				}
				return nil, nil
			case "label":
				if x, ok := n.Object.(*meta.Enum); ok {
					return val.String(x.Ident()), nil
				}
				if x, ok := n.Object.(*meta.Bit); ok {
					return val.String(x.Ident()), nil
				}
			case "length", "range":
				if x, ok := n.Object.(*meta.Range); ok {
					return val.String(x.String()), nil
				}
			case "base":
				if x, ok := n.Object.(*meta.Type); ok {
					if x.Base() != nil {
						ids := make([]string, len(x.Base()))
						for i, base := range x.Base() {
							ids[i] = base.Ident()
						}
						return val.StringList(ids), nil
					}
					return nil, nil
				}
			}
			return n.DoGetField(r)
		},
	}
}

type schemaPointer struct {
	parent   meta.HasDataDefinitions
	delegate meta.Definition
}

func (p *schemaPointer) Parent() meta.Meta {
	return p.parent
}

func (p *schemaPointer) Path() string {
	return meta.SchemaPathNoModule(p.delegate)
}

func (p *schemaPointer) Ident() string {
	return p.delegate.Ident()
}

func (p *schemaPointer) When() *meta.When {
	return p.delegate.(meta.HasWhen).When()
}

func (p *schemaPointer) Description() string {
	return p.delegate.(meta.Describable).Description()
}

func (p *schemaPointer) Reference() string {
	return p.delegate.(meta.Describable).Reference()
}

func (p *schemaPointer) Status() meta.Status {
	return p.delegate.(meta.HasStatus).Status()
}

func (p *schemaPointer) Extensions() []*meta.Extension {
	return p.delegate.Extensions()
}

func (p *schemaPointer) addExtension(x *meta.Extension) {

}

func hasRecursiveChild(p meta.HasDataDefinitions) bool {
	for _, d := range p.DataDefinitions() {
		if d.Parent() != p {
			return true
		}
	}
	return false
}

func (api schema2) recursivePath(sel *node.Selection, def meta.Meta) (string, bool) {
	//
	// Important: we're watching for recursive schema definitions to avoid infinite recursion
	// we can accomplish this by leveraging the fact that if you get to a child data def thru
	// a parent data def and the child has a different parent data def, then you are at a
	// recursive point.
	//
	psel := sel
	for psel != nil {
		if pmeta, valid := psel.Peek(nil).(meta.Meta); valid {
			if pmeta != def {
				if pmeta != def.Parent() {
					return meta.SchemaPathNoModule(def), true
				}
				return "", false
			}
		}
		psel = psel.Parent()
	}
	return "", false
}

func (api schema2) uniques(uniques [][]string, row int) node.Node {
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if r.Row < len(uniques) {
				return api.uniques(uniques, r.Row), nil, nil
			}
			return nil, nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			switch r.Meta.Ident() {
			case "leafs":
				hnd.Val = val.StringList(uniques[row])
			}
			return nil
		},
	}
}

func (api schema2) defType(data meta.Meta) string {
	switch data.(type) {
	case *meta.List:
		return "list"
	case *meta.Choice:
		return "choice"
	case *meta.Any:
		return "anyxml"
	case *meta.Notification:
		return "notification"
	case *meta.Rpc:
		return "action"
	case *meta.Leaf:
		return "leaf"
	case *meta.LeafList:
		return "leaf-list"
	case *meta.Container:
		return "container"
	case *meta.ChoiceCase:
		return "case"
	case *schemaPointer:
		return "pointer"
	}
	panic(fmt.Sprintf("unhandled type %T", data))
}
