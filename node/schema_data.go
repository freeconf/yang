package node

import (
	"fmt"
	"github.com/c2g/meta"
	"github.com/c2g/meta/yang"
)

/**
 * This is used to encode YANG models. In order to navigate the YANG model it needs a model
 * which is the YANG YANG model.  Note: It can be confusing which is the data and which is the
 * goober.
 */
type SchemaData struct {
	// resolve all uses, groups and typedefs.  if this is false, then depth must be
	// used to avoid infinite recursion
	Resolve bool
}

func SelectModule(m *meta.Module, resolve bool) *Browser {
	return NewBrowser(YangModule(),
		func() Node {
			return SchemaData{Resolve:resolve}.Yang(m)
		})
}

var yang1_0 *meta.Module

func init() {
	yang.InternalYang()["yang"] = `
module yang {
    namespace "http://meta.org/yang";
    prefix "meta";
    import yanglib;
    revision 0;

    uses module;
}
`
}
func YangModule() *meta.Module {
	if yang1_0 == nil {
		yang1_0 = yang.InternalModule("yang")
	}
	return yang1_0
}


type MetaListSelector func(m meta.Meta) (Node, error)

func (self SchemaData) Yang(module *meta.Module) Node {
	s := &MyNode{}
	s.OnSelect = func(r ContainerRequest) (Node, error) {
		switch r.Meta.GetIdent() {
		case "module":
			return self.Module(module), nil
		}
		return nil, nil
	}
	return s
}

func (self SchemaData) Module(module *meta.Module) (Node) {
	return &Extend{
		Label:"Module",
		Node:self.MetaList(module),
		OnSelect : func(parent Node, r ContainerRequest) (child Node, err error) {
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
			return parent.Select(r)
		},
	}
}

func (self SchemaData) Revision(rev *meta.Revision) (Node) {
	s := &MyNode{}
	s.OnRead = func(r FieldRequest) (*Value, error) {
		switch r.Meta.GetIdent() {
		case "rev-date":
			return &Value{Str: rev.Ident, Type:r.Meta.GetDataType()}, nil
		default:
			return ReadField(r.Meta, rev)
		}
	}
	s.OnWrite = func(r FieldRequest, val *Value) error {
		switch r.Meta.GetIdent() {
		case "rev-date":
			rev.Ident = val.Str
		default:
			return WriteField(r.Meta, rev, val)
		}
		return nil
	}
	return s
}

func (self SchemaData) Type(typeData *meta.DataType) (Node) {
	return &MyNode{
		OnRead: func(r FieldRequest) (*Value, error) {
			switch r.Meta.GetIdent()  {
			case "ident":
				return SetValue(r.Meta.GetDataType(), typeData.Ident)
			case "minLength":
				if self.Resolve || typeData.MinLengthPtr != nil {
					return SetValue(r.Meta.GetDataType(), typeData.MinLength())
				}
			case "maxLength":
				if self.Resolve || typeData.MaxLengthPtr != nil {
					return SetValue(r.Meta.GetDataType(), typeData.MaxLength())
				}
			case "path":
				if self.Resolve || typeData.PathPtr != nil {
					return SetValue(r.Meta.GetDataType(), typeData.Path())
				}
			case "enumeration":
				if self.Resolve || len(typeData.EnumerationRef) > 0 {
					return SetValue(r.Meta.GetDataType(), typeData.Enumeration())
				}
			}
			return nil, nil
		},
		OnWrite: func(r FieldRequest, val *Value) error {
			switch r.Meta.GetIdent() {
			case "ident":
				typeData.Ident = val.Str
				typeData.SetFormat(meta.DataTypeImplicitFormat(val.Str))
			case "minLength":
				typeData.SetMinLength(val.Int)
			case "maxLength":
				typeData.SetMaxLength(val.Int)
			case "path":
				typeData.SetPath(val.Str)
			case "enumeration":
				typeData.SetEnumeration(val.Strlist)
			}
			return nil
		},
	}
}

func (self SchemaData) Groupings(groupings meta.MetaList) (Node) {
	s := &MyNode{}
	i := listIterator{dataList: groupings, resolve: self.Resolve}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		var key = r.Key
		var group *meta.Grouping
		if r.New {
			group = &meta.Grouping{Ident:r.Key[0].Str}
			groupings.AddMeta(group)
		} else {
			if i.iterate(r.Selection, r.Meta, r.Key, r.First, r.Row) {
				group = i.data.(*meta.Grouping)
				if len(key) == 0 {
					key = SetValues(r.Meta.KeyMeta(), group.Ident)
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

func (self SchemaData) RpcIO(i *meta.RpcInput, o *meta.RpcOutput) (Node) {
	var io meta.MetaList
	if i != nil {
		io = i
	} else {
		io = o
	}
	return self.MetaList(io)
}

func (self SchemaData) createGroupingsTypedefsDefinitions(parent meta.MetaList, childMeta meta.Meta) (meta.Meta) {
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

func (self SchemaData) Rpc(rpc *meta.Rpc) (Node) {
	return &Extend{
		Label:"rpc",
		Node: MarshalContainer(rpc),
		OnSelect: func(parent Node, r ContainerRequest) (Node, error) {
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
			return parent.Select(r)
		},
	}
}

func (self SchemaData) Typedefs(typedefs meta.MetaList) (Node) {
	s := &MyNode{}
	i := listIterator{dataList: typedefs, resolve: self.Resolve}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		var key = r.Key
		var typedef *meta.Typedef
		if r.New {
			typedef = &meta.Typedef{Ident:r.Key[0].Str}
			typedefs.AddMeta(typedef)
		} else {
			if i.iterate(r.Selection, r.Meta, r.Key, r.First, r.Row) {
				typedef = i.data.(*meta.Typedef)
				if len(key) == 0 {
					key = SetValues(r.Meta.KeyMeta(), typedef.Ident)
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

func (self SchemaData) Typedef(typedef *meta.Typedef) Node {
	return &Extend{
		Label:"Typedef",
		Node: MarshalContainer(typedef),
		OnSelect :func(parent Node, r ContainerRequest) (Node, error) {
			switch r.Meta.GetIdent() {
			case "type":
				if r.New {
					typedef.SetDataType(&meta.DataType{Parent:typedef})
				}
				if typedef.DataType != nil {
					return self.Type(typedef.DataType), nil
				}
			}
			return nil, nil
		},
	}
}

func (self SchemaData) MetaList(data meta.MetaList) (Node) {
	return &Extend{
		Label: "MetaList",
		Node: MarshalContainer(data),
		OnSelect : func(parent Node, r ContainerRequest) (Node, error) {
			hasGroupings, implementsHasGroupings := data.(meta.HasGroupings)
			hasTypedefs, implementsHasTypedefs := data.(meta.HasTypedefs)
			switch r.Meta.GetIdent() {
			case "groupings":
				if ! self.Resolve && implementsHasGroupings {
					groupings := hasGroupings.GetGroupings()
					if r.New || ! meta.ListEmpty(groupings) {
						return self.Groupings(groupings), nil
					}
				}
				return nil, nil
			case "typedefs":
				if ! self.Resolve && implementsHasTypedefs {
					typedefs := hasTypedefs.GetTypedefs()
					if r.New || ! meta.ListEmpty(typedefs) {
						return self.Typedefs(typedefs), nil
					}
				}
				return nil, nil
			case "definitions":
				defs := data.(meta.MetaList)
				if r.New || ! meta.ListEmpty(defs) {
					return self.Definitions(defs), nil
				}
				return nil, nil
			}
			return parent.Select(r)
		},
	}
}

func (self SchemaData) Leaf(leaf *meta.Leaf, leafList *meta.LeafList, any *meta.Any) (Node) {
	var leafy meta.HasDataType
	if leaf != nil {
		leafy = leaf
	} else if leafList != nil {
		leafy = leafList
	} else {
		leafy = any
	}
	s := &MyNode{
		Peekable: leafy,
	}
	details := leafy.(meta.HasDetails).Details()
	s.OnSelect = func(r ContainerRequest) (Node, error) {
		switch r.Meta.GetIdent() {
		case "type":
			if r.New {
				leafy.SetDataType(&meta.DataType{Parent:leafy})
			}
			if leafy.GetDataType() != nil {
				return self.Type(leafy.GetDataType()), nil
			}
		}
		return nil, nil
	}
	s.OnRead = func(r FieldRequest) (*Value, error) {
		switch r.Meta.GetIdent() {
		case "config":
			if self.Resolve || details.ConfigPtr != nil {
				return &Value{Bool: details.Config(r.Selection.Path()), Type:r.Meta.GetDataType()}, nil
			}
		case "mandatory":
			if self.Resolve || details.MandatoryPtr != nil {
				return &Value{Bool: details.Mandatory(), Type:r.Meta.GetDataType()}, nil
			}
		default:
			return ReadField(r.Meta, leafy)
		}
		return nil, nil
	}
	s.OnWrite = func(r FieldRequest, val *Value) error {
		switch r.Meta.GetIdent() {
		case "config":
			details.SetConfig(val.Bool)
		case "mandatory":
			details.SetMandatory(val.Bool)
		default:
			return WriteField(r.Meta, leafy, val)
		}
		return nil
	}
	return s
}

func (self SchemaData) Uses(data *meta.Uses) (Node) {
	// TODO: uses has refine container(s)
	return MarshalContainer(data)
}

func (self SchemaData) Cases(choice *meta.Choice) (Node) {
	s := &MyNode{
		Peekable: choice,
	}
	i := listIterator{dataList: choice, resolve: self.Resolve}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		key := r.Key
		var choiceCase *meta.ChoiceCase
		if r.New {
			choiceCase = &meta.ChoiceCase{}
			choice.AddMeta(choiceCase)
		} else {
			if i.iterate(r.Selection, r.Meta, key, r.First, r.Row) {
				choiceCase = i.data.(*meta.ChoiceCase)
				key = SetValues(r.Meta.KeyMeta(), choiceCase.Ident)
			}
		}
		if choiceCase != nil {
			return self.MetaList(choiceCase), key, nil
		}
		return nil, nil, nil
	}
	return s
}

func (self SchemaData) Choice(data *meta.Choice) (Node) {
	return &Extend{
		Label:"Choice",
		Node: MarshalContainer(data),
		OnSelect: func(parent Node, r ContainerRequest) (Node, error) {
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
	iterator meta.MetaIterator
	resolve  bool
	temp     int
}

func (i *listIterator) iterate(sel *Selection, m *meta.List, key []*Value, first bool, row int) bool {
	i.data = nil
	if i.dataList == nil {
		return false
	}
	if len(key) > 0 {
		sel.path.key = key
		if first {
			i.data = meta.FindByIdent2(i.dataList, key[0].Str)
		}
	} else {
		if first {
			i.iterator = meta.NewMetaListIterator(i.dataList, i.resolve)
			for j := 0; j < row && i.iterator.HasNextMeta(); j++ {
			}
		}
		if i.iterator.HasNextMeta() {
			i.data = i.iterator.NextMeta()
			if i.data == nil {
				panic(fmt.Sprintf("Bad iterator at %s, item number %d", sel.String(), i.temp))
			}
			sel.path.key = SetValues(m.KeyMeta(), i.data.GetIdent())
		}
	}
	return i.data != nil
}

func (self SchemaData) Definition(parent meta.MetaList, data meta.Meta) (Node) {
	s := &MyNode{
		Peekable: data,
	}
	s.OnChoose = func(state *Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		caseType := self.DefinitionType(data)
		return choice.GetCase(caseType), nil
	}
	s.OnSelect = func(r ContainerRequest) (Node, error) {
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
	s.OnRead = func(r FieldRequest) (*Value, error) {
		return ReadField(r.Meta, data)
	}
	s.OnWrite = func(r FieldRequest, val *Value) (err error) {
		switch r.Meta.GetIdent() {
		case "ident":
			// if data is nil then we're creating a def and we'll get name again

			if data != nil {
				return WriteField(r.Meta, data, val)
			}
		default:
			return WriteField(r.Meta, data, val)
		}
		return nil
	}
	return s
}

func (self SchemaData) Definitions(dataList meta.MetaList) (Node) {
	s := &MyNode{
		Peekable: dataList,
	}
	i := listIterator{dataList: dataList, resolve: self.Resolve}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		key := r.Key
		if r.New {
			return self.Definition(dataList, nil), key, nil
		} else {
			if i.iterate(r.Selection, r.Meta, r.Key, r.First, r.Row) {
				if len(key) == 0 {
					key = SetValues(r.Meta.KeyMeta(), i.data.GetIdent())
				}
				return self.Definition(dataList, i.data), key, nil
			}
		}
		return nil, nil, nil
	}
	return s
}

func (self SchemaData) DefinitionType(data meta.Meta) string {
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
