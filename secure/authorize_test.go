package secure

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/freeconf/c2g/val"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
)

type testAc struct {
	path string
	perm Permission
}

const (
	xAllowed string = "allowed"
	xHidden         = "hidden"
	xUnauth         = "unauthorized"
)

func TestAuthConstraints(t *testing.T) {
	c2.DebugLog(true)
	m, err := yang.LoadModuleFromString(nil, `module birding { revision 0;
leaf count {
	type int32;
}
container owner {
	leaf name {
		type string;
	}
}
action fieldtrip {
	input {}
}
notification identified {}
	}`)
	if err != nil {
		t.Fatal(err)
	}
	dataStr := `{
		"count" : 10,
		"owner": {"name":"ethel"}
	}`
	var data map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(dataStr)).Decode(&data); err != nil {
		panic(err)
	}
	n := &nodes.Extend{
		Base: nodes.ReflectChild(data),
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			r.Send(&nodes.Basic{})
			closer := func() error { return nil }
			return closer, nil
		},
		OnAction: func(p node.Node, r node.ActionRequest) (node.Node, error) {
			return nil, nil
		},
	}
	b := node.NewBrowser(m, n)
	tests := []struct {
		desc      string
		acls      []*AccessControl
		read      string
		readPath  string
		write     string
		writePath string
		notify    string
		action    string
	}{
		{
			desc: "default",
			acls: []*AccessControl{
			/* empty */
			},
			read:      xHidden,
			readPath:  xHidden,
			write:     xUnauth,
			writePath: xHidden,
			notify:    xUnauth,
			action:    xUnauth,
		},
		{
			desc: "none",
			acls: []*AccessControl{
				{
					Path:        "birding",
					Permissions: Read,
				},
			},
			read:      xAllowed,
			readPath:  xAllowed,
			write:     xUnauth,
			writePath: xUnauth,
			notify:    xUnauth,
			action:    xUnauth,
		},
		{
			desc: "full",
			acls: []*AccessControl{
				{
					Path:        "birding",
					Permissions: Full,
				},
			},
			read:      xAllowed,
			readPath:  xAllowed,
			write:     xAllowed,
			writePath: xAllowed,
			notify:    xAllowed,
			action:    xAllowed,
		},
		{
			desc: "mixed",
			acls: []*AccessControl{
				{
					Path:        "birding",
					Permissions: Full,
				},
				{
					Path:        "birding/owner",
					Permissions: None,
				},
			},
			read:      xAllowed,
			readPath:  xHidden,
			write:     xAllowed,
			writePath: xHidden,
			notify:    xAllowed,
			action:    xAllowed,
		},
	}
	for _, test := range tests {
		acl := NewRole()
		for _, testAcDef := range test.acls {
			acl.Access[testAcDef.Path] = testAcDef
		}

		s := b.Root()
		s.Constraints.AddConstraint("auth", 0, 0, acl)
		s.Context = s.Constraints.ContextConstraint(s)

		t.Log(test.desc + " read")
		c2.AssertEqual(t, test.read, val2auth(s.GetValue("count")))

		t.Logf(test.desc + " read path")
		pathSel := s.Find("owner")
		c2.AssertEqual(t, test.readPath, sel2auth(pathSel))

		t.Log(test.desc + " write")
		writeErr := s.Set("count", 100)
		c2.AssertEqual(t, test.write, err2auth(writeErr))

		t.Log(test.desc + " write path")
		if pathSel.IsNil() {
			c2.AssertEqual(t, test.writePath, xHidden)
		} else {
			writePathErr := pathSel.Set("name", "Harvey")
			c2.AssertEqual(t, test.writePath, err2auth(writePathErr))
		}

		t.Log(test.desc + " execute")
		actionErr := s.Find("fieldtrip").Action(nil).LastErr
		c2.AssertEqual(t, test.action, err2auth(actionErr))

		t.Log(test.desc + " notify")
		var notifyErr error
		s.Find("identified").Notifications(func(m node.Selection) {
			notifyErr = m.LastErr
		})
		c2.AssertEqual(t, test.notify, err2auth(notifyErr))
	}
}

func val2auth(v val.Value, err error) string {
	if v == nil {
		return xHidden
	}
	if err == UnauthorizedError {
		return xUnauth
	}
	return xAllowed
}

func sel2auth(s node.Selection) string {
	if s.IsNil() {
		return xHidden
	}
	return xAllowed
}

func err2auth(err error) string {
	if err == nil {
		return xAllowed
	} else if err == UnauthorizedError {
		return xUnauth
	}
	panic(err.Error())
}
