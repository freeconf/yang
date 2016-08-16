package node

import (
	"fmt"
	"github.com/dhubler/c2g/meta"
)

type Selection struct {
	browser    *Browser
	parent     *Selection
	node       Node
	path       *Path
	insideList bool
}

func (self *Selection) Browser() *Browser {
	return self.browser
}

func (self *Selection) Parent() *Selection {
	return self.parent
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
	copy.browser = NewBrowser2(self.path.meta.(meta.MetaList), node)
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

func (self *Selection) SelectChild(m meta.MetaList, node Node) *Selection {
	child := &Selection{
		browser: self.browser,
		parent:  self,
		path:    &Path{parent: self.path, meta: m},
		node:    node,
	}
	return child
}

func (self *Selection) Selector() Selector {
	return Selector{
		Selection:   self,
		constraints: &Constraints{},
		handler:     &ConstraintHandler{},
	}
}

func (self *Selection) SelectListItem(node Node, key []*Value) *Selection {
	var parentPath *Path
	if self.parent != nil {
		parentPath = self.parent.path
	}
	child := &Selection{
		browser:    self.browser,
		parent:     self,
		node:       node,
		// NOTE: Path.parent is lists parentPath, not self.path
		path:       &Path{parent: parentPath, meta: self.path.meta, key: key},
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
		if e.Type.Bubbles() && !e.state.propagationStopped {
			if target.parent != nil {
				target = target.parent
				continue
			}
		}
		break
	}
	return self.browser.Triggers.Fire(self.Path().String(), e)
}

func (self *Selection) Peek() interface{} {
	return self.node.Peek(self)
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

func (self *Selection) FindOrCreate(ident string, autoCreate bool) (*Selection, error) {
	m := meta.FindByIdent2(self.path.meta, ident)
	var err error
	var child Node
	if m != nil {
		r := ContainerRequest{
			Request: Request{
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
