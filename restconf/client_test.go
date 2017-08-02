package restconf

import (
	"fmt"
	"io"
	"testing"

	"strings"

	"bytes"

	"io/ioutil"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
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

	s := b.sel(test.def, test.data)

	// read
	support.reset().node().Child(b.cr(s, "y"))
	checkEqual(t, "GET path=x", support.log())

	ls := b.sel(`list x { key "y"; leaf y { type string; } }`, `{"x":[{"y":"hi"}]}`)
	support.reset().node().Next(b.lr(ls, "hi"))
	checkEqual(t, "GET path=x", support.log())

	// nav
	nav := b.cr(s, "y")
	navPath, _ := node.ParsePath("y", s.Meta().(meta.MetaList))
	nav.Target = navPath.Tail
	if !nav.IsNavigation() {
		t.Error("assumed navigation mode")
	}
	support.reset().node().Child(nav)
	checkEqual(t, "OPTIONS path=x/y", support.log())

	// delete
	support.reset().node().Delete(b.nr(s))
	checkEqual(t, "DELETE path=x", support.log())

	// notify
	notifyDef := fmt.Sprintf(`notification x { %s }`, test.def)
	support.reset().node().Notify(b.nor(b.sel(notifyDef, test.data), support.stream))
	checkEqual(t, 1, len(support._subs))

	// action
	support.reset().node().Action(b.ar(b.sel(`action x { input { } }`, ""), s))
	checkEqual(t, `POST path=x payload={"y":{},"z":"hi"}`, support.log())

	// edit
	n := support.reset().node()
	nr := b.nr(s)
	n.BeginEdit(nr)
	checkEqual(t, "GET path=x params=depth=1&content=config&with-defaults=trim", support.log())

	n.Child(b.crw(s, "y"))
	checkEqual(t, "", support.log())

	n.Field(b.frw(s, "z", "hi"))
	checkEqual(t, "", support.log())

	n.EndEdit(nr)
	checkEqual(t, `PUT path=x payload={"y":{},"z":"hi"}`, support.log())
}

type testDriverSupport struct {
	_log        string
	doResponse  node.Node
	_subs       map[string]*clientSubscription
	ws          bytes.Buffer
	subPayloads bytes.Buffer
}

func (self *testDriverSupport) reset() *clientNode {
	self._log = ""
	self._subs = make(map[string]*clientSubscription)
	self.doResponse = &nodes.Basic{}
	self.ws.Reset()
	return &clientNode{support: self}
}

func (self *testDriverSupport) stream(payload node.Selection) {
	if err := payload.InsertInto(nodes.NewJsonWriter(&self.subPayloads).Node()).LastErr; err != nil {
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

func checkEqual(t *testing.T, a interface{}, b interface{}) {
	if err := c2.CheckEqual(a, b); err != nil {
		t.Error(err)
	}
}

type requestBuilder struct {
}

func (self requestBuilder) sel(y string, payloadJson string) node.Selection {
	return node.Selection{
		Node: self.dn(payloadJson),
		Path: node.NewRootPath(self.m(y).(meta.MetaList)),
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
	m, err := meta.FindByIdent2(s.Meta(), field)
	if err != nil {
		panic(err)
	}
	r := node.FieldRequest{
		Request: node.Request{
			Selection: s,
			Path:      s.Path,
		},
		Meta: m.(meta.HasDataType),
	}
	vv, err := node.NewValue(r.Meta.GetDataType(), v)
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
	m, err := meta.FindByIdent2(s.Meta(), child)
	if err != nil {
		panic(err)
	}
	return node.ChildRequest{
		Meta: m.(meta.MetaList),
		Request: node.Request{
			Selection: s,
		},
	}
}

func (requestBuilder) dn(payloadJson string) node.Node {
	return nodes.NewJsonReader(strings.NewReader(payloadJson)).Node()
}

func (requestBuilder) m(y string) meta.Meta {
	mstr := fmt.Sprint(`module m { namespace ""; prefix ""; revision 0; `, y, `}`)
	m := yang.RequireModuleFromString(nil, mstr)
	l := m.DataDefs()
	// heuristic; if there's only one item, assume that's the one they want
	if meta.ListLen(l) == 1 {
		return l.GetFirstMeta()
	}
	// otherwise if there's more, assume they want the module
	return m
}
