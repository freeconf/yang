package restconf

import (
	"strings"
	"testing"
)

func TestSseDecode(t *testing.T) {
	tests := []struct {
		payload  string
		expected []string
	}{
		{
			payload: `
data: x

data: hello
data:  world`,
			expected: []string{"x", "hello world"},
		},
		{
			payload:  "data: z",
			expected: []string{"z"},
		},
		{
			payload: `
ignored: z
data: foo
`,
			expected: []string{"foo"},
		},
	}
	for _, test := range tests {
		events := decodeSse(strings.NewReader(test.payload))
		for _, expected := range test.expected {
			actual := <-events
			if expected != string(actual) {
				t.Errorf("expected '%s' got '%s'", expected, actual)
			}
		}
	}
}
