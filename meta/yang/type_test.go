package yang

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
)

func TestTypeResolve(t *testing.T) {
	yang := `
module ff {
	namespace "ns";

	description "mod";

	revision 99-99-9999 {
	  description "bingo";
	}

	leaf x {
		type int32;
	}
	typedef q {
		type string;
	}
	list y {
		key "id";
		leaf id {
			type string;
		}
	}
	container z {
	  description "z1";
	  leaf z1 {
	    type leafref {
	    	path "../x";
	    }
	  }
	  leaf z2 {
	    type leafref {
	    	path "../y/id";
	    }
	  }
		leaf z3 {
			type q;
		}
	}
}
`
	m, err := LoadModuleCustomImport(yang, nil)
	if err != nil {
		t.Fatal(err)
	}
	z1, err := meta.FindByPath(m, "z/z1")
	if err != nil {
		t.Error(err)
	} else if z1 == nil {
		t.Errorf("No z1")
	}
	dt := z1.(meta.HasDataType).GetDataType()
	i, _ := dt.Info()
	if i.Format != val.FmtInt32 {
		t.Errorf("actual type %s", i.Format)
	}
	z3, err := meta.FindByPath(m, "z/z3")
	if err != nil {
		t.Error(err)
	}
	dt = z3.(meta.HasDataType).GetDataType()
	i, _ = dt.Info()
	if i.Format != val.FmtString {
		t.Errorf("actual type %s", i.Format)
	}
}
