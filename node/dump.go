package node

import (
	"fmt"
	"io"
	"os"
	"github.com/c2stack/c2g/meta"
)

const Padding = "                                                                                       "

type Dumper struct {
	Out io.Writer
}

func DevNull() Node {
	n := &MyNode{}
	n.OnSelect = func(r ContainerRequest) (Node, error) {
		if r.New {
			return n, nil
		}
		return nil, nil
	}
	n.OnNext = func(r ListRequest) (Node, []*Value, error) {
		if r.New {
			return n, nil, nil
		}
		return nil, nil, nil
	}
	n.OnField = func(FieldRequest, *ValueHandle) error {
		return nil
	}
	return n
}

func Dumpf(n Node, fname string) Node {
	f, ferr := os.Create(fname)
	if ferr != nil {
		panic(ferr)
	}
	return Dumper{
		Out: f,
	}.Node(0, n)
}

func (self Dumper) Close() {
	if closer, ok := self.Out.(io.ReadCloser); ok {
		closer.Close()
	}
}

func Dump(n Node, out io.Writer) Node {
	return Dumper{
		Out: out,
	}.Node(0, n)
}

func (self Dumper) write(s string, args ...interface{}) {
	self.Out.Write([]byte(fmt.Sprintf(s, args...)))
}

func (self Dumper) eol() {
	self.Out.Write([]byte("\n"))
}

func (self Dumper) check(e error) {
	if e != nil {
		self.write(",! ! ! err=%s ! ! !", e.Error())
	}
}

func (self Dumper) Node(level int, target Node) Node {
	n := &MyNode{}
	n.OnChoose = func(sel Selection, choice *meta.Choice) (choosen *meta.ChoiceCase, err error) {
		self.write("%schoose %s=", Padding[:level], choice.GetIdent())
		choosen, err = target.Choose(sel, choice)
		if choosen != nil {
			self.write(choosen.GetIdent())
		} else {
			self.write("nil")
		}
		self.check(err)
		self.eol()
		return choosen, err
	}
	n.OnSelect = func(r ContainerRequest) (child Node, err error) {
		self.write("%s{%s", Padding[:level], r.Meta.GetIdent())
		if r.New {
			self.write(", new")
		}
		child, err = target.Select(r)
		if child != nil {
			self.write(", found")
		} else {
			self.write(", !found")
		}
		self.check(err)
		self.eol()
		if child == nil {
			return nil, err
		}
		return self.Node(level + 1, child), nil
	}
	n.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if r.Write {
			self.write("%s->%s=%s(", Padding[:level], r.Meta.GetIdent(), hnd.Val.Type.Ident)
			err = target.Field(r, hnd)
			self.write("%v)", hnd.Val.String())
			self.check(err)
			self.eol()
		} else {
			self.write("%s<-%s=", Padding[:level], r.Meta.GetIdent())
			err = target.Field(r, hnd)
			if hnd.Val != nil {
				self.write("%s(%v)", hnd.Val.Type.Ident, hnd.Val.String())
			} else {
				self.write("nil")
			}
			self.check(err)
			self.eol()
		}
		return
	}
	n.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
		self.write("%s[%s, row=%d", Padding[:level], r.Meta.GetIdent(), r.Row)
		if r.New {
			self.write(", new")
		}
		if r.First {
			self.write(", first")
		}
		next, key, err = target.Next(r)
		if next != nil {
			self.write(", found")
		} else {
			self.write(", !found")
		}
		self.check(err)
		self.eol()
		if next == nil {
			return nil, key, err
		}
		return self.Node(level, next), key, err
	}
	n.OnEvent = func(sel Selection, e Event) (err error) {
		self.write("%s@%s", Padding[:level], e.Type.String())
		err = target.Event(sel, e)
		self.check(err)
		self.eol()
		return err
	}
	return n
}
