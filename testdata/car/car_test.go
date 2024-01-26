package car

import (
	"fmt"
	"testing"
	"time"

	"github.com/freeconf/yang/fc"
)

// Quick test of car's features using direct access to fields and methods
// again, nothing to do with FreeCONF.
func TestCar(t *testing.T) {
	c := New()
	c.pollInterval = time.Millisecond
	c.Speed = 1000

	events := make(chan UpdateEvent)
	unsub := c.OnUpdate(func(e UpdateEvent) {
		fmt.Printf("got event %s\n", e)
		events <- e
	})
	t.Log("waiting for car events...")
	c.Start()

	fc.AssertEqual(t, CarStarted, <-events)
	fc.AssertEqual(t, FlatTire, <-events)
	fc.AssertEqual(t, CarStopped, <-events)
	c.ReplaceTires()
	c.Start()

	fc.AssertEqual(t, CarStarted, <-events)
	unsub.Close()
	c.Stop()
}
