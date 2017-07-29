package node

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
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
				config "false";
			}
		}
		container y {
			config "false";
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
	x, _ := meta.FindByIdent2(m, "x")
	y, _ := meta.FindByIdent2(m, "y")
	mSel := NewBrowser(m, nil).Root()
	containerTests := []struct {
		sel      Selection
		m        meta.MetaList
		expected bool
	}{
		{
			mSel,
			x.(meta.MetaList),
			true,
		},
		{
			mSel,
			y.(meta.MetaList),
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
	xa, _ := meta.FindByIdent2(x, "a")
	ySel := Selection{
		Parent: &mSel,
		Path:   &Path{parent: mSel.Path, meta: y},
	}
	ya, _ := meta.FindByIdent2(y, "a")
	z, _ := meta.FindByIdent2(m, "z")
	zSel := Selection{
		Parent: &mSel,
		Path:   &Path{parent: mSel.Path, meta: z},
	}
	za, _ := meta.FindByIdent2(z, "a")
	fieldTests := []struct {
		sel      Selection
		m        meta.HasDataType
		expected bool
	}{
		{
			xSel,
			xa.(meta.HasDataType),
			false,
		},
		{
			ySel,
			ya.(meta.HasDataType),
			false,
		},
		{
			zSel,
			za.(meta.HasDataType),
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
