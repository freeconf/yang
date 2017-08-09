[![Build Status](https://travis-ci.org/c2stack/c2g.svg?branch=master)](https://travis-ci.org/c2stack/c2g)

# C2Stack
C2Stack let's you manage configuration, metrics, operations and events for any microservice written in the Go programming language.


```

       +---------+---------+      config
       |  your   |  REST   |      metrics
       | service |         | <=>  operations
       |         |         |      alerts (over websockets)
       +---------+---------+
                       
```

## Why?
All applications require some sort of management.

## How does this compare to ___?
Most likely c2stack complements what you're using today for management. There are no agents to install, plugins to build or servers to start.

## Source code
To download the source into your project:

`go get -d -u github.com/c2stack/c2g/...`
 
## Benefits
* Supports IETF management standards [YANG](http://tools.ietf.org/html/rfc6020), and [RESTCONF](https://tools.ietf.org/html/rfc8040)
* no dependencies beyond Go Standard Library
* no code generation or code annotations required
* includes tools to generate documentation
* client and server implementations including examples
* enables live configuration changes w/o service restarts


## License
Licensed under BSD-3-Clause license.

## Getting started
Full source for this example is [here](https://github.com/c2stack/c2g/tree/master/examples/car).

### Step 1. Write your application as you normal would
Here we are implementing a car application.  

```go
type Car struct {
	Speed     int
	Miles     int64
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
	   config "false";
	}	    	    
}
```
 
### Step 3. Connect everything together
```go
import (
	"github.com/c2stack/c2g/restconf"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/device"
)

func main() {

	// Your app
	car := &Car{}
		
	// Add management
	d := device.New(yang.YangPath())
	// use Go reflection, but many other options exist
	d.Add("car", nodes.Reflect(car)) 
	restconf.NewServer(d)
		
	// trick to sleep forever...
	select {}
}
```

### Step 4. Using your management API
Start your application

`YANGPATH=. go run ./main.go`

#### Get Configuration
`curl http://localhost:8080/restconf/data/car:?content=config`

```json
{
  "speed": 100
}
```

#### Change Configuration
`curl -XPUT -d @- http://localhost:8080/restconf/data/car: <<< '{"speed":99}'`


#### Metrics
`curl http://localhost:8080/restconf/data/car:?content=nonconfig`

```json
{
  "miles": 133
}
```

#### Alerts
`curl` doesn't support websockets, so we'll write a little node app.

```JavaScript
var ws = require('ws');
var notify = require('./notify');
var driver = new ws('ws://localhost:8080/restconf/streams','', {origin:'localhost:8080'});
var n = new notify.handler(ws_driver);
n.on('', 'update', 'car', (car, err) => {
  console.log(car);
});
```

`node my-app.js`

```json
{
  "speed": 99,
  "miles": 253
}
```

## Security
Configuration settings exist for server certifications and client certification authoritation.  You can implement fine grained control of any management operation based on whatever authentication management you decide.

## More Examples
* [App Examples](https://github.com/c2stack/c2g/blob/master/examples) - Complete applications that each have management APIs.
* [Code Examples](https://godoc.org/github.com/c2stack/c2g/examples) - Mostly examples on management node handlers options.
* Example generated docs. Templates exist for Markdown, HTML and SVG (thru dot)
  * [Car Doc](https://github.com/c2stack/c2g/blob/master/examples/car/api/car.md) - Car example generated doc. 
  * [Car Model](https://github.com/c2stack/c2g/blob/master/examples/car/api/car.svg) - Graphical representation
  * [RESTConf Doc](https://github.com/c2stack/c2g/blob/master/yang/api/restconf.md) - RESTConf is itself managable.
* [Example YANG files](https://github.com/c2stack/c2g/tree/master/yang) - Used internally by C2Stack
* [Industry YANG files](https://github.com/openconfig/public/tree/master/release/models) - From openconfig.net project

## Resources
* [Go API](https://godoc.org/github.com/c2stack/c2g)
* [YANG/RESTCONF](https://en.wikipedia.org/wiki/YANG) on wikipedia
* [Slides](https://docs.google.com/presentation/d/1g1QLtu7E3acSfeIOH7bc8vZHAULCpgccoQTHRLeczx0/edit?usp=sharing) on why we need a standard like YANG and RESTCONF for microservices.
* [Manual](https://docs.google.com/document/d/1EMTn8dDsMjOc6f4u0D7kTONQbD2C4hTNFuuqrvXav7o/edit?usp=sharing) - Work in progress
