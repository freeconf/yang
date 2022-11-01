package node

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/val"
)

func TestFieldConstraints(t *testing.T) {
	y := `module m { prefix ""; namespace ""; revision 0;
		leaf l1 {
			type string {
				pattern "a.*";
			}
		}
		leaf l2 {
			type string {
				pattern "a.*";
				pattern "b.*";
			}
		}
		leaf l3 {
			type string {
				length 5;
			}
		}
		leaf l4 {
			type string {
				length 3..5;
			}
		}
		leaf l5 {
			type string {
				length 3..5|8..10;
			}
		}
		leaf l6 {
			type string {
				pattern "a.*";
				pattern "b.*";
				length 3..5|8..10;
			}
		}
		leaf-list l7 {
			type string {
				pattern "x.*x";
				length 4|8;
			}
		}
		leaf l8 {
			type int32 {
				range 50..100;
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, y)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		MetaPath      string
		Sample        val.Value
		ExpectedCheck bool
	}{
		{
			"l1", val.String("a"), true,
		},
		{
			"l1", val.String("b"), false,
		},
		{
			"l2", val.String("b"), true,
		},
		{
			"l3", val.String("12345"), true,
		},
		{
			"l3", val.String("123"), false,
		},
		{
			"l4", val.String("12345"), true,
		},
		{
			"l4", val.String("123"), true,
		},
		{
			"l4", val.String("12"), false,
		},
		{
			"l4", val.String("123456"), false,
		},
		{
			"l5", val.String("1234"), true,
		},
		{
			"l5", val.String("123456789"), true,
		},
		{
			"l5", val.String("123456"), false,
		},
		{
			"l5", val.String("12"), false,
		},
		{
			"l5", val.String("12345678901"), false,
		},
		{
			"l6", val.String("a23456789"), true,
		},
		{
			"l6", val.String("b23"), true,
		},
		{
			"l6", val.String("b3"), false,
		},
		{
			"l6", val.String("x23"), false,
		},
		{
			"l7", val.StringList([]string{"xxxx", "x--x"}), true,
		},
		{
			"l7", val.StringList([]string{"1234", "x--x"}), false,
		},
		{
			"l7", val.StringList([]string{"x--x", "123x"}), false,
		},
		{
			"l7", val.StringList([]string{"xxxx", "x-x"}), false,
		},
		{
			"l7", val.StringList([]string{"xxxx", "x-x"}), false,
		},
		{
			"l7", val.StringList([]string{"xxxx", "x234567x"}), true,
		},
		{
			"l8", val.Int32(99), true,
		},
		{
			"l8", val.Int32(100), true,
		},
		{
			"l8", val.Int32(101), false,
		},
	}
	for _, test := range tests {
		r := FieldRequest{
			Meta: meta.Find(m, test.MetaPath).(meta.Leafable),
		}
		check := newFieldConstraints()
		ok, err := check.CheckFieldPreConstraints(&r, &ValueHandle{Val: test.Sample})
		// it should always either be:
		//  a.) true, no err
		//  b.) false, err
		if err != nil && test.ExpectedCheck {
			t.Errorf("%s, %s, %s", err, test.MetaPath, test.Sample.String())
		}
		fc.AssertEqual(t, test.ExpectedCheck, ok, test.MetaPath, test.Sample.String())
	}
}
