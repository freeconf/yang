# ![FreeCONF](https://s3.amazonaws.com/freeconf-static/freeconf-no-wrench.svg)

<font size="3">[Add support or configuration, metrics, alerts and management functions to your application!](https://freeconf.org)</font>

# In this repository

*  [IETF YANG RFC7950](https://tools.ietf.org/html/rfc7950)  file parser which can be used to parse YANG files into complete AST without the loss of data including YANG language extensions.  See compliance for all [supporting RFCs](https://freeconf.org/docs/reference/compliance/rfcs/).
*  Core logic to build a management API in Go.  You would need to include either [FreeCONF's RESTCONF](https://github.com/freeconf/restconf) or [FreeCONF's gNMI](https://github.com/freeconf/gnmi) projects to make management API available over a network interface.
*  Management API documentation generator from YANG files

# Requirements

Requires Go version 1.20 or greater.

# Getting the source

```bash
go get -u github.com/freeconf/yang
```

# Resources
* [Web site](https://freeconf.org)
* [Documentation](https://freeconf.org/docs)
  * [Getting Started](https://freeconf.org/docs/gettingstarted/)
  * [Next Steps](https://freeconf.org/docs/examples/next-step/)
  * [Generating Documentation](https://freeconf.org/docs/reference/docs/) from YANG files
  * [RFC Compliance](https://freeconf.org/docs/reference/compliance/rfcs/)
* [Go API Docs](https://pkg.go.dev/github.com/freeconf/yang)
* [Issues](https://github.com/freeconf/yang/issues)
* [Discussions](https://github.com/freeconf/restconf/discussions)
