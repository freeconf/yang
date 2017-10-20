package nodes

import (
	"fmt"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
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
func Schema(m *meta.Module, resolve bool) *node.Browser {
	return node.NewBrowser(yangModule(yang.YangPath()), schema{resolve: resolve}.Yang(m))
}

func SchemaWithYangPath(ypath meta.StreamSource, m *meta.Module, resolve bool) *node.Browser {
	return node.NewBrowser(yangModule(ypath), schema{resolve: resolve}.Yang(m))
}

func (self schema) Yang(module *meta.Module) node.Node {
	s := &Basic{}
	s.OnChild = func(r node.ChildRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "module":
			return self.Module(module), nil
		}
		return nil, nil
	}
	return s
}

func (self schema) Module(module *meta.Module) node.Node {
	return &Extend{
		Base: self.MetaList(module),
		OnChild: func(parent node.Node, r node.ChildRequest) (child node.Node, err error) {
			switch r.Meta.GetIdent() {
			case "revision":
				if r.New {
					module.Revision = &meta.Revision{}
				}
				if module.Revision != nil {
					return self.Revision(module.Revision), nil
				}
				return nil, nil
			}
			return parent.Child(r)
		},
	}
}

func (self schema) Revision(rev *meta.Revision) node.Node {
	return &Basic{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			switch r.Meta.GetIdent() {
			case "rev-date":
				if r.Write {
					rev.Ident = hnd.Val.String()
				} else {
					hnd.Val = val.String(rev.Ident)
				}
			default:
				if r.Write {
					err = node.WriteField(r.Meta, rev, hnd.Val)
				} else {
					hnd.Val, err = node.ReadField(r.Meta, rev)
				}
			}
			return nil
		},
	}
}

func (self schema) Type(typeData *meta.DataType) (node.Node, error) {
	info, err := typeData.Info()
	if err != nil {
		return nil, err
	}
	return &Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "enumeration":
				var l val.EnumList
				if self.resolve {
					l = info.Enum
				} else {
					l = typeData.EnumerationRef
				}
				if r.New || len(l) > 0 {
					return self.Enum(typeData, l), nil
				}
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			switch r.Meta.GetIdent() {
			case "ident":
				if r.Write {
					typeData.Ident = hnd.Val.String()
					typeData.SetFormat(val.TypeAsFormat(hnd.Val.String()))
				} else {
					hnd.Val = val.String(typeData.Ident)
				}
			case "minLength":
				if r.Write {
					typeData.SetMinLength(hnd.Val.Value().(int))
				} else {
					if self.resolve {
						hnd.Val = val.Int32(info.MinLength)
					} else {
						if typeData.MinLengthPtr != nil {
							hnd.Val = val.Int32(*typeData.MinLengthPtr)
						}
					}
				}
			case "maxLength":
				if r.Write {
					typeData.SetMaxLength(hnd.Val.Value().(int))
				} else {
					if self.resolve {
						hnd.Val = val.Int32(info.MaxLength)
					} else {
						if typeData.MaxLengthPtr != nil {
							hnd.Val = val.Int32(*typeData.MaxLengthPtr)
						}
					}
				}
			case "path":
				if r.Write {
					typeData.SetPath(hnd.Val.String())
				} else {
					if self.resolve {
						hnd.Val = val.String(info.Path)
					} else {
						if typeData.PathPtr != nil {
							hnd.Val = val.String(*typeData.PathPtr)
						}
					}
				}
			}
			return
		},
	}, nil
}

func (self schema) Enum(typeData *meta.DataType, orig val.EnumList) node.Node {
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var key = r.Key
			var ref val.Enum
			if r.New {
				ref.Label = r.Key[0].String()
			} else if key != nil {
				ref, _ = orig.ByLabel(r.Key[0].String())
			} else {
				if len(orig) < r.Row {
					ref = orig[r.Row]
					key = []val.Value{val.String(ref.Label)}
				}
			}
			if !ref.Empty() {
				n := &Extend{
					Base: ReflectChild(ref),
					OnEndEdit: func(node.Node, node.NodeRequest) error {
						typeData.EnumerationRef = append(typeData.EnumerationRef, ref)
						return nil
					},
				}
				return n, key, nil
			}
			return nil, nil, nil
		},
	}
}

func (self schema) Groupings(groupings meta.MetaList) node.Node {
	s := &Basic{}
	i := listIterator{dataList: groupings, resolve: self.resolve}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		var key = r.Key
		var group *meta.Grouping
		if r.New {
			group = &meta.Grouping{Ident: r.Key[0].String()}
			groupings.AddMeta(group)
		} else {
			if more, err := i.iterate(r.Selection, r.Meta, r.Key, r.First, r.Row); err != nil {
				return nil, nil, err
			} else if more {
				group = i.data.(*meta.Grouping)
				if len(key) == 0 {
					var err error
					if key, err = node.NewValues(r.Meta.KeyMeta(), group.Ident); err != nil {
						return nil, nil, err
					}
				}
			}
		}
		if group != nil {
			return self.MetaList(group), key, nil
		}
		return nil, nil, nil
	}
	return s
}

func (self schema) RpcIO(i *meta.RpcInput, o *meta.RpcOutput) node.Node {
	var io meta.MetaList
	if i != nil {
		io = i
	} else {
		io = o
	}
	return self.MetaList(io)
}

func (self schema) createGroupingsTypedefsDefinitions(parent meta.MetaList, childMeta meta.Meta) meta.Meta {
	var child meta.Meta
	switch childMeta.GetIdent() {
	case "leaf":
		child = &meta.Leaf{}
	case "anyxml":
		child = &meta.Any{}
	case "leaf-list":
		child = &meta.LeafList{}
	case "container":
		child = &meta.Container{}
	case "list":
		child = &meta.List{}
	case "uses":
		child = &meta.Uses{}
	case "grouping":
		child = &meta.Grouping{}
	case "typedef":
		child = &meta.Typedef{}
	case "rpc", "action":
		child = &meta.Rpc{}
	case "notification":
		child = &meta.Notification{}
	case "choice":
		child = &meta.Choice{}
	case "case":
		child = &meta.ChoiceCase{}
	default:
		panic("Unknown type:" + childMeta.GetIdent())
	}
	parent.AddMeta(child)
	return child
}

func (self schema) Rpc(rpc *meta.Rpc) node.Node {
	return &Extend{
		Base: ReflectChild(rpc),
		OnChild: func(parent node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "input":
				if r.New {
					rpc.AddMeta(&meta.RpcInput{})
				}
				if rpc.Input != nil {
					return self.RpcIO(rpc.Input, nil), nil
				}
				return nil, nil
			case "output":
				if r.New {
					rpc.AddMeta(&meta.RpcOutput{})
				}
				if rpc.Output != nil {
					return self.RpcIO(nil, rpc.Output), nil
				}
				return nil, nil
			}
			return parent.Child(r)
		},
	}
}

func (self schema) Typedefs(typedefs meta.MetaList) node.Node {
	s := &Basic{}
	i := listIterator{dataList: typedefs, resolve: self.resolve}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		var key = r.Key
		var typedef *meta.Typedef
		if r.New {
			typedef = &meta.Typedef{Ident: r.Key[0].String()}
			typedefs.AddMeta(typedef)
		} else {
			if more, err := i.iterate(r.Selection, r.Meta, r.Key, r.First, r.Row); err != nil {
				return nil, nil, err
			} else if more {
				typedef = i.data.(*meta.Typedef)
				if len(key) == 0 {
					if key, err = node.NewValues(r.Meta.KeyMeta(), typedef.Ident); err != nil {
						return nil, nil, err
					}
				}
			}
		}
		if typedef != nil {
			return self.Typedef(typedef), key, nil
		}
		return nil, nil, nil
	}
	return s
}

func (self schema) Typedef(typedef *meta.Typedef) node.Node {
	return &Extend{
		Base: ReflectChild(typedef),
		OnChild: func(parent node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "type":
				if r.New {
					typedef.SetDataType(&meta.DataType{Parent: typedef})
				}
				if typedef.DataType != nil {
					return self.Type(typedef.DataType)
				}
			}
			return nil, nil
		},
	}
}

func (self schema) MetaList(data meta.MetaList) node.Node {
	var details *meta.Details
	if hasDetails, ok := data.(meta.HasDetails); ok {
		details = hasDetails.Details()
	}
	return &Extend{
		Base: ReflectChild(data),
		OnChild: func(parent node.Node, r node.ChildRequest) (node.Node, error) {
			hasGroupings, implementsHasGroupings := data.(meta.HasGroupings)
			hasTypedefs, implementsHasTypedefs := data.(meta.HasTypedefs)
			switch r.Meta.GetIdent() {
			case "groupings":
				if !self.resolve && implementsHasGroupings {
					groupings := hasGroupings.GetGroupings()
					if r.New || !meta.HasChildren(groupings) {
						return self.Groupings(groupings), nil
					}
				}
				return nil, nil
			case "typedefs":
				if !self.resolve && implementsHasTypedefs {
					typedefs := hasTypedefs.GetTypedefs()
					if r.New || !meta.HasChildren(typedefs) {
						return self.Typedefs(typedefs), nil
					}
				}
				return nil, nil
			case "definitions":
				defs := data.(meta.MetaList)
				if r.New || !meta.HasChildren(defs) {
					return self.Definitions(defs), nil
				}
				return nil, nil
			}
			return parent.Child(r)
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			switch r.Meta.GetIdent() {
			case "config":
				if r.Write {
					details.SetConfig(hnd.Val.Value().(bool))
				} else {
					if self.resolve || details.ConfigPtr != nil {
						hnd.Val = val.Bool(details.Config(r.Selection.Path))
					}
				}
			case "mandatory":
				if r.Write {
					details.SetMandatory(hnd.Val.Value().(bool))
				} else {
					if self.resolve || details.MandatoryPtr != nil {
						hnd.Val = val.Bool(details.Mandatory())
					}
				}
			default:
				return p.Field(r, hnd)
			}
			return
		},
	}
}

func (self schema) Leaf(leaf *meta.Leaf, leafList *meta.LeafList, any *meta.Any) node.Node {
	var leafy meta.HasDataType
	if leaf != nil {
		leafy = leaf
	} else if leafList != nil {
		leafy = leafList
	} else {
		leafy = any
	}
	s := &Basic{
		Peekable: leafy,
	}
	details := leafy.(meta.HasDetails).Details()
	s.OnChild = func(r node.ChildRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "type":
			if r.New {
				leafy.SetDataType(&meta.DataType{Parent: leafy})
			}
			if leafy.GetDataType() != nil {
				return self.Type(leafy.GetDataType())
			}
		}
		return nil, nil
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		switch r.Meta.GetIdent() {
		case "config":
			if r.Write {
				details.SetConfig(hnd.Val.Value().(bool))
			} else {
				if self.resolve || details.ConfigPtr != nil {
					hnd.Val = val.Bool(details.Config(r.Selection.Path))
				}
			}
		case "mandatory":
			if r.Write {
				details.SetMandatory(hnd.Val.Value().(bool))
			} else {
				if self.resolve || details.MandatoryPtr != nil {
					hnd.Val = val.Bool(details.Mandatory())
				}
			}
		default:
			if r.Write {
				node.WriteField(r.Meta, leafy, hnd.Val)
			} else {
				hnd.Val, err = node.ReadField(r.Meta, leafy)
			}
		}
		return

	}
	return s
}

func (self schema) Uses(data *meta.Uses) node.Node {
	// TODO: uses has refine container(s)
	return ReflectChild(data)
}

func (self schema) Cases(choice *meta.Choice) node.Node {
	s := &Basic{
		Peekable: choice,
	}
	i := listIterator{dataList: choice, resolve: self.resolve}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		key := r.Key
		var choiceCase *meta.ChoiceCase
		if r.New {
			choiceCase = &meta.ChoiceCase{}
			choice.AddMeta(choiceCase)
		} else {
			if more, err := i.iterate(r.Selection, r.Meta, key, r.First, r.Row); err != nil {
				return nil, nil, err
			} else if more {
				choiceCase = i.data.(*meta.ChoiceCase)
				if key, err = node.NewValues(r.Meta.KeyMeta(), choiceCase.Ident); err != nil {
					return nil, nil, err
				}
			}
		}
		if choiceCase != nil {
			return self.MetaList(choiceCase), key, nil
		}
		return nil, nil, nil
	}
	return s
}

func (self schema) Choice(data *meta.Choice) node.Node {
	return &Extend{
		Base: ReflectChild(data),
		OnChild: func(parent node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "cases":
				// TODO: Not sure how to do create w/o what type to create
				return self.Cases(data), nil
			}
			return nil, nil
		},
	}
}

type listIterator struct {
	data     meta.Meta
	dataList meta.MetaList
	iterator meta.Iterator
	resolve  bool
	temp     int
}

func (i *listIterator) iterate(sel node.Selection, m *meta.List, key []val.Value, first bool, row int) (bool, error) {
	i.data = nil
	if i.dataList == nil {
		return false, nil
	}
	if len(key) > 0 {
		sel.Path.SetKey(key)
		if first {
			var err error
			i.data, err = meta.Find(i.dataList, key[0].String())
			if err != nil {
				return false, err
			}
		}
	} else {
		if first {
			if i.resolve {
				i.iterator = meta.Children(i.dataList)
			} else {
				i.iterator = meta.ChildrenNoResolve(i.dataList)
			}
			for j := 0; j < row && i.iterator.HasNext(); j++ {
			}
		}
		if i.iterator.HasNext() {
			var err error
			i.data, err = i.iterator.Next()
			if err != nil {
				return false, err
			}
			if i.data == nil {
				panic(fmt.Sprintf("Bad iterator at %s, item number %d", sel.Path.String(), i.temp))
			}
			sel.Path.SetKey([]val.Value{val.String(i.data.GetIdent())})
		}
	}
	return i.data != nil, nil
}

func (self schema) Definition(parent meta.MetaList, data meta.Meta) node.Node {
	s := &Basic{
		Peekable: data,
	}
	s.OnChoose = func(state node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		caseType := self.DefinitionType(data)
		return choice.GetCase(caseType)
	}
	s.OnChild = func(r node.ChildRequest) (node.Node, error) {
		if r.New {
			data = self.createGroupingsTypedefsDefinitions(parent, r.Meta)
		}
		if data == nil {
			return nil, nil
		}
		switch r.Meta.GetIdent() {
		case "anyxml":
			return self.Leaf(nil, nil, data.(*meta.Any)), nil
		case "leaf":
			return self.Leaf(data.(*meta.Leaf), nil, nil), nil
		case "leaf-list":
			return self.Leaf(nil, data.(*meta.LeafList), nil), nil
		case "uses":
			return self.Uses(data.(*meta.Uses)), nil
		case "choice":
			return self.Choice(data.(*meta.Choice)), nil
		case "rpc", "action":
			return self.Rpc(data.(*meta.Rpc)), nil
		default:
			return self.MetaList(data.(meta.MetaList)), nil
		}
		return nil, nil
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			if data != nil {
				err = node.WriteField(r.Meta, data, hnd.Val)
			}
		} else {
			hnd.Val, err = node.ReadField(r.Meta, data)
		}
		return
	}
	return s
}

func (self schema) Definitions(dataList meta.MetaList) node.Node {
	s := &Basic{
		Peekable: dataList,
	}
	i := listIterator{dataList: dataList, resolve: self.resolve}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		key := r.Key
		if r.New {
			return self.Definition(dataList, nil), key, nil
		} else {
			if more, err := i.iterate(r.Selection, r.Meta, r.Key, r.First, r.Row); err != nil {
				return nil, nil, err
			} else if more {
				if len(key) == 0 {
					if key, err = node.NewValues(r.Meta.KeyMeta(), i.data.GetIdent()); err != nil {
						return nil, nil, err
					}
				}
				return self.Definition(dataList, i.data), key, nil
			}
		}
		return nil, nil, nil
	}
	return s
}

func (self schema) DefinitionType(data meta.Meta) string {
	switch data.(type) {
	case *meta.List:
		return "list"
	case *meta.Uses:
		return "uses"
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
	default:
		return "container"
	}
}

var yangYang *meta.Module

func yangModule(ypath meta.StreamSource) *meta.Module {
	if yangYang == nil {
		var err error
		if yangYang, err = yang.LoadModule(ypath, "yang"); err != nil {
			panic(err)
		}
	}
	return yangYang
}
