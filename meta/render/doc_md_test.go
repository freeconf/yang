package render

import (
	"testing"

	"github.com/c2stack/c2g/c2"
)

func Test_mdCleanDescription(t *testing.T) {
	actual := mdCleanDescription("hello\n        more text")
	expected := "hello\nmore text"
	if err := c2.CheckEqual(expected, actual); err != nil {
		t.Error(err)
	}
}
