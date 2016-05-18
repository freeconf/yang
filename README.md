# Conf2
Conf2 is a library allowing you add a standards compliant management API any running microservice or device. Conf2 implements the following two standards:

* [YANG (RFC6020)](http://tools.ietf.org/html/rfc6020) - A data modeling language that can be used to model configuration, metrics, RPCs and events.
* [RESTCONF (RFC pending)](https://tools.ietf.org/html/draft-ietf-netconf-restconf-07) - A RESTful oriented protocol.

## Benefits:
* Standards compliance means automatic infrastrucure integration with other standards based controller systems.
* Receiving configuration through the network obviates file-based configuration tools such as Puppet or Chef
* Exporting health and metrics data through the network obviates log scraping tools like Splunk
* Sending alerts as they happen to subscribed systems obviates poll-based systems like watchdog or Nagios.
* Exporting RPC through the network obviates tools like Ansible.
* Written in the Go with a C-compatible API enabling support for Java, PHP, Python, Ruby, JavaScript, C/C++ and others.
* Experimental support for Java
* No dependencies beyond Go Standard Library
* API designed to integrate into any existing codebase without modification.
* No code generation
* Access to meta-data for model-driven UIs and tools
* Ability to add custom protocols beyond RESTCONF including NETCONF, SNMP, Weave or others
* Ability to add custom formats beyond JSON or XML
* Experimental support for distributed UI using Web Components (emerging W3C standard).

## Setup
Conf2 requires Go 1.5 for shared library support.

    mkdir -p $GOPATH/src/github.com
    cd $GOPATH/src/github.com
    git clone https://github.com/dhubler/conf2.git
    cd conf2
    make build

## Code Examples:
* [HelloGo](https://github.com/dhubler/conf2-examples/blob/master/helloGo/hello.go) - Basic Go application that says hello to you
* [Todo](https://github.com/dhubler/conf2-examples/tree/master/todo) - Todo Go application

