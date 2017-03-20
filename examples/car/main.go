package main

import (
	"container/list"
	"math/rand"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our Car micro-service application with C2Stack for
// RESTful based management
//
// To run:
//    cd ./src/vendor/github.com/c2stack/c2g/examples/car
//    go run ./main.go
//
// Then open web browser to
//   http://localhost:8080/
//
func main() {

	// Your application
	car := newCar()

	// Where to looks for yang files, this tells library to use cwd
	ypath := &meta.FileStreamSource{Root: "."}

	// Every management has a "device" container for
	device := conf.NewLocalDeviceWithUi(ypath, ypath)

	// Browser is the management handle to your entire application
	// see how we're connecting 3 things here
	//  1.) model - schema map for your application, nothing get's in/out with adhereing to model
	//  2.) car - this can be multiple objects, or often just one but that's up to you and how
	//      you designed your application.
	//  3.) carNode - your bridge between your model and your application
	device.Add("car", carNode(car))

	// in our car app, we start off by running start.
	car.start()

	// RESTful management.  Only required for RESTful based management.
	restconf.NewManagement(device, ":8080")

	select {}
}

//////////////////////////
// C A R
// Your application, no reference to C2Stack, unit testable as any business code
// would be.  Not auto-generated code and no golang tags.

type car struct {
	Tire      []*tire
	Engine    *engine
	Miles     int64
	Running   bool
	listeners *list.List
}

type carListener func(c *car)

func newCar() *car {
	c := &car{
		listeners: list.New(),
	}
	c.newTires()
	return c
}

func (c *car) newTires() {
	c.Tire = make([]*tire, 4)
	for pos := 0; pos < len(c.Tire); pos++ {
		c.Tire[pos] = &tire{
			Pos:  pos,
			Wear: 100,
		}
	}
}

func (c *car) start() {
	if c.Running {
		return
	}
	go func() {
		c.Running = true
		c.updateListeners()
		for {
			<-time.After(1000 * time.Millisecond)
			for _, t := range c.Tire {
				previousWorn := t.Worn()
				t.Wear -= float64(t.Pos) * (rand.Float64() / 2)
				t.checkFlat()
				if t.Flat {
					c.Running = false
					c.updateListeners()
					return
				}
				if previousWorn != t.Worn() {
					c.updateListeners()
				}
			}
			c.Miles++
		}
	}()
}

func (c *car) onUpdate(l carListener) c2.Subscription {
	return c2.NewSubscription(c.listeners, c.listeners.PushBack(l))
}

func (c *car) updateListeners() {
	e := c.listeners.Front()
	for e != nil {
		e.Value.(carListener)(c)
		e = e.Next()
	}
}

func (c *car) hasWornTire() bool {
	for _, t := range c.Tire {
		if t.Worn() {
			return true
		}
	}
	return false
}

func (c *car) replaceTires() {
	for _, t := range c.Tire {
		t.replace()
	}
	c.start()
}

func (c *car) rotateTires() {
	x := c.Tire[0]
	c.Tire[0] = c.Tire[1]
	c.Tire[1] = c.Tire[2]
	c.Tire[2] = c.Tire[3]
	c.Tire[3] = x
	for i, t := range c.Tire {
		t.Pos = i
	}
}

// T I R E
type tire struct {
	Pos  int
	Size string
	Flat bool
	Wear float64
}

type tireStatus int

const (
	tireFlat tireStatus = iota + 1
	tireLow
	tireWorn
)

func (t *tire) replace() {
	t.Wear = 100
	t.Flat = false
}

type tireListener func(t *tire)

type tireListenerRecord struct {
	previous tire
	l        tireListener
}

func (t *tire) checkFlat() bool {
	if !t.Flat {
		// really need gausian distribution
		t.Flat = (t.Wear - (rand.Float64() * 10)) < 0
		return t.Flat
	}
	return false
}

func (t *tire) Worn() bool {
	return t.Wear < 20
}

// E N G I N E
type engine struct {
	Specs map[string]interface{}
}

/////////////////////////
// C A R  N O D E
//  Bridge from model to car app.

// carNode is root handler from car.yang
//    module car { ... }
func carNode(c *car) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(c),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "tire":
				return tiresNode(c.Tire), nil
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		OnAction: func(p node.Node, r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "rotateTires":
				c.rotateTires()
			case "replaceTires":
				c.replaceTires()
			}
			return nil, nil
		},
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "update":
				sub := c.onUpdate(func(*car) {
					r.Send(r.Context, carNode(c))
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

// tiresNode handles list of tires.
//     list tire { ... }
func tiresNode(tires []*tire) node.Node {
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			var t *tire
			key := r.Key
			var pos int
			if key != nil {
				pos = key[0].Int
				if pos >= len(tires) {
					return nil, nil, nil
				}
			}
			if key != nil {
				t = tires[pos]
			} else {
				if r.Row < len(tires) {
					t = tires[r.Row]
					key = node.SetValues(r.Meta.KeyMeta(), r.Row)
				}
			}
			if t != nil {
				return tireNode(t), key, nil
			}
			return nil, nil, nil
		},
	}
}

// tireNode handles each tire node.  Everything *inside* list tire { ...}
func tireNode(t *tire) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(t),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "worn":
				hnd.Val = &node.Value{Bool: t.Worn()}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}
