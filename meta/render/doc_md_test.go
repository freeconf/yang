package render

import (
	"testing"

	"github.com/freeconf/gconf/c2"
)

func Test_mdCleanDescription(t *testing.T) {
	actual := mdCleanDescription("hello\n        more text")
	c2.AssertEqual(t, "hello\nmore text", actual)
}
