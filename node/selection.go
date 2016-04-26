package node

import (
	"regexp"
	"github.com/c2g/meta"
	"fmt"
)

type Selection struct {
	parent     *Selection
	events     Events
	node       Node
	path       *Path
	insideList bool
}

func (self *Selection) Parent() *Selection {
	return self.parent
}

func (self *Selection) Events() Events {
	return self.events
}

func (self *Selection) Meta() meta.Meta {
	return self.path.meta
}

func (self *Selection) Node() Node {
	return self.node
}

func (self *Selection) InsideList() bool {
	return self.insideList
}

func (self *Selection) Fork(node Node) *Selection {
	copy := *self
	copy.events = &EventsImpl{}
	copy.node = node

	// this has the desired effect of stopping event propagation up the selection chain on
	// forked selections. If you remove this code such as inserting into json writer will cause
	// the source node to get unwatned edit events.
	copy.parent = nil

	return &copy
}

func (self *Selection) Key() []*Value {
	return self.path.key
}

func (self *Selection) String() string {
	return fmt.Sprint(self.node.String(), ":", self.path.String())
}

func Select(m meta.MetaList, node Node) *Selection {
	return &Selection{
		events: &EventsImpl{},
		path: &Path{meta: m},
		node:   node,
	}
}

func (self *Selection) SelectChild(m meta.MetaList, node Node) *Selection {
	child := &Selection{
		parent: self,
		events: self.events,
		path: &Path{parent: self.path, meta: m},
		node:   node,
	}
	return child
}

func (self *Selection) SelectListItem(node Node, key []*Value) *Selection {
	var parentPath *Path
	if self.parent != nil {
		parentPath = self.parent.path
	}
	child := &Selection{
		parent:     self.parent, // NOTE: list item's parent is list's parent, not list!
		events:     self.events,
		node:       node,
		path:		&Path{parent:parentPath, meta: self.path.meta, key: key},
		insideList: true,
	}
	return child
}

func (self *Selection) Path() *Path {
	return self.path
}

func (self *Selection) Fire(e Event) (err error) {
	target := self
	for {
		err = target.node.Event(target, e)
		if err != nil {
			return err
		}
		if e.Type.Bubbles() && ! e.state.propagationStopped {
			if target.parent != nil {
				target = target.parent
				continue
			}
		}
		break
	}
	return self.events.Fire(self.path, e)
}

func (self *Selection) On(e EventType, listener ListenFunc) *Listener {
	return self.OnPath(e, self.Path().String(), listener)
}

func (self *Selection) OnPath(e EventType, path string, handler ListenFunc) *Listener {
	listener := &Listener{event: e, path: path, handler: handler}
	self.events.AddListener(listener)
	return listener
}

func (self *Selection) OnChild(e EventType, m meta.MetaList, listener ListenFunc) *Listener {
	fullPath := self.path.String() + "/" + m.GetIdent()
	return self.OnPath(e, fullPath, listener)
}

func (self *Selection) OnRegex(e EventType, regex *regexp.Regexp, handler ListenFunc) *Listener {
	listener := &Listener{event: e, regex: regex, handler: handler}
	self.events.AddListener(listener)
	return listener
}

func (self *Selection) Peek(peekId string) interface{} {
	return self.node.Peek(self, peekId)
}

func isFwdSlash(r rune) bool {
	return r == '/'
}

func (self *Selection) IsConfig(m meta.Meta) bool {
	if hasDetails, ok := m.(meta.HasDetails); ok {
		return hasDetails.Details().Config(self.path)
	}
	return true
}

func (self *Selection) ClearAll() error {
	return self.node.Event(self, DELETE.New(self.path))
}

func (self *Selection) FindOrCreate(ident string, autoCreate bool) (*Selection, error) {
	m := meta.FindByIdent2(self.path.meta, ident)
	var err error
	var child Node
	if m != nil {
		r := ContainerRequest{
			Request:Request {
				Selection: self,
			},
			Meta: m.(meta.MetaList),
		}
		child, err = self.node.Select(r)
		if err != nil {
			return nil, err
		} else if child == nil && autoCreate {
			r.New = true
			child, err = self.node.Select(r)
			if err != nil {
				return nil, err
			}
		}
		if child != nil {
			return self.SelectChild(r.Meta, child), nil
		}
	}
	return nil, nil
}

