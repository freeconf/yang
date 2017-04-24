package garage

import "testing"

func TestMaintenance(t *testing.T) {
	g := NewGarage()
	g.ApplyOptions(Options{TireRotateMiles: 1000})
	tests := []struct {
		state    CarState
		expected testCar
	}{
		{
			state: CarState{
				Miles:        999,
				Running:      true,
				LastRotation: 0,
			},
			expected: testCar{
				rotate:  0,
				replace: 0,
			},
		},
		{
			state: CarState{
				Miles:        1001,
				Running:      true,
				LastRotation: 0,
			},
			expected: testCar{
				rotate:  1,
				replace: 0,
			},
		},
		{
			state: CarState{
				Miles:        999,
				Running:      false,
				LastRotation: 0,
			},
			expected: testCar{
				rotate:  0,
				replace: 1,
			},
		},
	}
	for _, test := range tests {
		var car testCar
		g.maintainCar(&car, test.state)
		if car != test.expected {
			t.Error(test.expected, car)
		}
	}
}

type testCar struct {
	rotate  int
	replace int
}

func (self *testCar) Id() string {
	return ""
}

func (self *testCar) OnChange(l CarChangeListener) (err error) {
	return
}

func (self *testCar) Close() {
}

func (self *testCar) State() (state CarState, err error) {
	return
}

func (self *testCar) ReplaceTires(state CarState) (err error) {
	self.replace++
	return
}

func (self *testCar) RotateTires(state CarState) (err error) {
	self.rotate++
	return
}
