package orchestrator

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
)

var testYPath = meta.PathStreamSource(".:../car:../garage:../../yang")

func Test_LocalFactory(t *testing.T) {
	m := device.NewMap()
	f := &LocalFactory{
		Ypath: testYPath,
		Map:   m,
	}
	app := &App{
		Id:      "c1",
		Type:    "car",
		Startup: map[string]interface{}{},
	}
	if err := f.NewApp(app); err != nil {
		t.Error(err)
	}
	if mapWrongSize := c2.CheckEqual(1, m.Len()); mapWrongSize != nil {
		t.Error(mapWrongSize)
	}
}
