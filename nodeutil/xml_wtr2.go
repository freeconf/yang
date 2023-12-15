package nodeutil

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

type XMLWtr2 struct {
	XMLName   xml.Name
	ns        string `xml:"-"`
	EnumAsIds bool   `xml:"-"`
	Content   string `xml:",innerxml"`
	Elem      []*XMLWtr2
}

func XmlName(d meta.Definition) xml.Name {
	return xml.Name{
		Local: d.Ident(),
		Space: meta.OriginalModule(d).Namespace(),
	}
}

// WriteXMLDoc includes the xml of the selection root
func WriteXMLDoc(s *node.Selection, pretty bool) (string, error) {
	ns := meta.OriginalModule(s.Meta()).Namespace()
	w := XMLWtr2{
		XMLName: xml.Name{
			Local: s.Meta().Ident(),
			Space: ns,
		},
		ns: meta.OriginalModule(s.Meta()).Namespace(),
	}
	if err := s.UpsertInto(&w); err != nil {
		return "", err
	}
	return writeXml(&w, pretty)
}

// WriteXMLFrag does not include the selection root
func WriteXMLFrag(s *node.Selection, pretty bool) (string, error) {
	var w XMLWtr2
	if err := s.UpsertInto(&w); err != nil {
		return "", err
	}
	if len(w.Elem) == 0 {
		return "", errors.New("no content to xml doc")
	}
	if len(w.Elem) > 1 {
		return "", fmt.Errorf("%d elements found and expected one. consider using WriteXMLDoc", len(w.Elem))
	}
	return writeXml(w.Elem[0], pretty)
}

func writeXml(w *XMLWtr2, pretty bool) (string, error) {
	var buf bytes.Buffer
	e := xml.NewEncoder(&buf)
	if pretty {
		e.Indent("", "  ")
	}
	if err := e.Encode(w); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (x *XMLWtr2) Child(r node.ChildRequest) (node.Node, error) {
	if r.New {
		if meta.IsList(r.Meta) {
			return x, nil
		}
		c := x.new(r.Meta)
		x.Elem = append(x.Elem, c)
		return c, nil
	}
	return nil, nil
}

func (x *XMLWtr2) new(m meta.Definition) *XMLWtr2 {
	w := XMLWtr2{
		EnumAsIds: x.EnumAsIds,
		ns:        meta.OriginalModule(m).Namespace(),
	}
	if w.ns != x.ns {
		w.XMLName = xml.Name{
			Local: m.Ident(),
			Space: w.ns,
		}
	} else {
		w.XMLName = xml.Name{
			Local: m.Ident(),
		}
	}
	return &w
}

func (x *XMLWtr2) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	if r.New {
		c := x.new(r.Meta)
		x.Elem = append(x.Elem, c)
		return c, r.Key, nil
	}
	return nil, nil, nil
}

func (x *XMLWtr2) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	var err error
	if r.Write {
		if hnd.Val.Format().IsList() {
			val.ForEach(hnd.Val, func(_ int, v val.Value) {
				if ferr := x.writeFieldElement(r.Meta, v); ferr != nil {
					err = ferr
				}
			})

		} else {
			err = x.writeFieldElement(r.Meta, hnd.Val)
		}
	}
	return err
}

func (x *XMLWtr2) writeFieldElement(m meta.Definition, v val.Value) error {
	c := x.new(m)
	switch v.Format() {
	case val.FmtIdentityRef:
		id := v.String()
		leafMod := meta.OriginalModule(m)
		bases := m.(meta.HasType).Type().Base()
		idty := meta.FindIdentity(bases, id)
		if idty == nil {
			return fmt.Errorf("could not find ident '%s'", id)
		}
		idtyMod := meta.RootModule(idty)
		if idtyMod != leafMod {
			c.Content = fmt.Sprintf("%s:%s", idtyMod.Ident(), id)
		} else {
			c.Content = id
		}
	case val.FmtEnum:
		if x.EnumAsIds {
			c.Content = strconv.Itoa(v.(val.Enum).Id)
		} else {
			c.Content = v.String()
		}
	case val.FmtDecimal64:
		f := v.Value().(float64)
		c.Content = strconv.FormatFloat(f, 'f', -1, 64)
	default:
		//case val.FmtString, val.FmtBinary, val.FmtAny:
		c.Content = v.String()
	}
	x.Elem = append(x.Elem, c)
	return nil
}

// Begin Stubs

func (x *XMLWtr2) Choose(sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
	return nil, fc.NotImplementedError
}

func (x *XMLWtr2) BeginEdit(r node.NodeRequest) error {
	return nil
}

func (x *XMLWtr2) EndEdit(r node.NodeRequest) error {
	return nil
}

func (x *XMLWtr2) Action(r node.ActionRequest) (node.Node, error) {
	return nil, fc.NotImplementedError
}

func (x *XMLWtr2) Notify(r node.NotifyRequest) (node.NotifyCloser, error) {
	return nil, fc.NotImplementedError
}

func (x *XMLWtr2) Peek(sel *node.Selection, consumer interface{}) interface{} {
	return x
}

func (x *XMLWtr2) Context(sel *node.Selection) context.Context {
	return sel.Context
}

func (x *XMLWtr2) Release(sel *node.Selection) {
}
