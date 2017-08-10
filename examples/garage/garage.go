package garage

import (
	"container/list"
	"time"

	"github.com/c2stack/c2g/c2"
)

type Garage struct {
	options          Options
	cars             *list.List
	listeners        *list.List
	errs             chan error
	running          bool
	TireReplacements int
	TireRotations    int
}

func NewGarage() *Garage {
	return &Garage{
		cars:      list.New(),
		listeners: list.New(),
	}
}

type Options struct {
	TireRotateMiles int64
	PollTimeMs      int
}

type CarChangeListener func(c Car, state CarState)

type Car interface {
	Id() string
	OnChange(l CarChangeListener) error
	Close()
	State() (CarState, error)
	ReplaceTires(state CarState) error
	RotateTires(state CarState) error
}

type CarState struct {
	Miles        int64
	LastRotation int64
	Running      bool
	Speed        int
	Tire         map[int]interface{}
}

type WorkType int

const (
	workRotateTires WorkType = iota
	workChangeTires
)

func (self WorkType) String() string {
	switch self {
	case workRotateTires:
		return "workRotateTires"
	case workChangeTires:
		return "workChangeTires"
	}
	return ""
}

func (self *Garage) Options() Options {
	return self.options
}

func (self *Garage) ApplyOptions(options Options) error {
	if self.options == options {
		return nil
	}
	self.options = options
	if self.running {
		go self.start()
	}
	return nil
}

type CarHandle *list.Element

func (self *Garage) CarCount() int {
	return self.cars.Len()
}

func (self *Garage) CarsServiced() int {
	return self.TireReplacements + self.TireRotations
}

func (self *Garage) AddCar(c Car) CarHandle {
	c.OnChange(self.maintainCar)
	return CarHandle(self.cars.PushBack(c))
}

func (self *Garage) RemoveCar(h CarHandle) {
	self.cars.Remove(h)
	h.Value.(Car).Close()
}

func (self *Garage) start() {
	self.running = true
	for {
		if self.options.PollTimeMs > 0 {
			<-time.After(time.Duration(self.options.PollTimeMs) * time.Millisecond)
		}
		p := self.cars.Front()
		for p != nil {
			c := p.Value.(Car)
			state, err := c.State()
			if err != nil {
				self.errs <- err
			} else {
				self.maintainCar(c, state)
			}
			p = p.Next()
		}
	}
}

func (self *Garage) maintainCar(c Car, state CarState) {
	if !state.Running {
		c.ReplaceTires(state)
		self.TireReplacements++
		self.updateListeners(c, workChangeTires)
	} else if (state.Miles - state.LastRotation) > self.options.TireRotateMiles {
		c.RotateTires(state)
		self.TireRotations++
		self.updateListeners(c, workRotateTires)
	}
}

func (self *Garage) OnUpdate(l MaintenanceListener) c2.Subscription {
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

type MaintenanceListener func(c Car, work WorkType)

func (self *Garage) updateListeners(c Car, work WorkType) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(MaintenanceListener)(c, work)
		p = p.Next()
	}
}
