package node

//import (
//	"bytes"
//	"github.com/c2stack/c2g/meta"
//	"github.com/c2stack/c2g/meta/yang"
//	"strings"
//	"testing"
//)
//
//func TestKeyListBuilderInBufferStore(t *testing.T) {
//	tests := []struct {
//		path     string
//		expected string
//	}{
//		{"a/a", ""},
//		{"a/b", "x"},
//		{"a/c", "y|z"},
//	}
//	store := NewBufferStore()
//	v := &Value{}
//	store.Values["a/a/c"] = v
//	store.Values["a/b=x/c"] = v
//	store.Values["a/c=y/c"] = v
//	store.Values["a/c=y/c"] = v
//	store.Values["a/c=z/q/f=yy/fg=gf/gf"] = v
//	m := &meta.List{Ident: "c", Key: []string{"k"}}
//	m.AddMeta(meta.NewLeaf("k", "string"))
//	for _, test := range tests {
//		keys, err := store.KeyList(test.path, m)
//		if err != nil {
//			t.Error(err)
//		}
//		actual := strings.Join(keys, "|")
//		if actual != test.expected {
//			t.Errorf("Test failed for path %s\nExpected:%s\n  Actual:%s", test.path, test.expected, actual)
//		}
//	}
//}
//
//func keyValuesTestModule() *meta.Module {
//	mstr := `
//module m {
//	prefix "";
//	namespace "";
//	revision 0;
//	container a {
//		container aa {
//			leaf aaa {
//				type string;
//			}
//		}
//		leaf ab {
//			type string;
//		}
//	}
//	list b {
//		key "ba";
//		leaf ba {
//			type string;
//		}
//		container bb {
//			leaf bba {
//				type string;
//			}
//		}
//		list bc {
//			key "bca";
//			leaf bca {
//				type string;
//			}
//		}
//	}
//}
//`
//	m, err := yang.LoadModuleCustomImport(mstr, nil)
//	if err != nil {
//		panic(err)
//	}
//	return m
//}
//
//func TestStoreBrowserKeyValueRead(t *testing.T) {
//	store := NewBufferStore()
//	m := keyValuesTestModule()
//	kv := NewStoreData(m, store)
//	store.Values["a/aa/aaa"] = &Value{Str: "hi"}
//	store.Values["b=x/ba"] = &Value{Str: "x"}
//	var actualBytes bytes.Buffer
//	if err := kv.Browser().Root().InsertInto(NewJsonWriter(&actualBytes).Node()).LastErr; err != nil {
//		t.Error(err)
//	}
//	actual := string(actualBytes.Bytes())
//	expected := `{"a":{"aa":{"aaa":"hi"}},"b":[{"ba":"x"}]}`
//	if actual != expected {
//		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
//	}
//}
//
//func TestStoreBrowserValueEdit(t *testing.T) {
//	store := NewBufferStore()
//	m := keyValuesTestModule()
//	kv := NewStoreData(m, store)
//	inputJson := `{"a":{"aa":{"aaa":"hi"}},"b":[{"ba":"x"}]}`
//	json := NewJsonReader(strings.NewReader(inputJson)).Node()
//	if err := kv.Browser().Root().InsertFrom(json).LastErr; err != nil {
//		t.Fatal(err)
//	}
//	if len(store.Values) != 2 {
//		t.Error("Expected 2 items")
//	}
//	expectations := []struct {
//		path  string
//		value string
//	}{
//		{"a/aa/aaa", "hi"},
//		{"b=x/ba", "x"},
//	}
//	if len(expectations) != len(store.Values) {
//		t.Errorf("Expected %d items but got", len(expectations), len(store.Values))
//	}
//	for _, expected := range expectations {
//		v, found := store.Values[expected.path]
//		if !found {
//			t.Error("Could not find item", expected.path, "\nItems: ", store)
//		} else if v.Str != expected.value {
//			t.Error("Expected value to be %s but was %s", expected.value, v.Str)
//		}
//	}
//}
//
//func TestStoreBrowserKeyValueEdit(t *testing.T) {
//	store := NewBufferStore()
//	m := keyValuesTestModule()
//	kv := NewStoreData(m, store)
//	store.Values["b=x/ba"] = &Value{Str: "z"}
//
//	// change key
//	json := NewJsonReader(strings.NewReader(`{"ba":"y"}`)).Node()
//	if err := kv.Browser().Root().Find("b=x").UpdateFrom(json).LastErr; err != nil {
//		t.Fatal(err)
//	}
//	if v, newKeyExists := store.Values["b=y/ba"]; !newKeyExists {
//		t.Error("Edit key value not made")
//	} else if v.Str != "y" {
//		t.Error("Wrong key value")
//	}
//	if _, oldKeyExists := store.Values["/b=x/ba"]; oldKeyExists {
//		t.Error("Old key was not removed")
//	}
//}
//
//func TestStoreBrowserReadListList(t *testing.T) {
//	store := NewBufferStore()
//	m := keyValuesTestModule()
//	kv := NewStoreData(m, store)
//	store.Values["b=x/ba"] = &Value{Str: "x"}
//	store.Values["b=x/bc=y/bca"] = &Value{Str: "y"}
//	var actual bytes.Buffer
//	if err := kv.Browser().Root().UpsertInto(NewJsonWriter(&actual).Node()).LastErr; err != nil {
//		t.Error(err)
//	}
//	t.Log(actual.String())
//}
//
//func TestStoreRemoveAll(t *testing.T) {
//	store := NewBufferStore()
//	m := keyValuesTestModule()
//	store.Values["b=x/ba"] = &Value{Str: "x"}
//	store.Values["b=x/bc=y/bca"] = &Value{Str: "y"}
//	kv := NewStoreData(m, store)
//	sel := kv.Browser().Root().Find("b=x/bc")
//	if sel.LastErr != nil {
//		t.Fatal(sel.LastErr)
//	}
//	if err := sel.Delete(); err != nil {
//		t.Error(err)
//	}
//	if len(store.Values) != 1 {
//		t.Error("Expected only 1 value after delete")
//	}
//	if _, found := store.Values["b=x/ba"]; !found {
//		t.Error("Remaining value after delete is wrong")
//	}
//}
