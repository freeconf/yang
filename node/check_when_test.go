package node_test

import (
	"fmt"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
)

type whenTestData struct {
	in  string
	out string
}

func TestWhen(t *testing.T) {
	tests := []struct {
		y    string
		data []whenTestData
	}{
		{
			y: `
				container y {
					when "z>10";
					leaf z {
						type int32;
					}
				}
			`,
			data: []whenTestData{
				{
					in:  `{"y":{"z":99}}`,
					out: `{"x:y":{"z":99}}`,
				},
				{
					in:  `{"y":{"z":9}}`,
					out: `{}`,
				},
			},
		},
		{
			y: `
				leaf y {
					when "z>10";
					type int32;
				}
				leaf z {
					type int32;
				}
			`,
			data: []whenTestData{
				{
					in:  `{"y":100,"z":99}`,
					out: `{"x:y":100,"x:z":99}`,
				},
				{
					in:  `{"y":100,"z":9}`,
					out: `{"x:z":9}`,
				},
			},
		},
	}
	for _, test := range tests {
		mstr := fmt.Sprintf(`module x {revision 0;%s}`, test.y)
		m, err := parser.LoadModuleFromString(nil, mstr)
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range test.data {
			b := node.NewBrowser(m, nodeutil.ReadJSON(d.in))
			actual, err := nodeutil.WriteJSON(b.Root())
			if err != nil {
				t.Error(err)
			}
			fc.AssertEqual(t, d.out, actual)
		}
	}
}
