package nodeutil_test

import (
	"fmt"
	"net/netip"
	"reflect"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/testdata"
	"github.com/freeconf/yang/val"
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
		if actual = nodeutil.MetaNameToFieldName(test.in); actual != test.out {
			t.Error(test.out, "!=", actual)
		}
	}
}

func TestMetaNameToFieldNameExt(t *testing.T) {
	var actual string

	data := struct {
		Renamed    string `yang:"whatever"`
		Notrenamed string
	}{}

	tests := []struct {
		in  string
		out string
	}{
		{in: "Renamed", out: "Renamed"},
		{in: "renamed", out: "Renamed"},
		{in: "whatever", out: "Renamed"},
		{in: "Notrenamed", out: "Notrenamed"},
		{in: "notrenamed", out: "Notrenamed"},
	}
	for _, test := range tests {
		if actual = nodeutil.GetFieldName(reflect.ValueOf(data), test.in); actual != test.out {
			t.Errorf("%v should be %v, got %v", test.in, test.out, actual)
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
		if err = sel.UpsertFrom(nodeutil.ReadJSON(data)).LastErr; err != nil {
			t.Error(err)
		}
	}
	// structs
	{
		bird := &testdata.Bird{}
		write(nodeutil.ReflectChild(bird), m1, `{"name":"robin"}`)
		fc.AssertEqual(t, "robin", bird.Name)
	}
	// struct + field conversions on write
	{
		ipBird := &testdata.IPBird{}
		ref := nodeutil.Reflect{
			OnField: []nodeutil.ReflectField{
				{
					When: nodeutil.ReflectFieldByType(reflect.TypeOf(netip.Addr{})),
					OnWrite: func(_ meta.Leafable, fieldname string, elem reflect.Value, fieldElem reflect.Value, v val.Value) error {
						ip, err := netip.ParseAddr(fmt.Sprint(v.Value()))
						if err != nil {
							return err
						}
						fieldElem.Set(reflect.ValueOf(ip))
						return nil
					},
				},
			},
		}
		write(ref.Object(ipBird), m1, `{"name":"10.0.0.1","species":{"name":"10.0.0.2"}}`)
		fc.AssertEqual(t, netip.MustParseAddr("10.0.0.1"), ipBird.Name)
		fc.AssertEqual(t, netip.MustParseAddr("10.0.0.2"), ipBird.Species.Name)
	}
	// structs / structs
	{
		bird := &testdata.Bird{}
		write(nodeutil.ReflectChild(bird), m1, `{"name":"robin","species":{"name":"thrush"}}`)
		fc.AssertEqual(t, "robin", bird.Name)
		fc.AssertEqual(t, "thrush", bird.Species.Name)
	}
	// maps / maps
	{
		bird := map[string]interface{}{}
		write(nodeutil.ReflectChild(bird), m1, `{"name":"robin","species":{"name":"thrush"}}`)
		fc.AssertEqual(t, "robin", bird["name"])
		fc.AssertEqual(t, "thrush", fc.MapValue(bird, "species", "name"))

		// delete
		if err := b.Root().Find("species").Delete(); err != nil {
			t.Error(err)
		} else {
			fc.AssertEqual(t, nil, bird["species"])
		}
	}
	// maps(list) / maps
	{
		birds := map[string]interface{}{}
		write(nodeutil.ReflectChild(birds), m2, `{"birds":[{"name":"robin","species":{"name":"thrush"}}]}`)
		fc.AssertEqual(t, "thrush", fc.MapValue(birds, "birds", "robin", "species", "name"))

		// delete
		if err := b.Root().Find("birds=robin").Delete(); err != nil {
			t.Error(err)
		} else {
			b := birds["birds"].(map[string]interface{})
			fc.AssertEqual(t, 0, len(b))
		}
	}
	// maps(list) / structs
	{
		app := struct {
			Birds map[string]*testdata.Bird
		}{}
		n := nodeutil.ReflectChild(&app)
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"thrush"}}]}`)
		robin, exists := app.Birds["robin"]
		if !exists {
			t.Fail()
		}
		fc.AssertEqual(t, "robin", robin.Name)
		fc.AssertEqual(t, "thrush", robin.Species.Name)

		// update
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"DC Comics"}}]}`)
		fc.AssertEqual(t, "DC Comics", robin.Species.Name)

		// delete
		if err := b.Root().Find("birds=robin").Delete(); err != nil {
			t.Error(err)
		} else {
			fc.AssertEqual(t, 0, len(app.Birds))
		}
	}
	// slice(list) / structs
	{
		app := struct {
			Birds []*testdata.Bird
		}{}
		n := nodeutil.ReflectChild(&app)
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"thrush"}}]}`)
		if len(app.Birds) != 1 {
			t.Fail()
		}
		fc.AssertEqual(t, "robin", app.Birds[0].Name)
		fc.AssertEqual(t, "thrush", app.Birds[0].Species.Name)

		// update
		write(n, m2, `{"birds":[{"name":"robin","species":{"name":"DC Comics"}}]}`)
		fc.AssertEqual(t, "DC Comics", app.Birds[0].Species.Name)

		// delete
		if err := b.Root().Find("birds=robin").Delete(); err != nil {
			t.Error(err)
		} else {
			fc.AssertEqual(t, 0, len(app.Birds))
		}
	}
}

func Test_Reflect2Read(t *testing.T) {
	read := func(n node.Node, mstr string) string {
		m, err := parser.LoadModuleFromString(nil, mstr)
		if err != nil {
			t.Fatal(err)
		}
		s, err := nodeutil.WriteJSON(node.NewBrowser(m, n).Root())
		if err != nil {
			t.Error(err)
		}
		return s
	}
	// structs
	{
		bird := &testdata.Bird{Name: "robin"}
		fc.AssertEqual(t, `{"name":"robin"}`, read(nodeutil.ReflectChild(bird), m1))
	}
	// structs / structs
	{
		bird := &testdata.Bird{Name: "robin", Species: &testdata.Species{Name: "thrush"}}
		fc.AssertEqual(t, `{"name":"robin","species":{"name":"thrush"}}`, read(nodeutil.ReflectChild(bird), m1))
	}
	// struct w/ field conversion on read
	{
		ipbird := &testdata.IPBird{
			Name: netip.MustParseAddr("10.0.0.1"),
			Species: &testdata.IPSpecies{
				Name: netip.MustParseAddr("10.0.0.2"),
			},
		}
		ref := nodeutil.Reflect{
			OnField: []nodeutil.ReflectField{
				{
					When: nodeutil.ReflectFieldByType(reflect.TypeOf(netip.Addr{})),
					OnRead: func(leaf meta.Leafable, fieldname string, elem reflect.Value, fieldVal reflect.Value) (val.Value, error) {
						if leaf.Type().Format() != val.FmtString {
							return nil, fmt.Errorf("format should be string: %s", leaf.Type().Format().String())
						}
						if fieldVal.Type() != reflect.TypeOf(netip.Addr{}) {
							return nil, fmt.Errorf("input should be netip.Addr: %v", fieldVal.Type())
						}
						addr := fieldVal.Interface().(netip.Addr)
						return node.NewValue(leaf.Type(), fmt.Sprintf("ip: %v", addr.String()))
					},
				},
			},
		}
		fc.AssertEqual(t, `{"name":"ip: 10.0.0.1","species":{"name":"ip: 10.0.0.2"}}`, read(ref.Object(ipbird), m1))
	}
	// maps
	{
		bird := map[string]interface{}{"name": "robin"}
		fc.AssertEqual(t, `{"name":"robin"}`, read(nodeutil.ReflectChild(bird), m1))
	}
	// maps / maps
	{
		bird := map[string]interface{}{
			"name": "robin",
			"species": map[string]interface{}{
				"name": "thrush",
			},
		}
		fc.AssertEqual(t, `{"name":"robin","species":{"name":"thrush"}}`, read(nodeutil.ReflectChild(bird), m1))
	}
	// maps(list) / struct
	{
		birds := map[string]interface{}{
			"birds": map[string]*testdata.Bird{
				"robin": {
					Name: "robin",
				},
			},
		}
		actual := read(nodeutil.ReflectChild(birds), m2)
		fc.AssertEqual(t, `{"birds":[{"name":"robin"}]}`, actual)
	}
	// maps(list) / non-pointer struct
	{
		birds := map[string]interface{}{
			"birds": map[string]testdata.Bird{
				"robin": {
					Name: "robin",
				},
			},
		}
		actual := read(nodeutil.ReflectChild(birds), m2)
		fc.AssertEqual(t, `{"birds":[{"name":"robin"}]}`, actual)
	}
	// maps(list) / struct(stringer key), sorting only fails when at least two keys are present
	{
		birds := map[string]interface{}{
			"birds": map[netip.Addr]*testdata.IPBird{
				netip.MustParseAddr("10.0.0.1"): {
					Name: netip.MustParseAddr("10.0.0.1"),
				},
				netip.MustParseAddr("10.0.0.2"): {
					Name: netip.MustParseAddr("10.0.0.2"),
				},
			},
		}
		actual := read(nodeutil.ReflectChild(birds), m2)
		fc.AssertEqual(t, `{"birds":[{"name":"10.0.0.1"},{"name":"10.0.0.2"}]}`, actual)
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
		actual := read(nodeutil.ReflectChild(birds), m2)
		fc.AssertEqual(t, `{"birds":[{"name":"robin"}]}`, actual)
	}
	// stringer
	{
		birds := map[string]interface{}{
			"birds": map[string]interface{}{
				"robin": &struct {
					Name netip.Addr
				}{
					Name: netip.MustParseAddr("127.0.0.1"),
				},
			},
		}
		actual := read(nodeutil.ReflectChild(birds), m2)
		fc.AssertEqual(t, `{"birds":[{"name":"127.0.0.1"}]}`, actual)
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
	c := nodeutil.ReflectChild(&obj)
	sel := node.NewBrowser(m, c).Root()
	r := nodeutil.ReadJSON(`{"message":{"hello":"bob"}}`)
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
		path []interface{}
	}{
		{
			`{"a":{"b":{"x":"waldo"}}}`,
			[]interface{}{"a", "b", "x"},
		},
		{
			`{"p":[{"k":"waldo"},{"k":"walter"},{"k":"weirdo"}]}`,
			[]interface{}{"p", "waldo", "k"},
		},
	}
	for _, test := range tests {
		root := make(map[string]interface{})
		bd := nodeutil.ReflectChild(root)
		sel := node.NewBrowser(m, bd).Root()
		if err = sel.InsertFrom(nodeutil.ReadJSON(test.data)).LastErr; err != nil {
			t.Error(err)
		}
		actual := fc.MapValue(root, test.path...)
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
			`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
		},
	}
	for _, test := range tests {
		bd := nodeutil.ReflectChild(test.root)
		sel := node.NewBrowser(m, bd).Root()
		if actual, err := nodeutil.WriteJSON(sel); err != nil {
			t.Error(err)
		} else if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}

func TestCollectionNonStringKey(t *testing.T) {
	mstr := `module m {
		namespace "";
		prefix "";
		revision 0;	
		list x {
			key id;
			leaf id {
				type int32;
			}
			leaf data {
				type string;
			}
		}			
}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	data := map[string]interface{}{
		"x": map[int]interface{}{
			100: map[string]interface{}{
				"id":   100,
				"data": "hello",
			},
		},
	}
	b := node.NewBrowser(m, nodeutil.ReflectChild(data))
	actual, err := nodeutil.WriteJSON(b.Root())
	if err != nil {
		t.Error(err)
	}
	expected := `{"x":[{"id":100,"data":"hello"}]}`
	fc.AssertEqual(t, expected, actual)

	wtr := make(map[string]interface{})
	err = b.Root().UpsertInto(nodeutil.ReflectChild(wtr)).LastErr
	if err != nil {
		t.Error(err)
	}
	fc.AssertEqual(t, "map[x:map[100:map[data:hello id:100]]]", fmt.Sprintf("%v", wtr))
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
		bd := nodeutil.ReflectChild(test.root)
		sel := node.NewBrowser(m, bd).Root()
		if err := sel.Find(test.path).Delete(); err != nil {
			t.Error(err)
		}
		if actual, err := nodeutil.WriteJSON(sel); err != nil {
			t.Error(err)
		} else if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}
