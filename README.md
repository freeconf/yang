# C2Stack
C2Stack let's you add management capability to any application written in the Go programming language. Manage configuration, metrics, RPCs and events using a RESTful API. 

*IETF Standards : [YANG RFC6020](http://tools.ietf.org/html/rfc6020), [RESTCONF RFC8040](https://tools.ietf.org/html/rfc8040)*

```

custom scripts ----+   +----------+------------+
                   |   |   REST   |  your app  |
 ...and/or...      |   |   mgmt   |            |
                   +==>|          |            |             
RFC8040 compliant  |   |          |            | 
management tools --+   +----------+------------+
                       
```

## Why?
All applications require some sort of management.  By using an industry standard you can offload select management duties to standard compliant tools.  Because the RESTCONF standard is built on-top of REST, users of your management API may not even be aware they are using a standard.

## How does this compare to ___?
Most likely c2stack complements what you're using today for management. There are no agents to install or plugins to build.  The more your infrastructure supports this standard the less integration code you'll need to develop and maintain.

## Source code

To download the source into your project:

 `go get -d -u github.com/ctstack/c2g/...`
 
## Benefits
* no dependencies beyond Go Standard Library
* no code generation or code annotations required
* includes tools to generate documentation
* client and server implementations including examples
* optionally serve and aggregate custom user interfaces
* embraces live configuration changes w/o service restarts
* support for multiple programing languages in roadmap


## License
Licensed under BSD-3-Clause license.

## Getting started

### Step 1. Write your application as you normal would
Here we are implementing a car application.  None of the code will have references to c2stack libraries and therefore reusable outside c2stack and unit testable.

```go
type Car struct {
	FrontSeat   *Seat
	Owner       string
	...
}

type Seat struct {
   ...
}
```

You'll notice you **do not** even have to write code that reads configuration files.

### Step 2. Model your application in YANG
Using your knowledge of your application, you probably have an idea of what data you want configurable, what metrics you want available, what events/alerts you want to communicate and what internal functions you want to make available.  Write a YANG file capturing your model.  You can use the [YANG specification](https://tools.ietf.org/html/rfc6020) as a language guide. YANG is an incredibly powerful data modeling language with readability, extensibility, reusablilty and data validation as core principals.

```YANG
/* 
  Root management model of car application.  Here you describe 
  your management API using:
    container - for data structures 
    list - for lists/arrays of other data structures
    leaf - for fields/properties/primative data types.
    leaf-list - for lists/arrays of fields
    notification - for events
    action - for RPCs
*/
module car {

    /* required in spec, but only useful in interop situations */
	namespace "";
	prefix "";
	
	/* useful for API versioning */
	revision 0;
	
	/* 
	  container is like a data structure.  It can contain the
	  same data modeling things a module can including other containers
	*/
	container frontSeat {
		description "Front seat of the car";
		
		/*  more definition about the front seat ... */
	}	
	
	/*  
	 leaf describes a primative data types including strings, ints,
	 floats, enumerations, bits, leaf-ref (pointer) and even an 
	 "any" type for unstructured data
	 
	 By default, leafs are read/write but you can use 'config "false"' 
	 to mark leaf as read-only useful for operational/metric data like
	 cpu or RAM usage
	*/
	leaf owner { 
	    description "Owner of the car";
	    type string;
	}
	
	/* 
	  other very useful and common constructs include:
	    grouping - to re-use chunks of definitions
	    typedef - to re-use leaf types and data validation
	    choose - to denote one of many possibilities
	    augment - to extend an existing definition
	    many others....
	*/
}
```
 
### Step 3. Write node handlers in C2Stack
Based on your model, you need to write management handlers that map your management model to your application.  

```go
import "github.com/c2stack/c2g/node"

// Your management handling. This method signature is up to you, accept any
// parameters you need to implement management on a car
func carNode(car *Car) node.Node {

  // MyNode is for complete custom management handling, but many other starter
  // management handlers exist including: ReflectNode, Extend, MapNode,
  // ListNode.  Every handler you create is reusable and extendable by default.
  return &node.MyNode{
  
    // Handle management for data structures (e.g. 'container') which normally
    // means data structure construction and then delegating managment to 
    // other handlers
    OnChild : func(r node.ChildRequest) (node.Node, error) {
       switch r.Meta.GetIdent() {
         case "seat":
            // ...            
    }
    
    // Read/Write fields.  
    OnField : func(r node.FieldRequest, val *node.ValueHandle) error {
       switch r.Meta.GetIdent() {
         case "owner":
    	      // ...
    } 
    
    // Other optional entry points include:
    //  OnNext - navigating thru a list
    //  OnAction - RPC implementation
    //  OnNotify - Event implementation
    //  OnBeginEdit - Know when a data structure is being edited
    //  OnEndEdit - Know when a data structure is done being edited
    //  OnDelete - Know then a data structre is being deleted
    //  OnPeek - Gain access to undelying data structres being managed
    //  OnContext - Add to ansilary data about a request such as auth creds
  }
}
```

### Step 4. Connect everything together
How you connect everything together is up to you.  Here's just one example:

```go
import (
	"github.com/c2stack/c2g/restconf"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/device"
)

func main() {

	// Instantiate your app
	car := &Car{}
		
	// Here we use YANGPATH environment variable to point to 
	// YANG files. You may decide load them differently.
	modelSrc = yang.YangPath()
	
	// Your app's root management handler from Step 2
	mgmt := carNode(car)
	
	// Create a container for your management services. 
	d := device.New(modelSrc)
	
	// Register your management implementation. You can
	// register as many modules as you wish no more than one per module type
	d.Add("car", mgmt)
		
	// Pick RESTCONF as management protocol
	restconf.NewServer(d)
	
	// This helper function bootstraps config from local json file.
	d.ApplyStartupConfigFile("startup.json")
	
	// Before running your app, be sure to set environment variable
	//  YANGPATH=path/to/car.yang/file
	//
	// To get configuration and operational data
	//  curl http://localhost:8080/restconf/data/car:
}
```

## More Examples
* [App Examples](https://github.com/c2stack/c2g/blob/master/examples) - Complete applications that each have management APIs.
* [Code Examples](https://godoc.org/github.com/c2stack/c2g/examples) - Mostly examples on management node handlers options.
* Example generated docs. Templates exist for Markdown, HTML and SVG (thru dot)
  * [Car Doc](https://github.com/c2stack/c2g/blob/master/examples/car/car.md) - Car example generated doc. 
  * [Car Model](https://github.com/c2stack/c2g/blob/master/examples/car/car.svg) - Graphical representation
  * [RESTConf Doc](https://github.com/c2stack/c2g/blob/master/yang/restconf.md) - RESTConf is itself managable.  
* [Example YANG files](https://github.com/c2stack/c2g/tree/master/yang) - Used internally by C2Stack
* [Industry YANG files](https://github.com/openconfig/public/tree/master/release/models) - From openconfig.net project

## Resources
* [Questions or problems](https://stackoverflow.com/questions/ask?tags=c2stack) - StackOverflow
* [State of server-sode development](https://docs.google.com/presentation/d/1g1QLtu7E3acSfeIOH7bc8vZHAULCpgccoQTHRLeczx0/edit?usp=sharing) - Slides on pitching why we need a standard like YANG and RESTCONF.
* [User Manual](https://docs.google.com/document/d/1EMTn8dDsMjOc6f4u0D7kTONQbD2C4hTNFuuqrvXav7o/edit?usp=sharing) - Work in progress