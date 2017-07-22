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
// Notice there are no reference to C2Stack in this file.  This means your
// code remains:
// - unit test-able
// - Not auto-generated from model files
// - free of golang source code annotations/tags.
type Car struct {
	Tire []*tire

	// Not everything has to be structs, using a map may be useful
	// in early prototyping
	Specs map[string]interface{}

	Miles   int64
	Running bool

	// When the tires were last rotated
	LastRotation int64

	// Default speed value is in yang model file and free's your code
	// from hardcoded values, even if they are only default values
	// units milliseconds/mile
	Speed int

	// Listeners are common on manageable code.  Having said that, listeners
	// remain relevant to your application.  The node.go file is responsible
	// for bridging the conversion from application to management api.
	listeners *list.List
}

type CarListener func(c *Car)

func New() *Car {
	c := &Car{
		listeners: list.New(),
		Speed:     1000,
		Specs:     make(map[string]interface{}),
	}
	c.newTires()
	return c
}

func (c *Car) newTires() {
	c.Tire = make([]*tire, 4)
	c.LastRotation = c.Miles
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

			// tip: by using time.After instead of a time.Ticker, we don't
			// have to rebuild ticker object and restart this loop if Speed
			// is dynamically changed.  Simple little tricks like this make
			// your application support live updates
			<-time.After(time.Duration(c.Speed) * time.Millisecond)
			c.Miles++

			for _, t := range c.Tire {
				previousWorn := t.Worn()

				// put random wear on a tire.  Tires in 4th position
				// receive more wear on average to make application
				// more interesting
				t.Wear -= float64(t.Pos) * (rand.Float64() / 2)
				t.checkFlat()
				if t.Flat {
					goto done
				}
				if previousWorn != t.Worn() {
					c.updateListeners()
				}
			}
		}
	done:
		c.Running = false
		c.updateListeners()
	}()
}

func (c *Car) OnUpdate(l CarListener) c2.Subscription {
	return c2.NewSubscription(c.listeners, c.listeners.PushBack(l))
}

func (c *Car) updateListeners() {
	e := c.listeners.Front()
	for e != nil {
		e.Value.(CarListener)(c)
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
	c.LastRotation = c.Miles
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
	c.LastRotation = c.Miles
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
