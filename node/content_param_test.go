package node

import (
	"testing"

	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/meta/yang"
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
	m, _ := yang.LoadModuleFromString(nil, mstr)
	x := meta.Find(m, "x").(meta.HasDataDefs)
	y := meta.Find(m, "y").(meta.HasDataDefs)
	mSel := NewBrowser(m, nil).Root()
	containerTests := []struct {
		sel      Selection
		m        meta.HasDataDefs
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
		Path:   &Path{parent: mSel.Path, meta: x},
	}
	xa := meta.Find(x, "a")
	ySel := Selection{
		Parent: &mSel,
		Path:   &Path{parent: mSel.Path, meta: y},
	}
	ya := meta.Find(y, "a")
	z := meta.Find(m, "z").(meta.HasDataDefs)
	zSel := Selection{
		Parent: &mSel,
		Path:   &Path{parent: mSel.Path, meta: z},
	}
	za := meta.Find(z, "a")
	fieldTests := []struct {
		sel      Selection
		m        meta.HasType
		expected bool
	}{
		{
			xSel,
			xa.(meta.HasType),
			false,
		},
		{
			ySel,
			ya.(meta.HasType),
			false,
		},
		{
			zSel,
			za.(meta.HasType),
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
