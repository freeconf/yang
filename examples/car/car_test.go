package car

import "testing"

func Test_Car(t *testing.T) {
	c := New()
	c.Speed = 1
	update := make(chan bool)
	c.OnUpdate(func(c *Car) {
		update <- true
	})
	c.Start()
	<-update
	if !c.Running {
		t.Error("not running when starting")
	}
}
