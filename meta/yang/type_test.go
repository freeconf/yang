package yang
import (
	"testing"
	"github.com/c2g/meta"
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
	}
}
`
	m, err := LoadModuleCustomImport(yang, nil)
	if err != nil {
		t.Fatal(err)
	}
	z1 := meta.FindByPath(m, "z/z1")
	if z1 == nil {
		t.Errorf("No z1")
	}
	dt := z1.(meta.HasDataType).GetDataType()
	if dt.Format() != meta.FMT_INT32 {
		t.Errorf("actual type %d", dt.Format())
	}
}
