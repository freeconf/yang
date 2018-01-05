package val

import (
	"testing"

	"github.com/freeconf/gconf/c2"
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
		c2.AssertEqual(t, test.Expected, actual)
	}
}

func TestSingle(t *testing.T) {
	c2.AssertEqual(t, FmtString, FmtStringList.Single())
}
