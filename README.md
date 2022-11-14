# ![FreeCONF](https://s3.amazonaws.com/freeconf-static/freeconf-no-wrench.svg)

For more information about this project, [see wiki](https://github.com/freeconf/restconf/wiki).

# YANG library parser

This library parses YANG files into Go structures (AST).  YANG files are management interface definition files following the [IETF standard RFC7950](https://datatracker.ietf.org/doc/html/rfc7950) and used to describe the management API of a given service.

# FreeCONF Mission

FreeCONF plays an important role in the greater mission to browse, inspect and program __every piece__ of running software in your entire IT infrastructure! FreeCONF uses IETF standards to support configuration, metrics, operations and events to any service written in the Go programming language.

# What can I do with this library?

1. Use it together with the [FreeCONF RESTCONF library](https://github.com/freeconf/restconf) to add management API to any service written in the Go language.
2. Parse YANG files into Go data structures (or AST) including resolving YANG features such as `typedefs`, `imports`, `includes` and `groupings`.
3. Generate documentation for YANG files in HTML, SVG diagrams and Markdown formats
4. Generate documentation using your own templates
5. Generate source code from YANG files
6. Validate YANG files for correct syntax
7. Implement other IETF standards protocols like [NETCONF RFC 6241](https://datatracker.ietf.org/doc/html/rfc6241) much like the [FreeCONF RESTCONF library](https://github.com/freeconf/restconf) has done.

## Requirements

Requires Go version 1.9 or greater.

## Getting the source

```bash
go get -u github.com/freeconf/yang
```

## Example parsing YANG files

### Step 1. Create empty project

We'll create a new project involving a car:

```bash
mkdir car
cd car
go mod init car
go get -u github.com/freeconf/yang
```

### Step 2. Create a YANG file

Feel free to use your own YANG files, but if you don't any here is one to get you started.  You must name this file `car.yang` so YANG parser can find the file on the file system.

```yang
module car {
	description "Car goes beep beep";

	revision 0;

	leaf speed {
		description "How fast the car goes";
	    type int32 {
		    range "0..120";
	    }
		units milesPerSecond;
	}

	leaf miles {
		description "How many miles has car moved";
	    type decimal64;
	    config false;
	}

	rpc reset {
		description "Reset the odometer";
	}
}
```

### Step 3. Create a simple program

```go
package main

import (
	"fmt"

	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

func main() {

	// One of the many ways to find *.yang files. This one allows for multiple
	// directories separated with ':' but here we are just using the current working
	// directory
	ypath := source.Path(".")

	// Parse the file into Go structures and report any errors.
	car, err := parser.LoadModule(ypath, "car")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded %s module successfully\n", car.Ident())
}
```

### Step 4. Run your program

```bash
go run .
Loaded car module successfully
```

## Generating documentation from YANG files

Continuing with the car example above, let's generate documentation for our `car.yang` file.

## Step 1. Get internal FreeCONF YANG files

FreeCONF needs a set of common YANG files in order to generate documentation. Lets extract the files from the code base and output them to the current directory

```bash
go run github.com/freeconf/yang/cmd/fc-yang get -dir '.'
```

you should now see bunch of *.yang files in the current directory.  They were actually extracted from the source, not downloaded. 

## Step 2. Generate docs

```bash
YANGPATH=. go run github.com/freeconf/yang/cmd/fc-yang doc -module car -f html > car.html
```

optionally, generate diagram.  You will need to [install Graphviz first](https://graphviz.org/download/), then run the following

```bash
YANGPATH=. go run github.com/freeconf/yang/cmd/fc-yang doc -module car -f dot > car.dot
dot -Tsvg car.dot -o car.svg
```

For the minimalists, there is also markdown format using `-f md`

# Generating code, alternative docs, graphql schema, ...

You have two good options. First option is to parse the YANG from Go and walk the tree and generate output that way.  Second option is to convert the yang file to JSON format and then feed that JSON to a scripting tool like jinja and generate files that way.  The JSON format is lossless, so you would have all the same information.

```bash
YANGPATH=. go run github.com/freeconf/yang/cmd/fc-yang doc -module car > car.json
jinja -d car.json my-template.j2 > my-result.dat
```

## Resources
* [Wiki](https://github.com/freeconf/restconf/wiki) - *Combined wiki with FreeCONF RESTCONF project*
* [Discussions](https://github.com/freeconf/restconf/discussions)
* [Issues](https://github.com/freeconf/yang/issues)

## RFCs

If you don't see an RFC here, open a discussion to see if there is interest or existing implementations.

* [RFC 7950](https://datatracker.ietf.org/doc/html/rfc7950) - YANG 1.1
* [RFC 7951](https://datatracker.ietf.org/doc/html/rfc7951) - JSON encoding