package browse

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
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
		acls         []*AccessControl
		expected     error
		expectedSub  error
		expectedExec int
	}{
		{
			desc: "regex",
			acls: []*AccessControl{
				{
					Path:        ".*",
					Permissions: Read,
				},
			},
			expected:     nil,
			expectedSub:  nil,
			expectedExec: 401,
		},
		{
			desc: "parent path, but not all children",
			acls: []*AccessControl{
				{
					Path:        "^a$",
					Permissions: Read,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 401,
		},
		{
			desc: "parent's childern, not parent",
			acls: []*AccessControl{
				{
					Path:        "^b/ba",
					Permissions: Read,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  nil,
			expectedExec: 401,
		},
		{
			desc: "execute",
			acls: []*AccessControl{
				{
					Path:        "^a/x",
					Permissions: Execute,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 501,
		},
		{
			desc: "different path protected:",
			acls: []*AccessControl{
				{
					Path:        "^wrong",
					Permissions: Read,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 401,
		},
		{
			desc: "empty path same as root path:",
			acls: []*AccessControl{
				{
					Path:        "",
					Permissions: Read,
				},
			},
			expected:     nil,
			expectedSub:  nil,
			expectedExec: 401,
		},
		{
			desc: "can write, but not read:",
			acls: []*AccessControl{
				{
					Path:        "",
					Permissions: Write,
				},
			},
			expected:     UnauthorizedError,
			expectedSub:  UnauthorizedError,
			expectedExec: 401,
		},
		{
			desc: "multiple acls",
			acls: []*AccessControl{
				{
					Path:        "",
					Permissions: Write,
				},
				{
					Path:        "",
					Permissions: Read,
				},
			},
			expected:     nil,
			expectedSub:  nil,
			expectedExec: 401,
		},
	}
	for _, test := range tests {
		acl := NewRole()
		t.Log(test.desc)
		for _, testAcDef := range test.acls {
			acl.Access.PushBack(testAcDef)
		}

		s := b.Root()
		s.Constraints.AddConstraint("auth", 0, 0, acl)
		actual := s.InsertInto(node.DevNull()).LastErr
		if actual != test.expected {
			t.Error("Insert into root\n", c2.CheckEqual(test.expected, actual).Error())
			continue
		}

		path := "b/ba/baa"
		actualSub := s.Find(path).InsertInto(node.DevNull()).LastErr
		if actualSub != test.expectedSub {
			t.Error("Insert into path\n", c2.CheckEqual(test.expectedSub, actualSub).Error())
			continue
		}

		actualExec := s.Find("a/x").Action(nil).LastErr.(c2.HttpError).HttpCode()
		if actualExec != test.expectedExec {
			t.Error("Run action\n", c2.CheckEqual(test.expectedExec, actualExec).Error())
			continue
		}
	}
}
