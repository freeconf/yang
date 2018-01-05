package nodes

import (
	"fmt"
	"io"
	"os"

	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/val"
)

const padding = "                                                                                       "

type dump struct {
	Out io.Writer
}

// DumpBin sends useful information to string when written to
func DumpBin(out io.Writer) node.Node {
	return Dump(Null(), out)
}

// Dumpf will send useful information to file while delegating reads/writes to the given
// node
func Dumpf(n node.Node, fname string) node.Node {
	f, ferr := os.Create(fname)
	if ferr != nil {
		panic(ferr)
	}
	return dump{
		Out: f,
	}.Node(0, n)
}

// Dump will send useful information to writer while delegating reads/writes to the given
// node
func Dump(n node.Node, out io.Writer) node.Node {
	return dump{
		Out: out,
	}.Node(0, n)
}

func (self dump) Close() {
	if closer, ok := self.Out.(io.ReadCloser); ok {
		closer.Close()
	}
}

func (self dump) write(s string, args ...interface{}) {
	self.Out.Write([]byte(fmt.Sprintf(s, args...)))
}

func (self dump) eol() {
	self.Out.Write([]byte("\n"))
}

func (self dump) check(e error) {
	if e != nil {
		self.write(",! ! ! err=%s ! ! !", e.Error())
	}
}

func (self dump) Node(level int, target node.Node) node.Node {
	n := &Basic{}
	n.OnChoose = func(sel node.Selection, choice *meta.Choice) (choosen *meta.ChoiceCase, err error) {
		self.write("%schoose %s=", padding[:level], choice.Ident())
		choosen, err = target.Choose(sel, choice)
		if choosen != nil {
			self.write(choosen.Ident())
		} else {
			self.write("nil")
		}
		self.check(err)
		self.eol()
		return choosen, err
	}
	n.OnChild = func(r node.ChildRequest) (child node.Node, err error) {
		self.write("%s{%s", padding[:level], r.Meta.Ident())
		if r.New {
			self.write(", new")
		}
		child, err = target.Child(r)
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
		return self.Node(level+1, child), nil
	}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			self.write("%s->%s=%s(", padding[:level], r.Meta.Ident(), hnd.Val.Format())
			err = target.Field(r, hnd)
			self.write("%v)", hnd.Val.String())
			self.check(err)
			self.eol()
		} else {
			self.write("%s<-%s=", padding[:level], r.Meta.Ident())
			err = target.Field(r, hnd)
			if hnd.Val != nil {
				self.write("%s(%v)", hnd.Val.Format(), hnd.Val.String())
			} else {
				self.write("nil")
			}
			self.check(err)
			self.eol()
		}
		return
	}
	n.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		self.write("%s[%s, row=%d", padding[:level], r.Meta.Ident(), r.Row)
		if r.New {
			self.write(", new")
		}
		if r.First {
			self.write(", first")
		}
		next, key, err = target.Next(r)
		if next != nil {
			self.write(", found key=%v", key)
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
	onNodeRequest := func(r node.NodeRequest, entry string) {
		self.write("%s%s, new=%v, src=%s", padding[:level], entry, r.New, r.Source.Path.String())
		self.eol()
	}
	n.OnBeginEdit = func(r node.NodeRequest) (err error) {
		onNodeRequest(r, "BeginEdit")
		return
	}
	n.OnEndEdit = func(r node.NodeRequest) (err error) {
		onNodeRequest(r, "EndEdit")
		return
	}
	n.OnDelete = func(r node.NodeRequest) (err error) {
		onNodeRequest(r, "Delete")
		return
	}
	return n
}
