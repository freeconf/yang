// This is a very simple example of using the c2g library to add RESTCONF
// support to an app that can say hello to you.
//
// To run this example, use:
//    go run ./hello.go
//
// Then run these commands:
//
//  $>curl http://localhost:8009/restconf/
//     {"count":0}
//
//  $>curl -X PUT -d '{"message":"hello"}' http://localhost:8009/restconf/
//
//  $>curl http://localhost:8009/restconf/
//      {"message":"hello","count":0}
//
//  $>curl -X POST -d '{"name":"joe"}' http://localhost:8009/restconf/say
//      {"message":"hello joe"}
//
//  $>curl http://localhost:8009/restconf/
//      {"message":"hello","count":1}
//

package main

import (
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
)

// This is YANG and usually stored in files but you can store your yang anywhere
// you want including in your source or in a database.  See YANG RFC for full
// options of definitions including typedefs, groupings, containers, lists
// actions, leaf-lists, choices, enumerations and others.
var helloApiDefinition = `
/*
  module is a collection or definitions, much like a YANG container except its the
  top-most container
*/
module hello {
  prefix "hello";
  namespace "hello";
  revision 0;

  /* a "leaf" is just a field */
  leaf message {
  	type string;
  }

  leaf count {
    /* marking fields as NOT config helps denotes them as metrics */
    config "false";
    type int32;
  }

  /* An "rpc" is a function you want to expose */
  rpc say {

    /* rpcs can have optional input defined */
    input {
      leaf name {
        type string;
      }
    }

    /* rpcs can have optional output defined */
    output {
      leaf message {
        type string;
      }
    }
  }
}
`

func main() {
	// Your app, no references to Conf2 enc.
	app := &MyApp{}

	// This is the connection between your app and Conf2.  ManageApp can then
	// navigate through your app's other structures to fulfil API.
	// Here we load from memory, but to load from YANGPATH environment variable use:
	//  yang.LoadModule(yang.YangPath(), "some-module")
	var browser *node.Browser
	if model, err := yang.LoadModuleFromString(nil, helloApiDefinition); err != nil {
		panic(err.Error())
	} else {
		browser = node.NewBrowser(model, management(app))
	}

	// You can register as many APIs as you want, The module name is the default RESTCONF base url
	// Create a RESTCONF service to register your APIs
	service := restconf.NewService(nil, browser)
	service.Port = ":8009"

	// you may want to start in background, but here we start in foreground to keep app running.
	// Hit Ctrl-C in terminal to quit
	service.Listen()
}

// Beginning of your existing application code and has no references to Conf2
type MyApp struct {
	Message string
	Count   int
}

// a random function we'll expose thru API using OnAction below
func (app *MyApp) SayHello(name string) string {

	// as a random metric, let's count how many times we've said hello
	app.Count++

	return app.Message + " " + name
}

// Each unique Go struct you want to manage typically has a corresponding management Node
// but similar structs may share Node definitions.
func management(app *MyApp) node.Node {

	// Node is an interface, but there's a convenient struct that implements the interface
	// and delegates operations to closure functions defined below
	return &node.MyNode{

		// If we had nested Go structs, we'd implement OnSelect to drill into data
		// If we were managing a list of structs, we'd implement OnNext

		// This is for reading data, URL is
		//
		//   GET http://localhost:8009/restconf/hello
		//   Example Response:
		//    {"message":"hello","count":10}
		//
		// This is for writing data, URL is
		//
		//   PUT http://localhost:8009/restconf/hello
		//
		//   Example Payload:
		//     {"message":"hello"}
		//
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			// Here we can use Go's reflection but if reflection isn't valid for some or all fields,
			// you can add a switch case to handle them separately
			if r.Write {
				err = node.WriteField(r.Meta, app, hnd.Val)
			} else {
				hnd.Val, err = node.ReadField(r.Meta, app)
			}
			return
		},

		// Any RPCs (YANG "rpc" or "action") with come thru here
		//
		//  POST http://localhost:8009/restconf/hello/say
		//
		//   Example Payload:
		//     {"name":"joe"}
		///
		//   Example Response:
		//     {"message":"hello joe"}
		//
		OnAction: func(r node.ActionRequest) (out node.Node, err error) {

			// You can use a variety of methods to unmarshal the input including sticking into go map
			// using the Bucket struct
			name, err := r.Input.GetValue("name")
			if err != nil {
				return
			}

			// See how we can call functions of our app from our management data
			// browser?
			s := app.SayHello(name.Str)

			// Build the response, we choose reflection marshaller again just like input data
			out = node.MapNode(map[string]interface{}{"message": s})
			return
		},
	}
}
