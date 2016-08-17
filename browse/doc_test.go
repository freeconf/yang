package browse

import (
	"testing"

	"github.com/dhubler/c2g/meta/yang"
)

func TestDocBuild(t *testing.T) {
	mstr := `module x {
	namespace "";
	prefix "";
	revision 0;
	container x {
		leaf z {
			type string;
		}
	}
}`
	m, err := yang.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	doc := &Doc{}
	if doc.Build(m, "html"); doc.LastErr != nil {
		t.Fatal(doc.LastErr)
	}
	// TODO: Compare to golden file
}
