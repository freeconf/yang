package doc

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func Test_mdCleanDescription(t *testing.T) {
	actual := mdCleanDescription("hello\n        more text")
	fc.AssertEqual(t, "hello\nmore text", actual)
}
