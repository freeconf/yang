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
	msg, _ := brwsr.Root().Selector().Get("message")
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
	msg, _ := brwsr.Root().Selector().Get("to")
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
	msg, _ := brwsr.Root().Selector().Find("suggestionBox").Get("message")
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
	msg, _ := brwsr.Root().Selector().Find("suggestionBox=joe").Get("message")
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

		// namespace, prefix and revision are required as part of YANG spec but empty
		// and zero values are allowed.
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
	msg, _ := brwsr.Root().Selector().Get("message")
	fmt.Println(msg)
	_ = brwsr.Root().Selector().Set("message", "Goodbye")
	msg, _ = brwsr.Root().Selector().Get("message")
	fmt.Println(msg)
}
