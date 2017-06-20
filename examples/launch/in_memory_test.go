package launch

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

var testYPath = meta.PathStreamSource(".:../car:../garage:../../yang")

func Test_LocalFactory(t *testing.T) {
	f := NewInMemory(testYPath)
	app := &App{
		Id:      "c1",
		Type:    "car",
		Startup: map[string]interface{}{},
	}
	if err := f.Launch(app); err != nil {
		t.Error(err)
	}
	if mapWrongSize := c2.CheckEqual(1, f.Map.Len()); mapWrongSize != nil {
		t.Error(mapWrongSize)
	}
}
