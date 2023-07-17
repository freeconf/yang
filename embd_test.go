package yang

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestYang(t *testing.T) {
	files, err := internal.ReadDir("yang")
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, 2, len(files))
}
