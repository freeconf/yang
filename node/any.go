package node

import (
	"encoding/json"
	"bytes"
	"strings"
	"io"
	"io/ioutil"
)

type AnyData interface {
	String(c *Context) (string, error)
	Node() Node
}

type AnyNode struct {
	TheNode Node
}

func (self AnyNode) Node() Node {
	return self.TheNode
}

func (self AnyNode) String(c *Context) string {
	panic("?")
}

type AnyReader struct {
	Reader io.Reader
}

func (any *AnyReader) Node() Node {
	return NewJsonReader(any.Reader).Node()
}

func (any *AnyReader) String(c *Context) (string, error) {
	b, err := ioutil.ReadAll(any.Reader)
	return string(b), err
}

type AnyJsonString struct {
	Json string
}

func (any *AnyJsonString) Node() Node {
	return NewJsonReader(strings.NewReader(any.Json)).Node()
}

func (any *AnyJsonString) String(c *Context) (string, error) {
	return any.Json, nil
}

type AnyJson struct {
	container map[string]interface{}
}

func (any *AnyJson) Node() Node {
	return JsonContainerReader(any.container)
}

func (any *AnyJson) String(c *Context) (string, error) {
	bytes, err := json.Marshal(any.container)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type AnySelection struct {
	Selection *Selection
}

func (any *AnySelection) Node() Node {
	return any.Selection.Node()
}

func (any *AnySelection) String(c *Context) (string, error) {
	var out bytes.Buffer
	e := c.Selector(any.Selection).InsertInto(NewJsonWriter(&out).Node()).LastErr
	return out.String(), e
}

