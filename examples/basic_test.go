package examples

import (
	"bytes"
	"fmt"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

// Example defines a model with a single string called "message" as it's only allowed
// field (or 'leaf').
//
// Models can be matched to a data for lot's of things including reading data, so we create a simple
// data source that always returns hello.
//
// Models and Data come together as a browser.  A browser is all you need to do anything with the data
// that confirms to the model.
func Example_01Basic() {

	// Model - Normally module definition are stored on disk and found using YANGPATH environment
	// variable, but you can also load module definitions from strings.
	model, _ := yang.LoadModuleFromString(nil,

		// namespace, prefix and revision are required as part of YANG spec
		`module x {
			namespace "";
			prefix "";
			revision 0;

			leaf message { type string; }
		}`)

	// Node backs your application code. There are all sorts of Node implemetations including ones that
	// read JSON, use reflection, read maps but you can use OnField for custom field handling
	data := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = &node.Value{Str: "Hello"}
			return nil
		},
	}

	// Browser - Unites Model and Data to create a powerful way to interact with your application.
	// For example you can pass a browser to restconf package to host a REST API.
	brwsr := node.NewBrowser(model, data)

	// Here we read a simple value from our browser object.
	msg, _ := brwsr.Root().Get("message")
	fmt.Println(msg)

	// Output: Hello
}

// TestReadingMultipleLeafs expands on TestSimplestExample by adding multiple leafs. In the
// OnField method, we now need to differentiate which field we want to return
func Example_02writeJSON() {

	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			leaf to { type string; }
			leaf message { type string; }
		}`)

	data := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "to":
				hnd.Val = &node.Value{Str: "Mary"}
			case "message":
				hnd.Val = &node.Value{Str: "Hello"}
			}
			return nil
		},
	}

	brwsr := node.NewBrowser(model, data)
	var out bytes.Buffer

	// Convert everything to JSON
	if err := brwsr.Root().InsertInto(node.NewJsonWriter(&out).Node()).LastErr; err != nil {
		panic(err)
	}
	fmt.Println(out.String())

	// Output: {"to":"Mary","message":"Hello"}
}

// TestReadingStruct expands on TestSimplestExample by wrapping a 'container' around the
// message.  Containers are like a Golang struct.
func Example_readingStruct() {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			container suggestionBox {
			   leaf message { type string; }
			}
		}`)

	// Data
	messageData := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = &node.Value{Str: "Hello"}
			return nil
		},
	}
	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			return messageData, nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Read
	msg, _ := brwsr.Root().Find("suggestionBox").Get("message")
	fmt.Println(msg)

	// Output: Hello
}

/*
TestReadingList expands on TestSimplestExample by wrapping a 'list' around the
message.  Lists are like a Golang slices or arrays of structs. NOTE: If you want a list
of strings or ints, then you would just use a leaf-list

Output:
==============
Hello
*/
func Example_readingList() {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			list suggestionBox {
			   key "id";
			   leaf id {type string;}
			   leaf message { type string; }
			}
		}`)

	// Data
	messageData := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = &node.Value{Str: "Hello"}
			return nil
		},
	}
	dataList := &node.MyNode{
		OnNext: func(r node.ListRequest) (child node.Node, key []*node.Value, err error) {
			if r.Key[0].Str == "joe" {
				return messageData, nil, nil
			}
			return
		},
	}
	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			return dataList, nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Read
	msg, _ := brwsr.Root().Find("suggestionBox=joe").Get("message")
	fmt.Println(msg)
}

/*
TestSimplestExample defines a model with a single string called "message" as it's only allowed
field (or 'leaf').

Models can be matched to a data for lot's of things including reading data, so we create a simple
data source that always returns hello.

Models and Data come together as a browser.  A browser is all you need to do anything with the data
that confirms to the model.

Output:
==============
Hello
Goodbye
*/
func Example_readWrite() {

	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			leaf message { type string; }
		}`)

	// Data
	message := "Hello"
	data := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			if r.Write {
				message = hnd.Val.Str
			} else {
				hnd.Val = &node.Value{Str: message}
			}
			return nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Read
	msg, _ := brwsr.Root().Get("message")
	fmt.Println(msg)
	_ = brwsr.Root().Set("message", "Goodbye")
	msg, _ = brwsr.Root().Get("message")
	fmt.Println(msg)
}

/*
You can create containers, lists and list items.

When struct is being created
you do not have any fields.  When a list item is being created you do get the
key value.  If you need fields before constructing data, then implement OnEvent
and listen to node.NEW event.  If the fields in your model have default values
those fields will automatically be called.

Output
============
Finished creating suggestion box
{hello}
*/
type exampleBox struct {
	message string
}

func Example_addContainer() {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			container suggestionBox {
			  leaf message { type string; }
			}
		}`)

	// Data
	var box *exampleBox
	boxNode := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			box.message = hnd.Val.Str
			return nil
		},
		OnEndEdit: func(node.NodeRequest) error {
			fmt.Println("Finished creating suggestion box")
			return nil
		},
	}
	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "suggestionBox":
				if r.New {
					// You do not have any additional information
					box = &exampleBox{}
				}
				if box != nil {
					return boxNode, nil
				}
			}
			return nil, nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Delete
	brwsr.Root().InsertFrom(node.ReadJson(`{"suggestionBox":{"message":"hello"}}`))
	fmt.Println(*box)
}

/*
You can create containers, lists and list items.

When struct is being created
you do not have any fields.  When a list item is being created you do get the
key value.  If you need fields before constructing data, then implement OnEvent
and listen to node.NEW event.  If the fields in your model have default values
those fields will automatically be called.

Output
============
Finished creating suggestion box
map[212ea:hello]
*/

func Example_addListItem() {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			list suggestionBox {
			  key "id";
			  leaf id { type string; }
			  leaf message { type string; }
			}
		}`)

	// Data
	var box map[string]string
	var id string
	boxNode := &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "message":
				box[id] = hnd.Val.Str
			}
			return nil
		},
	}
	boxListNode := &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			key := r.Key
			if key != nil {
				id = key[0].Str
			}
			if r.New {
				box[id] = "new object"
			}
			if _, found := box[id]; found {
				return boxNode, key, nil
			}
			return nil, nil, nil
		},
		OnEndEdit: func(node.NodeRequest) error {
			fmt.Println("Finished creating suggestion box")
			return nil
		},
	}
	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "suggestionBox":
				if r.New {
					// You do not have any additional information
					box = make(map[string]string)
				}
				if box != nil {
					return boxListNode, nil
				}
			}
			return nil, nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Delete
	brwsr.Root().InsertFrom(node.ReadJson(`{"suggestionBox":[{"id":"abc", "message":"hello"}]}`))
	fmt.Println(box)
}

/*
Deleting is only for containers and lists.  You implement this by receiving events
on the node with a reference to the struct being deleted and/or the node of the
struct itself.

Output
=============
Deleting message hello
map[owner:map[name:joe]]
*/
func Example_delete() {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			container suggestionBox {
			   leaf message { type string; }
			}
			container owner {
			  leaf name { type string; }
			}
		}`)

	// Data
	box := map[string]interface{}{
		"message": "hello",
	}
	app := map[string]interface{}{
		"suggestionBox": box,
		"owner": map[string]interface{}{
			"name": "joe",
		},
	}
	boxData := &node.MyNode{
		OnDelete: func(n node.NodeRequest) error {
			// catch this event is the struct itself needs
			// to know it's being deleted.  In this case of
			fmt.Println("Deleting message ", box["message"])
			return nil
		},
	}

	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			if r.Delete {
				// catch this event for the owner of the struct to remove
				// references to the struct being deleted
				fmt.Println("Removing reference to ", r.Meta.GetIdent())
				delete(app, r.Meta.GetIdent())
			}
			return boxData, nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Delete
	brwsr.Root().Find("suggestionBox").Delete()
	fmt.Println(app)
}

/*
Deleting is only for containers and lists.  You implement this by receiving events
on the node with a reference to the struct being deleted and/or the node of the
struct itself.

Output
=============
42
*/
func Example_action() {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			rpc sum {
			   input {
			     leaf a {
			       type int32;
			     }
			     leaf b {
			       type int32;
			     }
			   }
			   output {
			     leaf result {
			       type int32;
			     }
			   }
			}
		}`)

	// Data
	data := &node.MyNode{
		OnAction: func(r node.ActionRequest) (out node.Node, err error) {
			switch r.Meta.GetIdent() {
			case "sum":
				var a, b *node.Value
				if a, err = r.Input.GetValue("a"); err != nil {
					return
				}
				if b, err = r.Input.GetValue("b"); err != nil {
					return
				}
				// Use map to return result, but you can use any node that can
				// represent the output model
				result := map[string]interface{}{
					"result": a.Int + b.Int,
				}
				return node.MapNode(result), nil
			}
			return
		},
	}

	// JSON is a useful format to use as input, but this can come from any node
	// that would return "a" and "b" leaves.
	input := node.ReadJson(`{"a":10,"b":32}`)

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Delete
	result := brwsr.Root().Find("sum").Action(input)
	v, _ := result.GetValue("result")
	fmt.Println(v.Int)
}
