[![Build Status](https://travis-ci.org/freeconf/c2g.svg?branch=master)](https://travis-ci.org/freeconf/c2g)

# ![FreeCONF](https://s3.amazonaws.com/freeconf-static/freeconf-no-wrench.svg)
FreeCONF plays an important role in the greater mission to browse, inspect and program __every piece__ of running software in your entire IT infrastructure! 

FreeCONF is a library for adding IETF standards support for configuration, metrics, operations and events to any service written in the Go programming language.


## Why?
Every services needs integration with IT tools that provide:

1. fault monitoring
2. configuration management
3. administration API
4. performance metrics and analysis
5. security

With just **5 different services** you would need to develop and maintain **25** custom integration scripts or plugins. 

![No Standard](https://s3.amazonaws.com/freeconf-static/no-standard.png)

However with a **management standard** there are **NO** integration scripts or plugins to write.  

![With Standard](https://s3.amazonaws.com/freeconf-static/with-standard.png)

## How does it work?
Every running service publishes one or more files called "YANG files" describing their management capabilities.  IT tools can then read these "YANG files" directly from the running service to discover the service's management capabilties.  Once the management capabilties are known, IT tools can manage the running service even though it had no prior knowledge of the service.

For example let's say you wrote a new toaster service and you wanted it to be manageable. 

Steps as a developer:

1. Describe the management capabilities of the toaster in a [YANG file like this one.](https://github.com/YangModels/yang/blob/master/experimental/odp/toaster.yang). 
2. Use FreeCONF library *(or any other library that supports proper server-side IETF RFCs)* to serve YANG files and help developer implement the management capabilties.

Steps as a operator:

1. Start toaster service within your IT infrastruture *(doesn't matter how : docker container, bare metal or physical device)*.
2. Choose an alert service as part of your IT infrastructure *(or write your own with FreeCONF)* that supports proper client-side IETF RFCs.
3. Alert service will read selected services and discover there are two events exported by the toaster service: `toasterOutOfBread` and `toasterRestocked`.  Alert service can then ask any operator which events they'd like to be notified about.

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
	* [Call Home](https://tools.ietf.org/html/rfc8071)
* no dependencies beyond Go Standard Library and Go's `net` package
* code generation optional
* no code annotations (i.e. "go tags") required
* documentation generator
* client and server support including examples

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

	notification update {
		leaf state {
			type enumeration {
				enum outOfGas;
				enum running;
			}
		}
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
		Base: nodes.ReflectChild(car),

		// handle action request
		OnAction: func(parent node.Node, req node.ActionRequest) (node.Node, error) {
			switch req.Meta.Ident() {
			case "start":
				go car.Start()
			}
			return nil, nil
		},
		
		...
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
	
	// Select wire-protocol
	restconf.NewServer(d)
	
	// apply start-up config
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
      '{"restconf":{\
          "web":{"port":":8080"}},\
          "car":{}}'
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

To get updates to the car status, one options is to use websockets from Node.js or the web browser:

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
  * [Car Doc](https://github.com/freeconf/examples/blob/master/car/api/car.md) - Car example generated doc. 
  * [Car Model](https://github.com/freeconf/examples/blob/master/car/api/car.svg) - Graphical representation
  * [RESTConf Doc](https://github.com/freeconf/c2g/blob/master/yang/api/restconf.md) - RESTConf is itself managable.
* [Example YANG files](https://github.com/freeconf/c2g/tree/master/yang) - Used internally by FreeCONF
* [Industry YANG files](https://github.com/openconfig/public/tree/master/release/models) - From openconfig.net project
* [More Industry YANG files](https://www.yangcatalog.org/) - From yangcatalog.org project

## Resources
* [YANG/RESTCONF](https://en.wikipedia.org/wiki/YANG) on wikipedia
* Slides on why we need a DevOps standards from [an operatator's perspective](https://docs.google.com/presentation/d/1q6-kWQI9ahC6iX0EccxLJ32RWv1AUdno8TwX2g9UzYc/edit?usp=sharing) and [a developer's perspective](https://docs.google.com/presentation/d/1g1QLtu7E3acSfeIOH7bc8vZHAULCpgccoQTHRLeczx0/edit?usp=sharing).
* [Manual](https://docs.google.com/document/d/1EMTn8dDsMjOc6f4u0D7kTONQbD2C4hTNFuuqrvXav7o/edit?usp=sharing) - Work in progress
