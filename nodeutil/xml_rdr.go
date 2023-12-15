package nodeutil

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/patch/xml"
	"github.com/freeconf/yang/val"
)

// ReadXMLDoc to read xml doc where root node is assumed to be the correct
// element.
//
//	<my-module>
//	   <my-child>...
//
//	module my-module {
//	    container my-child {...
//
//	NewBrowser(myModule, ReadXMLDoc(myXmlDoc))
func ReadXMLDoc(buf io.Reader) (*XmlNode, error) {
	dec := xml.NewDecoder(buf)
	var n XmlNode
	err := dec.Decode(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// ReadXMLBlock to read xml doc where root node is assumed to be the correct
// element.
//
//	<my-child>
//	   <my-item>...
//
//	module my-module {
//	    container my-child {...
//
//	b := NewBrowser(myModule, myApp)
//	sel, _ := b.Root().Find("my-child")
//	sel.UpsertFrom(ReadXMLBlock(myXmlDoc))
func ReadXMLBlock(buf io.Reader) (*XmlNode, error) {
	n, err := ReadXMLDoc(buf)
	if err != nil {
		return nil, err
	}
	return &XmlNode{Nodes: []*XmlNode{n}}, nil
}

type XmlNode struct {
	XMLName xml.Name
	Content []byte     `xml:",chardata"`
	Nodes   []*XmlNode `xml:",any"`

	// Even tho attributes are not used here, make this avail. to other systems
	// like netconf edit actions
	Attr []xml.Attr `xml:",any,attr"`
}

func (x *XmlNode) Child(r node.ChildRequest) (node.Node, error) {
	ndx := x.Find(0, r.Meta)
	if ndx < 0 {
		return nil, nil
	}
	if meta.IsList(r.Meta) {
		var found []*XmlNode
		// 7.8.5.  XML Encoding Rules
		// The XML elements representing list entries MAY be interleaved with elements
		// for siblings of the list
		for ndx >= 0 {
			found = append(found, x.Nodes[ndx])
			ndx = x.Find(ndx+1, r.Meta)
		}
		return &XmlNode{XMLName: x.XMLName, Nodes: found}, nil
	}
	return x.Nodes[ndx], nil
}

func (x *XmlNode) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	if r.Key != nil {
		for _, n := range x.Nodes {
			for i, k := range r.Key {
				v, found := n.field(r.Meta.KeyMeta()[i])
				if !found {
					break
				}
				if k.String() != v {
					break
				}
				isLastKey := i == (len(r.Key) - 1)
				if isLastKey {
					return n, r.Key, nil
				}
			}
		}
	} else if r.Row < len(x.Nodes) {
		// assumes x.Nodes are all of the same element tag as would be the case if
		// this was created by Child() method
		target := x.Nodes[r.Row]
		var key []val.Value
		if len(r.Meta.KeyMeta()) > 0 {
			for _, kmeta := range r.Meta.KeyMeta() {
				sval, valid := target.field(kmeta)
				if !valid {
					return nil, nil, fmt.Errorf("key '%s' missing from %s", kmeta.Ident(), r.Path)
				}
				v, err := node.NewValue(kmeta.Type(), sval)
				if err != nil {
					return nil, nil, fmt.Errorf("error reading key '%s' from %s. %w", kmeta.Ident(), r.Path, err)
				}
				key = append(key, v)
			}
		}

		return x.Nodes[r.Row], key, nil
	}
	return nil, nil, nil
}

func (x *XmlNode) ContentTrim() string {
	return strings.TrimSpace(string(x.Content))
}

func (x *XmlNode) field(m meta.Leafable) (string, bool) {
	if ndx := x.Find(0, m); ndx >= 0 {
		return x.Nodes[ndx].ContentTrim(), true
	}
	return "", false
}

func (x *XmlNode) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	var err error
	ndx := x.Find(0, r.Meta)
	if ndx < 0 {
		return nil
	}
	if _, isList := r.Meta.(*meta.LeafList); isList {
		var found []string
		// 7.8.5.  XML Encoding Rules
		// The XML elements representing list entries MAY be interleaved with elements
		// for siblings of the list
		for ndx >= 0 {
			found = append(found, x.Nodes[ndx].ContentTrim())
			ndx = x.Find(ndx+1, r.Meta)
		}
		hnd.Val, err = node.NewValue(r.Meta.Type(), found)
	} else {
		hnd.Val, err = node.NewValue(r.Meta.Type(), x.Nodes[ndx].ContentTrim())
	}
	return err
}

func (x *XmlNode) Find(start int, m meta.Definition) int {
	for i := start; i < len(x.Nodes); i++ {
		if x.Nodes[i].XMLName.Local == m.Ident() {
			if x.Nodes[i].XMLName.Space != "" {
				ns := meta.OriginalModule(m).Namespace()
				if x.Nodes[i].XMLName.Space != ns {
					continue
				}
			}
			return i
		}
	}
	return -1
}

func (x *XmlNode) Choose(sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
	for _, c := range choice.Cases() {
		for _, m := range c.DataDefinitions() {
			if x.Find(0, m) >= 0 {
				return c, nil
			}
		}
	}
	return nil, nil
}

// Stubs non-reader funcs

func (n *XmlNode) Peek(sel *node.Selection, consumer interface{}) interface{} {
	return n
}

func (n *XmlNode) BeginEdit(r node.NodeRequest) error {
	return nil
}

func (n *XmlNode) EndEdit(r node.NodeRequest) error {
	return nil
}

func (n *XmlNode) Action(r node.ActionRequest) (node.Node, error) {
	return nil, nil
}

func (n *XmlNode) Notify(r node.NotifyRequest) (node.NotifyCloser, error) {
	return nil, nil
}

func (n *XmlNode) Context(sel *node.Selection) context.Context {
	return sel.Context
}

func (n *XmlNode) Release(sel *node.Selection) {}
