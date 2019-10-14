package nodes_test

import (
	"testing"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/testdata"
)

func TestMetaNameToFieldName(t *testing.T) {
	var actual string
	tests := []struct {
		in  string
		out string
	}{
		{"X", "X"},
		{"x", "X"},
		{"abc", "Abc"},
		{"ABC", "ABC"},
		{"abCd", "AbCd"},
		{"one-two", "OneTwo"},
	}
	for _, test := range tests {
		if actual = nodes.MetaNameToFieldName(test.in); actual != test.out {
			t.Error(test.out, "!=", actual)
		}
	}
}

var m1 = `module m {
	revision 0;

	leaf name {
		type string;
	}
	container species {
		leaf name {
			type string;
		}
		leaf class {
			type string;
		}
	}
}
`
var m2 = `module m {
	revision 0;

	list birds {
		key "name";
		leaf name {
			type string;
		}
		container species {
			leaf name {
				type string;
			}
			leaf class {
				type string;
			}
		}
	}
}
`

func TestReflect2Write(t *testing.T) {
	var b *node.Browser
	write := func(n node.Node, mstr string, data string) {
		m, err := parser.LoadModuleFromString(nil, mstr)
		if err != nil {
			t.Fatal(err)
		}
		b = node.NewBrowser(m, n)
		sel := b.Root()
		if err = sel.UpsertFrom(nodes.ReadJSON(data)).LastErr; err != nil {
			t.Error(err)
		}
	}
	// structs
	{
		bird := &testdata.Bird{}
		write(nodes.ReflectChild(bird), m1, `{"name":"robin"}`)
		c2.AssertEqual(t, "robin", bird.Name)
	}
	// structs / structs
	{
		bird := &testdata.Bird{}
		write(nodes.ReflectChild(bird), m1, `{"name":"robin","species":{"name":"thrush"}}`)
		c2.AssertEqual(t, "robin", bird.Name)
		c2.AssertEqual(t, "thrush", bird.Species.Name)
	}
	// maps / maps
	{
		bird := map[string]interface{}{}
		write(nodes.ReflectChild(bird), m1, `{"name":"robin","species":{"name":"thrush"}}`)
		c2.AssertEqual(t, "robin", bird["name"])
		c2.AssertEqual(t, "thrush", mapValue(bird, "species", "name"))

		// delete
		if err := b.Root().Find("species").Delete(); err != nil {
			t.Error(err)
		} else {
			c2.AssertEqual(t, nil, bird["species"])
		}
	}
	// maps(list) / maps
	{
		birds := map[string]interface{}{}
		write(nodes.ReflectChild(birds), m2, `{"birds":[{"name":"robin","species":{"name":"thrush"}}]}`)
		c2.AssertEqual(t, "thrush", mapValue(birds, "birds", "robin", "species", "name"))

		// delete
		if err := b.Root().Find("birds=robin").Delete(); err != nil {
			t.Error(err)
		} else {
			b := birds["birds"].(map[string]interface{})
			c2.AssertEqual(t, 0, len(b))
		}
	}
	// maps(list) / structs
	{
		app := struct {
			Birds map[string]*testdata.Bird
		}{}
		n := nodes.ReflectChild(&app)
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"thrush"}}]}`)
		robin, exists := app.Birds["robin"]
		if !exists {
			t.Fail()
		}
		c2.AssertEqual(t, "robin", robin.Name)
		c2.AssertEqual(t, "thrush", robin.Species.Name)

		// update
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"DC Comics"}}]}`)
		c2.AssertEqual(t, "DC Comics", robin.Species.Name)

		// delete
		if err := b.Root().Find("birds=robin").Delete(); err != nil {
			t.Error(err)
		} else {
			c2.AssertEqual(t, 0, len(app.Birds))
		}
	}
	// slice(list) / structs
	{
		app := struct {
			Birds []*testdata.Bird
		}{}
		n := nodes.ReflectChild(&app)
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"thrush"}}]}`)
		if len(app.Birds) != 1 {
			t.Fail()
		}
		c2.AssertEqual(t, "robin", app.Birds[0].Name)
		c2.AssertEqual(t, "thrush", app.Birds[0].Species.Name)

		// update
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"DC Comics"}}]}`)
		c2.AssertEqual(t, "DC Comics", app.Birds[0].Species.Name)

		// delete
		if err := b.Root().Find("birds=robin").Delete(); err != nil {
			t.Error(err)
		} else {
			c2.AssertEqual(t, 0, len(app.Birds))
		}
	}
}

func mapValue(m map[string]interface{}, key ...string) interface{} {
	r, exists := m[key[0]]
	if len(key) > 1 && exists {
		return mapValue(r.(map[string]interface{}), key[1:]...)
	}
	return r
}

func Test_Reflect2Read(t *testing.T) {
	read := func(n node.Node, mstr string) string {
		m, err := parser.LoadModuleFromString(nil, mstr)
		if err != nil {
			t.Fatal(err)
		}
		s, err := nodes.WriteJSON(node.NewBrowser(m, n).Root())
		if err != nil {
			t.Error(err)
		}
		return s
	}
	// structs
	{
		bird := &testdata.Bird{Name: "robin"}
		c2.AssertEqual(t, `{"name":"robin"}`, read(nodes.ReflectChild(bird), m1))
	}
	// structs / structs
	{
		bird := &testdata.Bird{Name: "robin", Species: &testdata.Species{Name: "thrush"}}
		c2.AssertEqual(t, `{"name":"robin","species":{"name":"thrush"}}`, read(nodes.ReflectChild(bird), m1))
	}
	// maps
	{
		bird := map[string]interface{}{"name": "robin"}
		c2.AssertEqual(t, `{"name":"robin"}`, read(nodes.ReflectChild(bird), m1))
	}
	// maps / maps
	{
		bird := map[string]interface{}{
			"name": "robin",
			"species": map[string]interface{}{
				"name": "thrush",
			},
		}
		c2.AssertEqual(t, `{"name":"robin","species":{"name":"thrush"}}`, read(nodes.ReflectChild(bird), m1))
	}
	// maps(list) / struct
	{
		birds := map[string]interface{}{
			"birds": map[string]*testdata.Bird{
				"robin": &testdata.Bird{
					Name: "robin",
				},
			},
		}
		actual := read(nodes.ReflectChild(birds), m2)
		c2.AssertEqual(t, `{"birds":[{"name":"robin"}]}`, actual)
	}
	// maps(list) / maps
	{
		birds := map[string]interface{}{
			"birds": map[string]interface{}{
				"robin": map[string]interface{}{
					"name": "robin",
				},
			},
		}
		actual := read(nodes.ReflectChild(birds), m2)
		c2.AssertEqual(t, `{"birds":[{"name":"robin"}]}`, actual)
	}
}

type TestMessage struct {
	Message struct {
		Hello string
	}
}

func TestMarshal(t *testing.T) {
	mstr := `
module m {
	prefix "";
	namespace "";
	revision 0;
	container message {
		leaf hello {
			type string;
		}
	}
}
`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	var obj TestMessage
	c := nodes.ReflectChild(&obj)
	sel := node.NewBrowser(m, c).Root()
	r := nodes.ReadJSON(`{"message":{"hello":"bob"}}`)
	if err = sel.UpsertFrom(r).LastErr; err != nil {
		t.Fatal(err)
	}
	if obj.Message.Hello != "bob" {
		t.Fatal("Not selected")
	}
}

type TestMessageItem struct {
	Id string
}

var mstr = `
module m {
	namespace "";
	prefix "";
	revision 0;
	container a {
		container b {
			leaf x {
				type string;
			}
		}
	}
	list p {
		key "k";
		leaf k {
			type string;
		}
		container q {
			leaf s {
				type string;
			}
		}
		list r {
			leaf z {
				type int32;
			}
		}
	}
}
`

func TestCollectionWrite(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		data string
		path string
	}{
		{
			`{"a":{"b":{"x":"waldo"}}}`,
			"a.b.x",
		},
		{
			`{"p":[{"k":"waldo"},{"k":"walter"},{"k":"weirdo"}]}`,
			"p.waldo.k",
		},
	}
	for _, test := range tests {
		root := make(map[string]interface{})
		bd := nodes.ReflectChild(root)
		sel := node.NewBrowser(m, bd).Root()
		if err = sel.InsertFrom(nodes.ReadJSON(test.data)).LastErr; err != nil {
			t.Error(err)
		}
		actual := node.MapValue(root, test.path)
		if actual != "waldo" {
			t.Error(actual)
		}
	}
}

func TestCollectionRead(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		root     map[string]interface{}
		expected string
	}{
		{
			map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"x": "waldo",
					},
				},
			},
			`{"a":{"b":{"x":"waldo"}}}`,
		},
		{
			map[string]interface{}{
				"p": []interface{}{
					map[string]interface{}{"k": "walter"},
					map[string]interface{}{"k": "waldo"},
					map[string]interface{}{"k": "weirdo"},
				},
			},
			`{"p":[{"k":"waldo"},{"k":"walter"},{"k":"weirdo"}]}`,
		},
	}
	for _, test := range tests {
		bd := nodes.ReflectChild(test.root)
		sel := node.NewBrowser(m, bd).Root()
		if actual, err := nodes.WriteJSON(sel); err != nil {
			t.Error(err)
		} else if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}

func TestCollectionDelete(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		root     map[string]interface{}
		path     string
		expected string
	}{
		{
			map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"x": "waldo",
					},
				},
			},
			"a/b",
			`{"a":{}}`,
		},
		{
			map[string]interface{}{
				"p": []interface{}{
					map[string]interface{}{"k": "walter"},
					map[string]interface{}{"k": "waldo"},
					map[string]interface{}{"k": "weirdo"},
				},
			},
			"p=walter",
			`{"p":[{"k":"waldo"},{"k":"weirdo"}]}`,
		},
	}
	for _, test := range tests {
		bd := nodes.ReflectChild(test.root)
		sel := node.NewBrowser(m, bd).Root()
		if err := sel.Find(test.path).Delete(); err != nil {
			t.Error(err)
		}
		if actual, err := nodes.WriteJSON(sel); err != nil {
			t.Error(err)
		} else if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}
