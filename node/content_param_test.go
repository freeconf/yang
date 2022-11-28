package node

import (
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/parser"
)

func TestContentConstraintParse(t *testing.T) {
	if c, _ := NewContentConstraint(nil, "config"); c != ContentConfig {
		t.Fail()
	}
}

func TestContentConstraintCheck(t *testing.T) {
	mstr := `
	module m {
		namespace "";
		prefix "";
		revision 0;
		container x {
			leaf a {
				type string;
				config false;
			}
		}
		container y {
			config false;
			leaf a {
				type string;
			}
		}
		container z {
			leaf a {
				type string;
			}
		}
	}
	`
	m, _ := parser.LoadModuleFromString(nil, mstr)
	x := meta.Find(m, "x").(meta.HasDataDefinitions)
	y := meta.Find(m, "y").(meta.HasDataDefinitions)
	mSel := NewBrowser(m, ErrorNode{}).Root()
	containerTests := []struct {
		sel      Selection
		m        meta.HasDataDefinitions
		expected bool
	}{
		{
			mSel,
			x,
			true,
		},
		{
			mSel,
			y,
			false,
		},
	}
	for i, test := range containerTests {
		r := &ChildRequest{
			Request: Request{
				Selection: test.sel,
			},
			Meta: test.m,
		}
		pass, _ := ContentConfig.CheckContainerPreConstraints(r)
		if pass != test.expected {
			t.Errorf("container test %d failed", i)
		}
	}

	xSel := Selection{
		Parent: &mSel,
		Path:   &Path{Parent: mSel.Path, Meta: x},
	}
	xa := meta.Find(x, "a")
	ySel := Selection{
		Parent: &mSel,
		Path:   &Path{Parent: mSel.Path, Meta: y},
	}
	ya := meta.Find(y, "a")
	z := meta.Find(m, "z").(meta.HasDataDefinitions)
	zSel := Selection{
		Parent: &mSel,
		Path:   &Path{Parent: mSel.Path, Meta: z},
	}
	za := meta.Find(z, "a")
	fieldTests := []struct {
		sel      Selection
		m        meta.Leafable
		expected bool
	}{
		{
			xSel,
			xa.(meta.Leafable),
			false,
		},
		{
			ySel,
			ya.(meta.Leafable),
			false,
		},
		{
			zSel,
			za.(meta.Leafable),
			true,
		},
	}
	for i, test := range fieldTests {
		r := &FieldRequest{
			Request: Request{
				Selection: test.sel,
			},
			Meta: test.m,
		}
		pass, _ := ContentConfig.CheckFieldPreConstraints(r, nil)
		if pass != test.expected {
			t.Errorf("field test %d failed", i)
		}
	}
}
