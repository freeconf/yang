package car

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ////////////////////////
// C A R - example application
//
// This has nothing to do with FreeCONF, just an example application written in Go.
// that models a running car that can get flat tires when tires are worn.
type Car struct {
	Tire         []*Tire
	Miles        float64
	Speed        int
	OilLevel     float64
	Running      bool
	LastRotation int64 // When the tires were last rotated

	// Listeners are common on manageable code.  Having said that, listeners
	// remain relevant to your application.  The manage.go file is responsible
	// for bridging the conversion from application to management api.
	listeners *list.List

	pollInterval time.Duration // adjust in unit testing to speed up tests
}

// CarListener for receiving car update events
type CarListener func(UpdateEvent)

// car event types
type UpdateEvent int

const (
	CarStarted UpdateEvent = iota + 1
	CarStopped
	FlatTire
	BadOilLevel
)

func (e UpdateEvent) String() string {
	strs := []string{
		"unknown",
		"carStarted",
		"carStopped",
		"flatTire",
		"badOilLevel",
	}
	if int(e) < len(strs) {
		return strs[e]
	}
	return "invalid"
}

func New() *Car {
	c := &Car{
		listeners:    list.New(),
		OilLevel:     10,
		pollInterval: time.Second,
	}
	c.NewTires()
	return c
}

func (c *Car) endureMileage(miles float64) {
	c.Miles += miles
	c.OilLevel -= (miles * 0.0001)
	if !c.checkOil() {
		c.updateListeners(BadOilLevel)
		c.Running = false
		return
	}
	for _, t := range c.Tire {
		// Wear down [0.0 - 0.5] of each tire proportionally to the tire position
		t.Wear -= miles * float64(t.Pos) * rand.Float64() * 5
		t.checkIfFlat()
		t.checkForWear()
		if t.Flat {
			c.Running = false
			c.updateListeners(FlatTire)
		}
	}
}

const oilOptimalLevel = 10.0
const oilDamageMargin = 2.0

func (c *Car) checkOil() bool {
	return math.Abs(oilOptimalLevel-c.OilLevel) <= oilDamageMargin
}

// Stop will take up to poll_interval seconds to come to a stop
func (c *Car) Stop() {
	c.Running = false
}

func (c *Car) Start() {
	if c.Running {
		return
	}
	ticker := time.NewTicker(c.pollInterval)
	c.Running = true
	go func() {
		c.updateListeners(CarStarted)
		defer func() {
			c.Running = false
			c.updateListeners(CarStopped)
		}()
		for range ticker.C {
			miles := float64(c.Speed) * (float64(time.Second) / float64(time.Hour))
			c.endureMileage(miles)
			if !c.Running {
				return
			}
		}
	}()
}

// OnUpdate to listen for car events like start, stop and flat tire
func (c *Car) OnUpdate(l CarListener) Subscription {
	return NewSubscription(c.listeners, c.listeners.PushBack(l))
}

func (c *Car) NewTires() {
	c.Tire = make([]*Tire, 4)
	c.LastRotation = int64(c.Miles)
	for pos := 0; pos < len(c.Tire); pos++ {
		c.Tire[pos] = &Tire{
			Pos:  pos,
			Wear: 100,
			Size: "H15",
		}
	}
}

// ChangeOil is just an example for a function that accepts input and returns output that is
// mapped to an "rpc" in the YANG file.
func (c *Car) AddOil(drainFirst bool, amount float64) (float64, error) {
	if drainFirst {
		if c.Running {
			return 0, errors.New("cannot change oil while car is running")
		}
		c.OilLevel = 0
	}
	c.OilLevel = c.OilLevel + amount
	if !c.checkOil() {
		return 0, fmt.Errorf("invalid oil change level %.2f liters", c.OilLevel)
	}
	return c.OilLevel, nil
}

func (c *Car) ReplaceTires() {
	for _, t := range c.Tire {
		t.Replace()
	}
	c.LastRotation = int64(c.Miles)
}

func (c *Car) RotateTires() {
	x := c.Tire[0]
	c.Tire[0] = c.Tire[1]
	c.Tire[1] = c.Tire[2]
	c.Tire[2] = c.Tire[3]
	c.Tire[3] = x
	for i, t := range c.Tire {
		t.Pos = i
	}
	c.LastRotation = int64(c.Miles)
}

func (c *Car) updateListeners(e UpdateEvent) {
	for i := c.listeners.Front(); i != nil; i = i.Next() {
		i.Value.(CarListener)(e)
	}
}

// T I R E
type Tire struct {
	Pos  int
	Size string
	Flat bool
	Wear float64
	Worn bool
}

func (t *Tire) Replace() {
	t.Wear = 100
	t.Flat = false
	t.Worn = false
}

func (t *Tire) checkIfFlat() {
	if !t.Flat {
		// emulate that the more wear a tire has, the more likely it will
		// get a flat, but there is always a chance.
		t.Flat = (t.Wear - (rand.Float64() * 10)) < 0
	}
}

func (t *Tire) checkForWear() bool {
	return t.Wear < 20
}

///////////////////////
// U T I L

// Subscription is handle into a list.List that when closed
// will automatically remove item from list.  Useful for maintaining
// a set of listeners that can easily remove themselves.
type Subscription interface {
	Close() error
}

// NewSubscription is used by subscription managers to give a token
// to caller the can close to unsubscribe to events
func NewSubscription(l *list.List, e *list.Element) Subscription {
	return &listSubscription{l, e}
}

type listSubscription struct {
	l *list.List
	e *list.Element
}

// Close will unsubscribe to events.
func (sub *listSubscription) Close() error {
	sub.l.Remove(sub.e)
	return nil
}
