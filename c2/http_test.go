package c2

import "testing"

func Test_AppendUrlSegment(t *testing.T) {
	tests := [][]string{
		{
			"a", "b", "a/b",
		},
		{
			"a/", "b", "a/b",
		},
		{
			"a/", "/b", "a/b",
		},
		{
			"a", "/b", "a/b",
		},
		{
			"", "", "",
		},
		{
			"a/", "", "a/",
		},
		{
			"", "/b", "/b",
		},
	}
	for _, test := range tests {
		actual := AppendUrlSegment(test[0], test[1])
		if err := CheckEqual(test[2], actual); err != nil {
			t.Error(err)
		}
	}
}
