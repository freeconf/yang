package nodeutil

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"

	"bytes"

	"github.com/freeconf/yang/meta"
)

const QUOTE1 = '"'
const XML1 = '<'
const XML_CLOSE = '/'
const XML2 = '>'

type XMLWtr struct {

	// stream to write contents.  contents will be flushed only at end of operation
	Out io.Writer

	// otherwise enumerations are written as their labels.  it may be
	// useful to know that json reader can accept labels or values
	EnumAsIds bool

	_out *bufio.Writer
}

func WriteXML(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr := &XMLWtr{Out: buff}
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func (wtr XMLWtr) XML(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr.Out = buff
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func NewXMLWtr(out io.Writer) *XMLWtr {
	return &XMLWtr{Out: out}
}

func (wtr *XMLWtr) Node() node.Node {
	wtr._out = bufio.NewWriter(wtr.Out)

	return &Extend{
		Base: wtr.container(0),
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			println("Extend OnEndEdit: " + wtr.ident(r.Selection.Path))
			ident := wtr.ident(r.Selection.Path)
			if !meta.IsLeaf(r.Selection.Meta()) {
				if err := wtr.endContainer(ident); err != nil {
					return err
				}
			}
			if err := wtr._out.Flush(); err != nil {
				return err
			}
			return nil
		},
	}
}

func (wtr *XMLWtr) container(lvl int) node.Node {
	first := true
	s := &Basic{}
	s.OnChild = func(r node.ChildRequest) (child node.Node, err error) {
		if !r.New {
			return nil, nil
		}
		if !meta.IsList(r.Meta) {
			println("onchild:" + wtr.ident(r.Path))
			if err = wtr.beginContainer(wtr.ident(r.Path)); err != nil {
				return nil, err
			}
		}

		return wtr.container(lvl + 1), nil
	}
	s.OnBeginEdit = func(r node.NodeRequest) error {
		ident := wtr.ident(r.Selection.Path)
		println("OnBeginEdit: " + ident)
		if !meta.IsLeaf(r.Selection.Meta()) && !r.Selection.InsideList && !meta.IsList(r.Selection.Meta()) {
			if lvl == 0 && first == true {
				ident = wtr.ident(r.Selection.Path) + " xmlns=" + "\"" + meta.OriginalModule(r.Selection.Path.Meta).Ident() + "\""
				println("OnBeginEdit: " + ident)
				if err := wtr.beginContainer(ident); err != nil {
					return err
				}
			}
			first = false
		}
		return nil
	}
	s.OnEndEdit = func(r node.NodeRequest) error {
		println("OnEndEdit: " + wtr.ident(r.Selection.Path))
		//if !meta.IsList(r.Selection.Meta()) {
		//if !r.Selection.IsList {
		if r.Selection.InsideList {
			if err := wtr.endContainer(wtr.ident(r.Selection.Path)); err != nil {
				return err
			}
		} else if !meta.IsList(r.Selection.Meta()) {
			if err := wtr.endContainer(wtr.ident(r.Selection.Path)); err != nil {
				return err
			}
		}

		return nil
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		println("OnField: " + wtr.ident(r.Path) + "  Value: " + hnd.Val.String())
		space := ""
		if lvl == 0 && first == true {
			space = meta.OriginalModule(r.Path.Meta).Ident()
		}
		wtr.writeLeafElement(space, r.Path, hnd.Val)
		return nil
	}
	s.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		if !r.New {
			return
		}
		println("OnNext: " + wtr.ident(r.Selection.Path))

		ident := wtr.ident(r.Selection.Path)

		if err = wtr.beginContainer(ident); err != nil {
			return
		}
		return wtr.container(lvl + 1), r.Key, nil
	}
	return s
}

func (wtr *XMLWtr) ident(p *node.Path /*, first bool*/) string {
	s := p.Meta.(meta.Identifiable).Ident()
	return s
}

func (wtr *XMLWtr) beginContainer(ident string) (err error) {
	if err = wtr.writeOpenIdent(ident); err != nil {
		return err
	}
	return
}

func (wtr *XMLWtr) writeOpenIdent(ident string) (err error) {
	if _, err = wtr._out.WriteRune(XML1); err != nil {
		return err
	}
	if _, err = wtr._out.WriteString(ident); err != nil {
		return err
	}
	if _, err = wtr._out.WriteRune(XML2); err != nil {
		return err
	}
	return nil
}

func (wtr *XMLWtr) writeCloseIdent(ident string) (err error) {
	if _, err = wtr._out.WriteRune(XML1); err != nil {
		return err
	}
	if _, err = wtr._out.WriteRune(XML_CLOSE); err != nil {
		return err
	}
	if _, err = wtr._out.WriteString(ident); err != nil {
		return err
	}
	if _, err = wtr._out.WriteRune(XML2); err != nil {
		return err
	}
	return nil
}

func (wtr *XMLWtr) endContainer(ident string) (err error) {
	if err = wtr.writeCloseIdent(ident); err != nil {
		return
	}
	return
}

func (wtr *XMLWtr) writeLeafElement(attibute string, p *node.Path, v val.Value) error {
	var err error
	stringValue, err := wtr.getStringValue(p, v)
	ident := p.Meta.(meta.Identifiable).Ident()
	test := xml.StartElement{Name: xml.Name{Local: ident, Space: attibute}}
	xml.NewEncoder(wtr._out).EncodeElement(stringValue, test)

	return err
}

func (wtr *XMLWtr) getStringValue(p *node.Path, v val.Value) (string, error) {
	stringValue := ""
	var err error
	switch v.Format() {
	case val.FmtIdentityRef:
		stringValue = v.String()
		leafMod := meta.OriginalModule(p.Meta)
		base := p.Meta.(meta.HasType).Type().Base()
		idty, found := base.Derived()[stringValue]
		if !found {
			err = fmt.Errorf("could not find ident '%s'", stringValue)
		}
		idtyMod := meta.RootModule(idty)
		if idtyMod != leafMod {
			stringValue = fmt.Sprint(idtyMod.Ident(), ":", stringValue)
		}
	case val.FmtString, val.FmtBinary, val.FmtAny:
		stringValue = v.String()
	case val.FmtEnum:
		if wtr.EnumAsIds {
			stringValue = strconv.Itoa(v.(val.Enum).Id)

		} else {
			stringValue = v.(val.Enum).Label
		}
	case val.FmtDecimal64:
		f := v.Value().(float64)
		stringValue = strconv.FormatFloat(f, 'f', -1, 64)
	default:
		stringValue = v.String()
	}
	return stringValue, err
}

func (wtr *XMLWtr) writeString(s string) error {
	clean := bytes.NewBuffer(make([]byte, len(s)+2))
	clean.Reset()
	writeString(clean, s, true)
	_, ioErr := wtr._out.Write(clean.Bytes())
	return ioErr
}
