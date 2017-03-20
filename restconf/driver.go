package restconf

import (
	"bytes"
	"context"
	"fmt"

	"io"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/node"
)

// driverSupport is interface between Device and driver.  Factored out as part of
// testing but also because a lot of what driver does is potentially universal to proxying
// for other protocols and might allow reusablity when other protocols are added
type driverSupport interface {
	driverSubs() map[string]*driverSub
	driverDo(method string, params string, p *node.Path, payload io.Reader) (node.Node, error)
	driverWebsocket() (io.Writer, error)
}

type driver struct {
	support driverSupport
	params  string
	read    node.Node
	edit    node.Node
	method  string
	changes node.Node
}

var noSelection node.Selection

func (self *driver) inProcess(r node.NodeRequest) bool {
	return !r.Source.Path.Equal(r.Selection.Path)
}

func (self *driver) node() node.Node {
	n := &node.MyNode{}
	n.OnBeginEdit = func(r node.NodeRequest) error {
		if self.inProcess(r) {
			return nil
		}
		if r.New {
			self.method = "POST"
		} else {
			self.method = "PUT"
		}
		return self.startEditMode(r.Selection.Path)
	}
	n.OnChild = func(r node.ChildRequest) (node.Node, error) {
		if r.IsNavigation() {
			return self.startNavigation(r.Context, r.Target, n)
		} else if self.edit != nil {
			return self.edit.Child(r)
		}
		if self.read == nil {
			if err := self.startReadMode(r.Selection.Path); err != nil {
				return nil, err
			}
		}
		return self.read.Child(r)
	}
	n.OnDelete = func(r node.NodeRequest) error {
		_, err := self.request(r.Context, "DELETE", r.Selection.Path, noSelection)
		return err
	}
	n.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if r.IsNavigation() {
			n, err := self.startNavigation(r.Context, r.Target, n)
			return n, r.Key, err
		} else if self.edit != nil {
			return self.edit.Next(r)
		}
		if self.read == nil {
			if err := self.startReadMode(r.Selection.Path); err != nil {
				return nil, nil, err
			}
		}
		return self.read.Next(r)
	}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		if r.IsNavigation() {
			return nil
		} else if self.edit != nil {
			return self.edit.Field(r, hnd)
		}
		if self.read == nil {
			if err := self.startReadMode(r.Selection.Path); err != nil {
				return err
			}
		}
		return self.read.Field(r, hnd)
	}
	n.OnNotify = func(r node.NotifyRequest) (node.NotifyCloser, error) {
		sub := &driverSub{
			sel:    r.Selection,
			stream: r.Stream,
		}
		ws, err := self.support.driverWebsocket()
		if err != nil {
			return nil, err
		}
		if err := sub.open(ws); err != nil {
			return nil, err
		}
		subKey := r.Path.StringNoModule()
		self.support.driverSubs()[subKey] = sub
		closer := func() error {
			sub.close(ws)
			delete(self.support.driverSubs(), subKey)
			return nil
		}
		return closer, nil
	}
	n.OnAction = func(r node.ActionRequest) (node.Node, error) {
		return self.request(r.Context, "POST", r.Selection.Path, r.Input)
	}
	n.OnEndEdit = func(r node.NodeRequest) error {
		// send request
		if self.inProcess(r) {
			return nil
		}
		_, err := self.request(r.Context, self.method, r.Selection.Path, r.Selection.Split(self.changes))
		return err
	}
	return n
}

func (self *driver) startReadMode(path *node.Path) (err error) {
	self.read, err = self.get(path, self.params)
	return
}

func (self *driver) startEditMode(path *node.Path) error {
	// add depth = 1 so we can pull first level containers and
	// know what container would be conflicts.  we'll have to pull field
	// values too because there's no url param to exclude those yet.
	params := "depth=1&content=config&with-defaults=trim"
	existing, err := self.get(path, params)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	self.changes = node.MapNode(data)
	self.edit = &node.Extend{
		Node: self.changes,
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			if !r.New && existing != nil {
				return existing.Child(r)
			}
			return p.Child(r)
		},
	}
	return nil
}

// we stay inside this node until we're not navigating or remote endpoint
// doesn't exist
func (self *driver) startNavigation(c context.Context, target *node.Path, targetNode node.Node) (node.Node, error) {
	_, err := self.request(c, "OPTIONS", target, noSelection)
	if herr, ok := err.(c2.HttpError); ok {
		if herr.HttpCode() == 404 {
			return nil, nil
		}
		return nil, err
	}
	e := &node.MyNode{}
	e.OnChild = func(r node.ChildRequest) (node.Node, error) {
		if !r.IsNavigation() {
			return targetNode, nil
		}
		return e, nil
	}
	e.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if !r.IsNavigation() {
			return targetNode, r.Key, nil
		}
		return e, r.Key, nil
	}
	return e, nil
}

func (self *driver) get(p *node.Path, params string) (node.Node, error) {
	return self.support.driverDo("GET", params, p, nil)
}

func (self *driver) request(c context.Context, method string, p *node.Path, in node.Selection) (node.Node, error) {
	var payload bytes.Buffer
	if !in.IsNil() {
		js := node.NewJsonWriter(&payload).Node()
		if err := in.InsertInto(c, js).LastErr; err != nil {
			return nil, err
		}
		c2.Debug.Printf("got %d bytes in payload", payload.Len())
	}
	return self.support.driverDo(method, "", p, &payload)
}

type driverSub struct {
	sel    node.Selection
	stream node.NotifyStream
}

func (self *driverSub) close(ws io.Writer) error {
	_, err := ws.Write(self.subscribe("-"))
	return err
}

func (self *driverSub) open(ws io.Writer) error {
	_, err := ws.Write(self.subscribe("+"))
	return err
}

func (self *driverSub) notify(c context.Context, msg node.Selection) {
	self.stream(c, msg)
}

func (self *driverSub) subscribe(op string) []byte {
	return []byte(fmt.Sprintf(`{"op":"%s","path":"%s","group":"n2-notify"}`, op, self.sel.Path.StringNoModule()))
}
