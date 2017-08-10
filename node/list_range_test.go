package node_test

import (
	"testing"

	"github.com/c2stack/c2g/node"
)

func TestNewListRange(t *testing.T) {
	tests := []struct {
		expression string
		start      int64
		end        int64
	}{
		{
			"100-200",
			100,
			200,
		},
		{
			"100-",
			100,
			-1,
		},
		{
			"100",
			100,
			-1,
		},
	}
	for i, test := range tests {
		lr, err := node.NewListRange("aaa(bbb;ccc)!" + test.expression)
		if err != nil {
			t.Error(i, err)
		}
		if lr.StartRow != test.start {
			t.Error(i, lr.StartRow)
		}
		if lr.EndRow != test.end {
			t.Error(i, lr.EndRow)
		}
	}
}
