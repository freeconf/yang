package render

import (
	"testing"

	"github.com/c2stack/c2g/c2"
)

func Test_mdCleanDescription(t *testing.T) {
	actual := mdCleanDescription("hello\n        more text")
	c2.AssertEqual(t, "hello\nmore text", actual)
}
