package restconf

import (
	"fmt"
	"io"
	"testing"

	"bytes"

	"io/ioutil"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
)

func Test_Client(t *testing.T) {
	support := &testDriverSupport{}
	b := requestBuilder{}
	test := struct {
		def  string
		data string
	}{
		def: `
			container x {
				container y {}
				leaf z { type string; }
			}`,
		data: `{"y":{},"z":"hi"}`,
	}

	s := b.sel(b.ddef(test.def), test.data)

	// read
	support.reset().node().Child(b.cr(s, "y"))
	c2.AssertEqual(t, "GET path=x", support.log())

	ls := b.sel(b.ddef(`list x { key "y"; leaf y { type string; } }`), `{"x":[{"y":"hi"}]}`)
	support.reset().node().Next(b.lr(ls, "hi"))
	c2.AssertEqual(t, "GET path=x", support.log())

	// nav
	nav := b.cr(s, "y")
	navPath, _ := node.ParsePath("y", s.Meta().(meta.HasDefinitions))
	nav.Target = navPath.Tail
	if !nav.IsNavigation() {
		t.Error("assumed navigation mode")
	}
	support.reset().node().Child(nav)
	c2.AssertEqual(t, "OPTIONS path=x/y", support.log())

	// delete
	support.reset().node().Delete(b.nr(s))
	c2.AssertEqual(t, "DELETE path=x", support.log())

	// notify
	notifyDef := fmt.Sprintf(`notification x { %s }`, test.def)
	support.reset().node().Notify(b.nor(b.sel(b.notify(notifyDef), test.data), support.stream))
	c2.AssertEqual(t, 1, len(support._subs))

	// action
	support.reset().node().Action(b.ar(b.sel(b.action(`action x { input { } }`), ""), s))
	c2.AssertEqual(t, `POST path=x payload={"y":{},"z":"hi"}`, support.log())

	// edit
	n := support.reset().node()
	nr := b.nr(s)
	n.BeginEdit(nr)
	c2.AssertEqual(t, "GET path=x params=depth=1&content=config&with-defaults=trim", support.log())

	n.Child(b.crw(s, "y"))
	c2.AssertEqual(t, "", support.log())

	n.Field(b.frw(s, "z", "hi"))
	c2.AssertEqual(t, "", support.log())

	n.EndEdit(nr)
	c2.AssertEqual(t, `PUT path=x payload={"y":{},"z":"hi"}`, support.log())
}

type testDriverSupport struct {
	_log        string
	doResponse  node.Node
	_subs       map[string]*clientSubscription
	ws          bytes.Buffer
	subPayloads string
}

func (self *testDriverSupport) reset() *clientNode {
	self._log = ""
	self._subs = make(map[string]*clientSubscription)
	self.doResponse = &nodes.Basic{}
	self.ws.Reset()
	return &clientNode{support: self}
}

func (self *testDriverSupport) stream(payload node.Selection) {
	var err error
	if self.subPayloads, err = nodes.WriteJSON(payload); err != nil {
		panic(err)
	}
}

func (self *testDriverSupport) log() string {
	s := self._log
	self._log = ""
	return s
}

func (self *testDriverSupport) clientDo(method string, params string, p *node.Path, payload io.Reader) (node.Node, error) {
	self._log += fmt.Sprintf("%s path=%s", method, p.String())
	if params != "" {
		self._log += " params=" + params
	}
	if payload != nil {
		if payloadBytes, err := ioutil.ReadAll(payload); err != nil {
			panic(err)
		} else if len(payloadBytes) > 0 {
			self._log += fmt.Sprintf(" payload=%s", string(payloadBytes))
		}
	}
	return self.doResponse, nil
}

func (self *testDriverSupport) clientSubscriptions() map[string]*clientSubscription {
	return self._subs
}

func (self *testDriverSupport) clientSocket() (io.Writer, error) {
	return &self.ws, nil
}

type requestBuilder struct {
}

func (self requestBuilder) sel(d meta.Definition, payloadJson string) node.Selection {
	return node.Selection{
		Constraints: &node.Constraints{},
		Node:        self.dn(payloadJson),
		Path:        node.NewRootPath(d),
	}
}

func (requestBuilder) lr(s node.Selection, key interface{}) node.ListRequest {
	r := node.ListRequest{
		Request: node.Request{
			Selection: s,
			Path:      s.Path,
		},
		Meta: s.Meta().(*meta.List),
	}
	if key != nil {
		var err error
		r.Key, err = node.NewValues(r.Meta.KeyMeta(), key)
		if err != nil {
			panic(err)
		}
	}
	return r
}

func (self requestBuilder) frw(s node.Selection, field string, v interface{}) (node.FieldRequest, *node.ValueHandle) {
	r, h := self.fr(s, field, v)
	r.Write = true
	return r, h
}

func (requestBuilder) fr(s node.Selection, field string, v interface{}) (node.FieldRequest, *node.ValueHandle) {
	m := meta.Find(s.Meta().(meta.HasDefinitions), field)
	if m == nil {
		panic("no field " + field)
	}
	r := node.FieldRequest{
		Request: node.Request{
			Selection: s,
			Path:      s.Path,
		},
		Meta: m.(meta.HasDataType),
	}
	vv, err := node.NewValue(r.Meta.DataType(), v)
	if err != nil {
		panic(err)
	}
	return r, &node.ValueHandle{Val: vv}
}

func (requestBuilder) ar(s node.Selection, in node.Selection) node.ActionRequest {
	return node.ActionRequest{
		Request: node.Request{
			Selection: s,
			Path:      s.Path,
		},
		Meta:  s.Meta().(*meta.Rpc),
		Input: in,
	}
}

func (requestBuilder) nor(s node.Selection, stream node.NotifyStream) node.NotifyRequest {
	return node.NotifyRequest{
		Request: node.Request{
			Selection: s,
			Path:      s.Path,
		},
		Meta:   s.Meta().(*meta.Notification),
		Stream: stream,
	}
}

func (requestBuilder) nr(s node.Selection) node.NodeRequest {
	return node.NodeRequest{
		Selection: s,
		Source:    s,
		EditRoot:  true,
	}
}

func (self requestBuilder) crw(s node.Selection, child string) node.ChildRequest {
	r := self.cr(s, child)
	r.New = true
	return r
}

func (requestBuilder) cr(s node.Selection, child string) node.ChildRequest {
	m := meta.Find(s.Meta().(meta.HasDefinitions), child)
	if m == nil {
		panic(child + " not found")
	}
	return node.ChildRequest{
		Meta: m.(meta.HasDataDefs),
		Request: node.Request{
			Selection: s,
		},
	}
}

func (requestBuilder) dn(payloadJson string) node.Node {
	return nodes.ReadJSON(payloadJson)
}

func (self requestBuilder) notify(y string) *meta.Notification {
	for _, n := range self.m(y).Notifications() {
		return n
	}
	panic("no notification")
}

func (self requestBuilder) action(y string) *meta.Rpc {
	for _, n := range self.m(y).Actions() {
		return n
	}
	panic("no actions")
}

func (self requestBuilder) ddef(y string) meta.DataDef {
	return self.m(y).DataDefs()[0]
}

func (requestBuilder) m(y string) *meta.Module {
	mstr := fmt.Sprint(`module m { namespace ""; prefix ""; revision 0; `, y, `}`)
	return yang.RequireModuleFromString(nil, mstr)
}
