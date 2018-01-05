package node_test

import "testing"
import "github.com/freeconf/gconf/meta/yang"
import "github.com/freeconf/gconf/node"
import "github.com/freeconf/gconf/nodes"
import "github.com/freeconf/gconf/c2"
import "fmt"

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
					out: `{"y":{"z":99}}`,
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
					out: `{"y":100,"z":99}`,
				},
				{
					in:  `{"y":100,"z":9}`,
					out: `{"z":9}`,
				},
			},
		},
	}
	for _, test := range tests {
		mstr := fmt.Sprintf(`module x {revision 0;%s}`, test.y)
		m := yang.RequireModuleFromString(nil, mstr)
		for _, d := range test.data {
			b := node.NewBrowser(m, nodes.ReadJSON(d.in))
			actual, err := nodes.WriteJSON(b.Root())
			if err != nil {
				t.Error(err)
			}
			c2.AssertEqual(t, d.out, actual)
		}
	}
}
