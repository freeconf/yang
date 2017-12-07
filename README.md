[![Build Status](https://travis-ci.org/freeconf/c2g.svg?branch=master)](https://travis-ci.org/freeconf/c2g)

# ![FreeCONF](https://s3.amazonaws.com/freeconf-static/freeconf.svg)
FreeCONF let's you manage configuration, metrics, operations and events for any microservice written in the Go programming language.


```

       +---------+---------+      config
       |  your   |  REST   |      metrics
       | service |         | <=>  operations
       |         |         |      alerts (over websockets)
       +---------+---------+
                       
```

## Why?
Close to half of all software engineers are working on infrastructure related development in some capacity.  This percentage is predicted to grow as microservices start gaining traction.  Unfortunately, most of this infrastructure development is custom and not reusable.

Infrastructure and cloud management tools are vital to an organization's productivity but the cost of integrating these tools can bring diminishing returns for many projects.  Incompatible tools reduce each tool's effectiveness and scope.  Conversely, compatible tools bring exponential value by enabling tools to focus on particular problems while allowing the combination of tools to match the right solution to the right problem.

In February 2014 the IETF standards organization introduced a proposal for network management using a REST based protocol.  This proposal would allow network management tools to manage a network using services from multiple vendors. In 2015, Douglas Hubler, recognized the quality of the proposal and it's value as a standard for microservice development and cloud management so he created a library that would enable these standards for any microservice or infrastructure tool.  In January 2017, the IETF published the RESTCONF specification.

## How does this compare to ___?
Most likely FreeCONF complements what you're using today for management. There are no agents to install, plugins to build or servers to start.

## Source code
Requires Go version 1.9 or greater and uses [go dep](https://github.com/golang/dep) to manage the one dependencies on `golang.org/x/net`.

If you just want to quickly download the source into your project, you can use this command:

`go get -d -u github.com/freeconf/c2g/...`
 
## Benefits
* Supports IETF management standards:
	* [YANG](http://tools.ietf.org/html/rfc6020)
	* [YANG 1.1](https://tools.ietf.org/html/rfc7950)
	* [RESTCONF](https://tools.ietf.org/html/rfc8040)
* no dependencies beyond Go Standard Library and Go's `net` package
* enables code generation but not required
* no code annotations (i.e. "go tags") required
* includes tools to generate documentation
* client and server implementations including examples
* enables live configuration changes w/o service restarts


## License
Licensed under Apache 2.0 license.

## Getting started
Full source for this example is [here](https://github.com/freeconf/examples/tree/master/src/intro).

### Step 1. Write your application as you normally would
Here we are implementing a car application.  

```go
type Car struct {
	Speed     int
	Miles     int64
}

func (c *Car) Start() {
	for {
		<-time.After(time.Duration(c.Speed) * time.Millisecond)
		c.Miles += 1
	}
}
```

### Step 2. Model your application in YANG
Use [YANG](https://tools.ietf.org/html/rfc6020) to model your management API.

```YANG
module car {

	revision 0;
	
	leaf speed {
	   type int32;
	}	    	    

	leaf miles {
	   type int64;
	   config false;
	}	    	    
	
	rpc start {}
}
```

### Step 3. Add Management

```go
// implement your mangement api
func manage(car *Car) node.Node {
	return &nodes.Extend {
	
		// use reflect when possible, here we're using to get/set speed AND
		// to read miles metrics.
		Base: nodes.Reflect(car),

		// handle action request
		OnAction: func(parent node.Node, req node.ActionRequest) (node.Node, error) {
			switch req.Meta.Ident() {
			case "start":
				go car.Start()
			}
			return nil, nil
		},
	}
}
```
 
### Step 4. Connect everything together
```go
import (
	"github.com/freeconf/c2g/restconf"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/nodes"
	"github.com/freeconf/c2g/device"
)

func main() {

	// Your app
	car := &Car{}
		
	// Add management
	d := device.New(yang.YangPath())
	d.Add("car", manage(car)) 
	restconf.NewServer(d)
	d.ApplyStartupConfig(os.Stdin)
		
	// trick to sleep forever...
	select {}
}
```

### Step 5. Using your management API

Start your application

```bash
YANGPATH=.:../../ \
    go run ./main.go <<< \
    '{"restconf":{"web":{"port":":8080"}},"car":{}}'
```

#### Get Configuration
`curl http://localhost:8080/restconf/data/car:?content=config`

```json
{"speed":100}
```

#### Change Configuration
`curl -XPUT -d @- http://localhost:8080/restconf/data/car: <<< '{"speed":99}'`


#### Metrics
`curl http://localhost:8080/restconf/data/car:?content=nonconfig`

```json
{"miles":133}
```


#### Operations
Start has no input or output defined, so simple POST will start the car

`curl -XPOST http://localhost:8080/restconf/data/car:start`

#### Alerts

This car example doesn't have alerts, but to get alerts we can use websockets from Node.js or the web browser:

```JavaScript
var notify = require('./notify');
...
var events = new notify.handler(ws_driver);
events.on('', 'update', 'car', (car, err) => {
  console.log(car);
});
```

## Security
Default authenication is certificate based and default authorization is based on the YANG model from __Step 2.__. of any management operation based on whatever authentication management you decide.  Each configuration change is logged by the server.


## More Examples
* [Robotic Bartender](https://github.com/dhubler/bartend) - Pour drinks automatically from Raspberry Pi
* [App Examples](https://github.com/freeconf/examples) - Complete applications that each have management APIs.
* [Code Examples](https://godoc.org/github.com/freeconf/examples) - Mostly examples on management node handlers options.
* Example generated docs. Templates exist for Markdown, HTML and SVG (thru dot)
  * [Car Doc](https://github.com/freeconf/examples/blob/master/src/car/api/car.md) - Car example generated doc. 
  * [Car Model](https://github.com/freeconf/examples/blob/master/src/car/api/car.svg) - Graphical representation
  * [RESTConf Doc](https://github.com/freeconf/c2g/blob/master/yang/api/restconf.md) - RESTConf is itself managable.
* [Example YANG files](https://github.com/freeconf/c2g/tree/master/yang) - Used internally by FreeCONF
* [Industry YANG files](https://github.com/openconfig/public/tree/master/release/models) - From openconfig.net project

## Resources
* [YANG/RESTCONF](https://en.wikipedia.org/wiki/YANG) on wikipedia
* [Slides](https://docs.google.com/presentation/d/1g1QLtu7E3acSfeIOH7bc8vZHAULCpgccoQTHRLeczx0/edit?usp=sharing) on why we need a standard like YANG and RESTCONF for microservices.
* [Manual](https://docs.google.com/document/d/1EMTn8dDsMjOc6f4u0D7kTONQbD2C4hTNFuuqrvXav7o/edit?usp=sharing) - Work in progress
