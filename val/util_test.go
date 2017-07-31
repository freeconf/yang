package val

import (
	"testing"

	"github.com/c2stack/c2g/c2"
)

func TestReduce(t *testing.T) {
	appender := func(index int, v Value, data interface{}) interface{} {
		s := data.(string)
		if index > 0 {
			s += ","
		}
		s += v.String()
		return s
	}
	tests := []struct {
		In       Value
		Expected string
	}{
		{
			In:       StringList([]string{"a", "b", "c"}),
			Expected: "a,b,c",
		},
		{
			In:       String("hey"),
			Expected: "hey",
		},
	}
	for _, test := range tests {
		actual := Reduce(test.In, "", appender)
		if err := c2.CheckEqual(test.Expected, actual); err != nil {
			t.Error(err)
		}
	}
}
