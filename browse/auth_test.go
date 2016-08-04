package browse

import (
	"testing"
	"github.com/c2g/meta/yang"
	"encoding/json"
	"strings"
	"github.com/c2g/node"
	"bytes"
	"github.com/c2g/c2"
	"fmt"
)

func TestAuthConstraints(t *testing.T) {
	m, err := yang.LoadModuleFromString(nil, `module m { namespace ""; prefix ""; revision 0;
container a {
	leaf aa {
		type string;
	}
}

	}`)
	if err != nil {
		t.Fatal(err)
	}
	dataStr := `{
		"a" : {
			"aa" : "hello"
		}
	}`
	var data map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(dataStr)).Decode(&data); err != nil {
		panic(err)
	}
	b := node.NewBrowser2(m, node.MapNode(data))

	tests := []struct{
		path string
		selector string
		perm Permission
		expected error
	} {
		{
			path : ".*",
			perm: Read,
			expected: nil,
		},
		{
			path : "x",
			perm: Read,
			expected: UnauthorizedError,
		},
		{
			path : "",
			perm: Read,
			expected: nil,
		},
		{
			path : "",
			perm: Write,
			expected: UnauthorizedError,
		},
	}

	for i, test := range tests {
		acl := NewAcl()
		ac := &AccessControl{Permissions:test.perm}
		ac.SetPath(test.path)
		acl.List.PushBack(ac)
		s := b.Root().Selector()
		s.Constraints().AddConstraint("auth", 0, 0, acl)
		var buff bytes.Buffer
		out := node.NewJsonWriter(&buff).Node()
		actual := s.InsertInto(out).LastErr
		if actual != test.expected {
			t.Error(fmt.Sprintf("Read Test %d%s", i + 1, c2.CheckEqual(test.expected, actual).Error()))
			continue
		}
	}
}
