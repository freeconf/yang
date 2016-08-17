package browse

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/dhubler/c2g/c2"
	"github.com/dhubler/c2g/meta/yang"
	"github.com/dhubler/c2g/node"
)

type testAc struct {
	path string
	perm Permission
}

func TestAuthConstraints(t *testing.T) {
	m, err := yang.LoadModuleFromString(nil, `module m { namespace ""; prefix ""; revision 0;
container a {
	leaf aa {
		type string;
	}
	action x {
		input {}
	}
}
container b {
	container ba {
		container baa {
			leaf baaa {
				type string;
			}
		}
	}
}

	}`)
	if err != nil {
		t.Fatal(err)
	}
	dataStr := `{
		"a" : { "aa" : "hello" },
		"b" : { "ba" : { "baa" : { "baaa" : "bye" } } }
	}`
	var data map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(dataStr)).Decode(&data); err != nil {
		panic(err)
	}
	b := node.NewBrowser2(m, node.MapNode(data))

	tests := []struct {
		desc         string
		acls         []testAc
		expected     error
		expectedSub  error
		expectedExec int
	}{
		{
			desc: "regex",
			acls: []testAc{
				{
					path: ".*",
					perm: Read,
				},
			},
			expected:     nil,
			expectedSub:  nil,
			expectedExec: 401,
		},
		{
			desc: "parent path, but not all children",
			acls: []testAc{
				{
					path: "^a$",
					perm: Read,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 401,
		},
		{
			desc: "parent's childern, not parent",
			acls: []testAc{
				{
					path: "^b/ba",
					perm: Read,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  nil,
			expectedExec: 401,
		},
		{
			desc: "execute",
			acls: []testAc{
				{
					path: "^a/x",
					perm: Execute,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 501,
		},
		{
			desc: "different path protected",
			acls: []testAc{
				{
					path: "^wrong",
					perm: Read,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 401,
		},
		{
			desc: "empty path same as root path",
			acls: []testAc{
				{
					path: "",
					perm: Read,
				},
			},
			expected:     nil,
			expectedSub:  nil,
			expectedExec: 401,
		},
		{
			desc: "can write, but not read",
			acls: []testAc{
				{
					path: "",
					perm: Write,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 401,
		},
		{
			desc: "multiple acls",
			acls: []testAc{
				{
					path: "",
					perm: Write,
				},
				{
					path: "",
					perm: Read,
				},
			},
			expected:     nil,
			expectedSub:  nil,
			expectedExec: 401,
		},
	}
	for _, test := range tests {
		acl := NewRole()
		for _, testAcDef := range test.acls {
			testAc := &AccessControl{Permissions: testAcDef.perm}
			testAc.SetPath(testAcDef.path)
			acl.Access.PushBack(testAc)
		}
		s := b.Root().Selector()
		s.Constraints().AddConstraint("auth", 0, 0, acl)
		actual := s.InsertInto(node.DevNull()).LastErr
		if actual != test.expected {
			t.Error(fmt.Sprintf("(root) %s Root - %s", test.desc, c2.CheckEqual(test.expected, actual).Error()))
			continue
		}

		path := "b/ba/baa"
		actualSub := s.Find(path).InsertInto(node.DevNull()).LastErr
		if actualSub != test.expectedSub {
			t.Error(fmt.Sprintf("(%s) %s - %s", path, test.desc, c2.CheckEqual(test.expectedSub, actualSub).Error()))
			continue
		}

		actualExec := s.Find("a/x").Action(nil).LastErr.(c2.HttpError).HttpCode()
		if actualExec != test.expectedExec {
			t.Error(fmt.Sprintf("Execute %s - %s", test.desc, c2.CheckEqual(test.expectedExec, actualExec).Error()))
			continue
		}
	}
}
