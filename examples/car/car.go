package car

import (
	"container/list"
	"math/rand"
	"time"

	"github.com/c2stack/c2g/c2"
)

//////////////////////////
// C A R
// Your application code.
//
// Notice there are no reference to C2Stack, unit testable as any business code
// would be.  Not auto-generated code and no golang source code annotations/tags.
//
type Car struct {
	Tire      []*tire
	Engine    *engine
	Miles     int64
	Running   bool
	Speed     int
	listeners *list.List
}

type carListener func(c *Car)

func New() *Car {
	c := &Car{
		listeners: list.New(),
		Speed:     1000,
	}
	c.newTires()
	return c
}

func (c *Car) newTires() {
	c.Tire = make([]*tire, 4)
	for pos := 0; pos < len(c.Tire); pos++ {
		c.Tire[pos] = &tire{
			Pos:  pos,
			Wear: 100,
		}
	}
}

func (c *Car) Start() {
	if c.Running {
		return
	}
	go func() {
		c.Running = true
		c.updateListeners()
		for c.Speed > 0 {
			<-time.After(time.Duration(c.Speed) * time.Millisecond)
			for _, t := range c.Tire {
				previousWorn := t.Worn()
				t.Wear -= float64(t.Pos) * (rand.Float64() / 2)
				t.checkFlat()
				if t.Flat {
					goto done
				}
				if previousWorn != t.Worn() {
					c.updateListeners()
				}
			}
			c.Miles++
		}
	done:
		c.Running = false
		c.updateListeners()
	}()
}

func (c *Car) onUpdate(l carListener) c2.Subscription {
	return c2.NewSubscription(c.listeners, c.listeners.PushBack(l))
}

func (c *Car) updateListeners() {
	e := c.listeners.Front()
	for e != nil {
		e.Value.(carListener)(c)
		e = e.Next()
	}
}

func (c *Car) hasWornTire() bool {
	for _, t := range c.Tire {
		if t.Worn() {
			return true
		}
	}
	return false
}

func (c *Car) replaceTires() {
	for _, t := range c.Tire {
		t.replace()
	}
	c.Start()
}

func (c *Car) rotateTires() {
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
