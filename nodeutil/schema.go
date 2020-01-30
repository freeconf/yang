package nodeutil

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"

	"github.com/freeconf/yang/meta"
)

type schema struct {
	// resolve all uses, groups and typedefs.  if this is false, then depth must be
	// used to avoid infinite recursion
	resolve bool
}

/**
 * Schema is used to browse YANG models. If resolve is true all references like
 * groupings, uses typedefs are resolved, otherwise they are not.
 */
func Schema(yangModule *meta.Module, yourModule *meta.Module) *node.Browser {
	return node.NewBrowser(yangModule, schema{}.Yang(yourModule))
}

func (self schema) Yang(module *meta.Module) node.Node {
	s := &Basic{}
	s.OnChild = func(r node.ChildRequest) (node.Node, error) {
		switch r.Meta.Ident() {
		case "module":
			return self.module(module), nil
		}
		return nil, nil
	}
	return s
}

func sval(s string) val.Value {
	if s == "" {
		return nil
	}
	return val.String(s)
}

func (self schema) module(module *meta.Module) node.Node {
	return &Extend{
		Base: self.definition(module),
		OnChild: func(p node.Node, r node.ChildRequest) (child node.Node, err error) {
			switch r.Meta.Ident() {
			case "revision":
				if r := module.Revision(); r != nil {
					return self.rev(r), nil
				}
			case "identity":
				if len(module.Identities()) > 0 {
					return self.identities(module.Identities()), nil
				}
			case "feature":
				if len(module.Features()) > 0 {
					return self.features(module.Features()), nil
				}
			case "extensionDef":
				if len(module.ExtensionDefs()) > 0 {
					return self.extensionDefs(module.ExtensionDefs()), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "namespace":
				hnd.Val = sval(module.Namespace())
			case "prefix":
				hnd.Val = sval(module.Prefix())
			case "contact":
				hnd.Val = sval(module.Contact())
			case "organization":
				hnd.Val = sval(module.Organization())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) extensionDefs(defs map[string]*meta.ExtensionDef) node.Node {
	index := node.NewIndex(defs)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		Peekable: defs,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var x *meta.ExtensionDef
			key := r.Key
			if key != nil {
				x = defs[key[0].String()]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					ident := v.String()
					x = defs[ident]
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), ident); err != nil {
						return nil, nil, err
					}
				}
			}
			if x != nil {
				return self.extensionDef(x), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) extensions(e []*meta.Extension) node.Node {
	return &Basic{
		Peekable: e,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if r.Row >= len(e) {
				return nil, nil, nil
			}
			return self.extension(e[r.Row]), nil, nil
		},
	}
}

func (self schema) extension(e *meta.Extension) node.Node {
	return &Extend{
		Base: self.meta(e),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "keyword":
				if e.Keyword() != "" {
					hnd.Val = val.String(e.Keyword())
				}
			case "arguments":
				if len(e.Arguments()) > 0 {
					hnd.Val = val.StringList(e.Arguments())
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) extensionDef(def *meta.ExtensionDef) node.Node {
	return &Extend{
		Base: self.meta(def),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "argument":
				if len(def.Arguments()) > 0 {
					return self.extensionDefArgs(def.Arguments()), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
	}
}

func (self schema) extensionDefArgs(args []*meta.ExtensionDefArg) node.Node {
	return &Basic{
		Peekable: args,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if r.Row >= len(args) {
				return nil, nil, nil
			}
			return self.extensionDefArg(args[r.Row]), nil, nil
		},
	}
}

func (self schema) extensionDefArg(arg *meta.ExtensionDefArg) node.Node {
	return &Extend{
		Base: self.meta(arg),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "ident":
				hnd.Val = val.String(arg.Ident())
			case "yinElement":
				if arg.YinElement() {
					hnd.Val = val.Bool(arg.YinElement())
				}
				hnd.Val = val.Bool(arg.YinElement())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) features(idents map[string]*meta.Feature) node.Node {
	index := node.NewIndex(idents)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		Peekable: idents,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var x *meta.Feature
			key := r.Key
			if key != nil {
				x = idents[key[0].String()]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					ident := v.String()
					x = idents[ident]
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), ident); err != nil {
						return nil, nil, err
					}
				}
			}
			if x != nil {
				return self.meta(x), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) identities(idents map[string]*meta.Identity) node.Node {
	index := node.NewIndex(idents)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		Peekable: idents,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var x *meta.Identity
			key := r.Key
			if key != nil {
				x = idents[key[0].String()]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					ident := v.String()
					x = idents[ident]
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), ident); err != nil {
						return nil, nil, err
					}
				}
			}
			if x != nil {
				return self.identity(x), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) identity(i *meta.Identity) node.Node {
	return &Extend{
		Base: self.meta(i),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "ids":
				superset := i.Identities()
				ids := make([]string, len(superset))
				i := 0
				for id := range superset {
					ids[i] = id
					i++
				}
				sort.Strings(ids)
				hnd.Val = val.StringList(ids)
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) definition(data meta.Definition) node.Node {
	details, _ := data.(meta.HasDetails)
	listDetails, _ := data.(meta.HasListDetails)
	return &Extend{
		Base: self.meta(data),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "action":
				if x, ok := data.(meta.HasActions); ok {
					if len(x.Actions()) > 0 {
						return self.actions(x.Actions()), nil
					}
				}
			case "notify":
				if x, ok := data.(meta.HasNotifications); ok {
					if len(x.Notifications()) > 0 {
						return self.notifys(x.Notifications()), nil
					}
				}
			// TODO: Change this to just 'data'
			case "dataDef":
				if x, ok := data.(meta.HasDataDefinitions); ok {
					if len(x.DataDefinitions()) > 0 {
						return self.dataDefs(x), nil
					}
				}
			case "extension":
				if len(data.Extensions()) > 0 {
					return self.extensions(data.Extensions()), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "key":
				l := data.(*meta.List)
				keyMeta := l.KeyMeta()
				keys := make([]string, len(keyMeta))
				for i, k := range keyMeta {
					keys[i] = k.Ident()
				}
				hnd.Val = val.StringList(keys)
			case "config":
				if !details.Config() {
					hnd.Val = val.Bool(details.Config())
				}
			case "mandatory":
				if details.Mandatory() {
					hnd.Val = val.Bool(details.Mandatory())
				}
			case "maxElements":
				if !listDetails.Unbounded() {
					hnd.Val = val.Int32(listDetails.MaxElements())
				}
			case "minElements":
				if listDetails.MinElements() > 0 {
					hnd.Val = val.Int32(listDetails.MinElements())
				}
			case "unbounded":
				if listDetails.Unbounded() {
					hnd.Val = val.Bool(listDetails.Unbounded())
				}
			case "must":
				if x, ok := data.(meta.HasMusts); ok {
					if len(x.Musts()) > 0 {
						slist := make([]string, len(x.Musts()))
						for i, m := range x.Musts() {
							slist[i] = m.Expression()
						}
						hnd.Val = val.StringList(slist)
					}
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) rev(rev *meta.Revision) node.Node {
	return &Extend{
		Base: self.meta(rev),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			switch r.Meta.Ident() {
			case "rev-date":
				hnd.Val = sval(rev.Ident())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) dataDefs(m meta.HasDataDefinitions) node.Node {
	ddefs := m.DataDefinitions()
	return &Basic{
		Peekable: ddefs,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var d meta.Definition
			key := r.Key
			if key != nil {
				d, _ = m.Definition(key[0].String()).(meta.Definition)
			} else if r.Row < len(ddefs) {
				d = ddefs[r.Row]
				var err error
				if key, err = node.NewValues(r.Meta.KeyMeta(), d.Ident()); err != nil {
					return nil, nil, err
				}
			}
			if d != nil {
				return self.dataDef(d), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) actions(actions map[string]*meta.Rpc) node.Node {
	index := node.NewIndex(actions)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		Peekable: actions,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var x *meta.Rpc
			key := r.Key
			if key != nil {
				x = actions[key[0].String()]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					ident := v.String()
					x = actions[ident]
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), ident); err != nil {
						return nil, nil, err
					}
				}
			}
			if x != nil {
				return self.action(x), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) notifys(notifys map[string]*meta.Notification) node.Node {
	index := node.NewIndex(notifys)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		Peekable: notifys,
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var x *meta.Notification
			key := r.Key
			if key != nil {
				x = notifys[key[0].String()]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					ident := v.String()
					x = notifys[ident]
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), ident); err != nil {
						return nil, nil, err
					}
				}
			}
			if x != nil {
				return self.definition(x), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) dataType(dt *meta.Type) node.Node {
	return &Extend{
		Base: self.meta(dt),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "enumeration":
				if dt.Enum() != nil {
					return self.enumList(dt, dt.Enums()), nil
				}
			case "union":
				if len(dt.Union()) > 0 {
					return self.types(dt.Union()), nil
				}
			}
			return nil, nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			switch r.Meta.Ident() {
			case "ident":
				hnd.Val = sval(dt.Ident())
			case "length":
				if len(dt.Length()) > 0 {
					hnd.Val = val.StringList(rangesToStrings(dt.Length()))
				}
			case "range":
				if len(dt.Range()) > 0 {
					hnd.Val = val.StringList(rangesToStrings(dt.Range()))
				}
			case "path":
				hnd.Val = sval(dt.Path())
			case "format":
				hnd.Val, err = node.NewValue(r.Meta.Type(), int(dt.Format()))
			case "base":
				if dt.Base() != nil {
					hnd.Val = sval(dt.Base().Ident())
				}
			case "pattern":
				if len(dt.Patterns()) > 0 {
					hnd.Val = val.StringList(dt.Patterns())
				}
				return
			case "fractionDigits":
				if dt.Format().Single() == val.FmtDecimal64 {
					hnd.Val = val.Int32(dt.FractionDigits())
				}
			default:
				return p.Field(r, hnd)
			}
			return
		},
	}
}

func rangesToStrings(ranges []*meta.Range) []string {
	slist := make([]string, len(ranges))
	for i, r := range ranges {
		slist[i] = r.String()
	}
	return slist
}

func (self schema) types(u []*meta.Type) node.Node {
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if r.Row < len(u) {
				return self.dataType(u[r.Row]), nil, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) enumList(typeData *meta.Type, orig []*meta.Enum) node.Node {
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var key = r.Key
			var ref *meta.Enum
			if key != nil {
				for _, e := range orig {
					if e.Ident() == r.Key[0].String() {
						ref = e
						break
					}
				}
			} else if r.Row < len(orig) {
				ref = orig[r.Row]
				var err error
				if key, err = node.NewValues(r.Meta.KeyMeta(), ref.Ident()); err != nil {
					return nil, nil, err
				}
			}
			if ref != nil {
				return self.enum(ref), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) enum(e *meta.Enum) node.Node {
	return &Extend{
		Base: self.meta(e),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "label":
				hnd.Val = sval(e.Ident())
			case "id":
				hnd.Val = val.Int32(e.Value())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) action(rpc *meta.Rpc) node.Node {
	return &Extend{
		Base: self.definition(rpc),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "input":
				if rpc.Input() != nil {
					return self.definition(rpc.Input()), nil
				}
			case "output":
				if rpc.Output() != nil {
					return self.definition(rpc.Output()), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
	}
}

func (self schema) meta(m interface{}) node.Node {
	desc, _ := m.(meta.Describable)
	ident, _ := m.(meta.Identifiable)
	stat, hasStatus := m.(meta.HasStatus)
	return &Basic{
		Peekable: m,
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "ident":
				hnd.Val = sval(ident.Ident())
			case "description":
				hnd.Val = sval(desc.Description())
			case "reference":
				hnd.Val = sval(desc.Reference())
			case "status":
				var err error
				if hasStatus && stat.Status() != meta.Current {
					hnd.Val, err = node.NewValue(r.Meta.Type(), stat.Status())
				}
				return err
			case "when":
				if hw, ok := m.(meta.HasWhen); ok {
					if hw.When() != nil {
						hnd.Val = sval(hw.When().Expression())
					}
				}
			}
			return nil
		},
	}
}

func (self schema) list(l *meta.List) node.Node {
	return &Extend{
		Base: self.definition(l),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "key":
				keyMeta := l.KeyMeta()
				keys := make([]string, len(keyMeta))
				for i, k := range keyMeta {
					keys[i] = k.Ident()
				}
				hnd.Val = val.StringList(keys)
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) leafy(leafy meta.Leafable) node.Node {
	return &Extend{
		Base: self.definition(leafy),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "type":
				return self.dataType(leafy.Type()), nil
			}
			return p.Child(r)
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "units":
				hnd.Val = sval(leafy.Units())
			case "default":
				if leafy.HasDefault() {
					hnd.Val = val.Any{Thing: leafy.Default()}
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func (self schema) dataDef(data meta.Definition) node.Node {
	return &Extend{
		Base: self.meta(data),
		OnChoose: func(p node.Node, state node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
			return choice.Cases()[self.defType(data)], nil
		},
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "leaf-list", "leaf", "anyxml", "anydata":
				return self.leafy(data.(meta.Leafable)), nil
			case "list":
				return self.list(data.(*meta.List)), nil
			case "choice":
				return self.choice(data.(*meta.Choice)), nil
			case "container", "case":
				return self.definition(data), nil
			}
			return p.Child(r)
		},
	}
}

func (self schema) choice(c *meta.Choice) node.Node {
	index := node.NewIndex(c.Cases())
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "dataDef":
				if len(c.Cases()) > 0 {
					return self.choice(c), nil
				}
			}
			return nil, nil
		},
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var d meta.Definition
			key := r.Key
			if key != nil {
				d, _ = c.Cases()[key[0].String()]
			} else if r.Row < index.Len() {
				k := index.NextKey(r.Row)
				if k != node.NO_VALUE {
					ident := k.String()
					d = c.Cases()[ident]
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), ident); err != nil {
						return nil, nil, err
					}
				}
			}
			if d != nil {
				return self.dataDef(d), key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) defType(data meta.Meta) string {
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
	}
	panic(fmt.Sprintf("unhandled type %T", data))
}
