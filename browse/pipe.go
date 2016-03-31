package browse

import (
	"github.com/c2g/node"
	"github.com/c2g/c2"
)

type Pipe struct {
	messages     chan *pipeMessage
	position *pipeMessage
}

type tok int

const (
	PipeSelect tok = iota + 1
	PipeListItem
	PipeLeaf
	PipeEnd
)

func NewPipe() *Pipe {
	return &Pipe{}
}

func (self *Pipe) peek() *pipeMessage {
	if self.position == nil {
		self.position = <- self.messages
		if self.position.tok == PipeEnd {
			close(self.messages)
		}
	}
	return self.position
}

type pipeMessage struct {
	tok tok
	ident string
	val *node.Value
	key []*node.Value
	err error
}

func (self *Pipe) consume() {
	self.position = nil
}

func (self *Pipe) Close(err error) {
	defer func() {
		if r:= recover(); r != nil {
			// channel was probably already closed so log err if there was one
			if err != nil {
				c2.Err.Printf(err.Error())
			}
		}
	}()
	self.messages <- &pipeMessage{
		tok:PipeEnd,
		err: err,
	}
}

func (self *Pipe) PullPush() (node.Node, node.Node) {
	self.messages = make(chan *pipeMessage)
	pull := &node.MyNode{}
	push := &node.MyNode{}
	pull.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		if r.New {
			return nil, c2.NewErr("Not a writer")
		}
		msg := self.peek()
		if msg.tok != PipeSelect || msg.ident != r.Meta.GetIdent() {
			return nil, msg.err
		}
		defer self.consume()
		return pull, msg.err
	}
	pull.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if r.New {
			return nil, nil, c2.NewErr("Not a writer")
		}
		msg := self.peek()
		if msg.tok != PipeListItem {
			return nil, nil, msg.err
		}
		defer self.consume()
		return pull, msg.key, msg.err
	}
	pull.OnRead = func(r node.FieldRequest) (v *node.Value, err error) {
		msg := self.peek()
		if msg.tok != PipeLeaf || msg.ident != r.Meta.GetIdent() {
			return nil, msg.err
		}
		defer self.consume()
		return msg.val, msg.err
	}
	push.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		if ! r.New {
			return nil, nil
		}
		self.messages <- &pipeMessage{
			tok: PipeSelect,
			ident: r.Meta.GetIdent(),
		}
		return push, nil
	}
	push.OnWrite = func(r node.FieldRequest, v *node.Value) error {
		self.messages <- &pipeMessage{
			tok: PipeLeaf,
			val: v,
			ident: r.Meta.GetIdent(),
		}
		return nil
	}
	push.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if ! r.New {
			return nil, nil, nil
		}
		self.messages <- &pipeMessage{
			tok: PipeListItem,
			key: r.Key,
			ident: r.Meta.GetIdent(),
		}
		return push, r.Key, nil
	}
	return pull, push
}
