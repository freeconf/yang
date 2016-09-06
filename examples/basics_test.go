package examples

/*
Basic Examples
================
Here are show some very basic operations. Nothing in these examples will be too exciting and
using C2G would be overkill in all these examples. These examples are only meant to show basic
API usage.  It's worth noting that these examples run without a web server so they are useful
beyond creating web services.  They also require extra steps in your development process like
code generatation.

Important notes about these examples:
 - rarely handle errors correctly to make the code easier to understand.
 - data models are defined in code and are normally located in separate *.yang files
   so they can be shared or used to create generated documentation.
 - You'll see two or more copies of model names.  One in model definition and another in
   data node. This is not very 'DRY' and prone to being fragile should the names change.
   In many cases, production code will use reflection, maps or potentially other means to
   keep things in sync.
 - Examples do not get very creative to reuse code when it distracts from communication how
   the API works
 */

import (
	"testing"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"fmt"
)

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
 */
func TestExampleSimplest(t *testing.T) {

	// Model
	model, _ := yang.LoadModuleFromString(nil,

		// namespace, prefix and revision are required as part of YANG spec but empty
		// and zero values are allowed.
		`module x {
			namespace "";
			prefix "";
			revision 0;

			leaf message { type string; }
		}`)

	// Data
	data := &node.MyNode{
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = &node.Value{Str:"Hello"}
			return nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser2(model, data)

	// Read
	msg, _ := brwsr.Root().Get("message")
	fmt.Println(msg)
}

/*
TestReadingMultipleLeafs expands on TestSimplestExample by adding multiple leafs. In the
OnField method, we now need to differentiate which field we want to return

Output:
==============
Mary
 */
func TestExampleReadingMultipleLeafs(t *testing.T) {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			leaf to { type string; }
			leaf message { type string; }
		}`)

	// Data
	data := &node.MyNode{
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "to":
				hnd.Val = &node.Value{Str:"Mary"}
			case "message":
				hnd.Val = &node.Value{Str:"Hello"}
			}
			return nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser2(model, data)

	// Read
	msg, _ := brwsr.Root().Get("to")
	fmt.Println(msg)
}

/*
TestReadingStruct expands on TestSimplestExample by wrapping a 'container' around the
message.  Containers are like a Golang struct.

Output:
==============
Hello
 */
func TestExampleReadingStruct(t *testing.T) {
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
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = &node.Value{Str:"Hello"}
			return nil
		},
	}
	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			return messageData, nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser2(model, data)

	// Read
	msg, _ := brwsr.Root().Find("suggestionBox").Get("message")
	fmt.Println(msg)
}

/*
TestReadingList expands on TestSimplestExample by wrapping a 'list' around the
message.  Lists are like a Golang slices or arrays of structs. NOTE: If you want a list
of strings or ints, then you would just use a leaf-list

Output:
==============
Hello
 */
func TestExampleReadingList(t *testing.T) {
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
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = &node.Value{Str:"Hello"}
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
	brwsr := node.NewBrowser2(model, data)

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
func TestExampleReadWrite(t *testing.T) {

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
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
			if r.Write {
				message = hnd.Val.Str
			} else {
				hnd.Val = &node.Value{Str:message}
			}
			return nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser2(model, data)

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
func TestExampleAddContainer(t *testing.T) {
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
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
			box.message = hnd.Val.Str
			return nil
		},
		OnEvent: func(s node.Selection, e node.Event) error {
			switch e.Type {
			case node.NEW:
				fmt.Println("Finished creating suggestion box")
			}
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
	brwsr := node.NewBrowser2(model, data)

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

func TestExampleAddListItem(t *testing.T) {
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
		OnField:func(r node.FieldRequest, hnd *node.ValueHandle) error {
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
		OnEvent: func(s node.Selection, e node.Event) error {
			switch e.Type {
			case node.NEW:
				fmt.Println("Finished creating suggestion box")
			}
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
	brwsr := node.NewBrowser2(model, data)

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
func TestExampleDelete(t *testing.T) {
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
	box := map[string]interface{} {
		"message" : "hello",
	}
	app := map[string]interface{} {
		"suggestionBox" : box,
		"owner" : map[string]interface{} {
			"name" : "joe",
		},
	}
	boxData := &node.MyNode{
		OnEvent: func(s node.Selection, e node.Event) error {
			switch e.Type {
			case node.DELETE:
				// catch this event is the struct itself needs
				// to know it's being deleted.  In this case of
				fmt.Println("Deleting message", box["message"])
			}
			return nil
		},
	}

	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			return boxData, nil
		},
		OnEvent: func(s node.Selection, e node.Event) error {
			switch e.Type {
			case node.REMOVE_CONTAINER:
				// catch this event for the owner of the struct to remove
				// references to the struct being deleted
				delete(app, e.Src.Meta().GetIdent())
			}
			return nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser2(model, data)

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
func TestExampleAction(t *testing.T) {
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
				result := map[string]interface{} {
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
	brwsr := node.NewBrowser2(model, data)

	// Delete
	result := brwsr.Root().Find("sum").Action(input)
	v, _ := result.GetValue("result")
	fmt.Println(v.Int)
}