package node

import (
	"context"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Node is at the root of where your application's data management and functionality is
// contained and made available.
//
// When building your application management logic, you will register a single instance
// of a node representing the root of the module that will lead to all of nodes, functions
// and event streams. Each of these nodes, functions and event streams will align with the
// definitions in the YANG file.
//
// You can use data structures nodeutil.Basic, nodeutil.Extend or nodeutil.Reflect that implement
// this interface to reduce the amount of required coding.
//
// The root Node together with the meta.Module represent the top most node and can be obtained
// using node.Browser's Root() function to return a node.Selection data structure that links the
// Node and the Module.  This root Selection then be used to navigate you manament application
// thru a series of nested Selections of meta.Meta and node.Node pairings.
// See https://freeconf.org/docs/reference/basic-components/ for how Node works with other objects
//
// Node are constructed as needed so two separate requests to the same management function would
// have different Node instances.  This means you can keep temporary state like
// lazy list indexes or references to  request specific data in your node implementation.  The
// only time a Node is not garage collected after a request
// is finished is if it is Root node or there is an active event stream via a Notification
// subscription.
type Node interface {

	// Child is called when navigating, creating or deleting a "container" or "list".
	// The request fields contain the nature of the request.
	//
	// Params:
	//  ChildRequest - contains the nature (e.g. get, delete, new) and details of the
	//                 request (e.g. identity name of the container or list )
	//
	// Return:
	//   Node - Return nil entity does not exist otherwise node implementation of the underlying
	//          node requested
	Child(r ChildRequest) (child Node, err error)

	// Next is called for items in a list when navigating, creating or deleting items in the
	// "list".
	//
	// Params:
	//  ListRequest - contains the nature (e.g. get, delete, new) and details of the
	//                 request (e.g. key, row number)
	//
	// Return:
	//   Node - Return nil entity does not exist otherwise node implementation of the underlying
	//          node requested
	//   []val.Value - If a key was defined in YANG, AND the request is for the next item
	//           in the list, then you must also return the key expressed in val.Value array
	Next(r ListRequest) (next Node, key []val.Value, err error)

	// Field is called to read or write "leaf" or "leaf-list" items on container/list items or
	// even root module.
	//
	// Params:
	//   FieldRequest - identity of the field and the contains the nature (e.g. read or write)
	//             tip: if the YANG defines the field as config:false, you can assume the field is
	//             only called for reading and you can ignore the flag for read v.s. write
	//
	//   *ValueHandle - contains a single the "Val" that with either:
	//                   a.) contain the value to be set or
	//                   b.) be expected to be set to the requested value to be read
	Field(r FieldRequest, hnd *ValueHandle) error

	// Choose is only called when there is a need to know which single case between several "choice/case"
	// sets are currently true.  This is only called when reading.
	//
	// Params:
	//   Selection - current selection that is handling the choice investigation
	//   *Choice - name for the choice (there could be several in a given container)
	//
	// Return:
	//   *ChoiceCase - which case is currently valid (there can be only one)
	//
	Choose(sel Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)

	// BeginEdit is called simply to inform a node when it is about to be edited including deleting or creating.
	// While there is nothing required of nodes to do anything on this call, implementations might have
	// very important, internal operations they need to perform like obtaining write locks for example. It is
	// also an opportunity to reject the edit for whatever reason by returning an error.
	//
	// It important to know this is called on **every** parent node for any edit to any of their children.
	// This is known as "event bubbling" and can be helpful when parents are in a better position to handle
	// operations for children.  You can easily distinguish this situation from the NodeRequest parameter.
	// See https://github.com/freeconf/restconf/wiki/Edit-Node-Traversal for details on when this is called
	// in particular in the order
	//
	// Params:
	//  NodeRequest - contains the nature (e.g. get, delete, new)
	BeginEdit(r NodeRequest) error

	// EndEdit is called simply to inform a node after it has been edited including deleting or creating.
	// While there is nothing required of nodes to do anything on this call, implementations might have
	// very important, internal operations they need to perform like release write locks for example. It is
	// also an opportunity to apply changes to live application or persist data to permanent storage.
	//
	// It important to know this is called on **every** parent node for any edit to any of their children.
	// This is known as "event bubbling" and can be helpful when parents are in a better position to handle
	// operations for children.  You can easily distinguish this situation from the NodeRequest parameter.
	// See https://github.com/freeconf/restconf/wiki/Edit-Node-Traversal for details on when this is called
	// in particular in the order
	//
	// Params:
	//  NodeRequest - contains the nature (e.g. get, delete, new)
	EndEdit(r NodeRequest) error

	// Action(rpc/action) is called when caller wished to run a 'action' or 'rpc' definition.  Input can
	// be found in request if an input is defined.  Output only has to be returned for
	// definitions that declare an output.
	//
	// Params:
	//  ActionRequest - In addition to the identity name of the action or rpc, it may contains the "input"
	//        to the action if one was defined in the YANG.
	//
	// Return:
	//  Node - If there is an expected response from action ("output" defined in YANG) then this would be
	//        the Node implementation of that data.
	Action(r ActionRequest) (output Node, err error)

	// Notify(notification) is called when caller wish to subscribe to events from a node.

	// Notificationsare unique with respect to resources are allocated that survive the original
	// request so implementation should try not keep references to any more resources than neccessary
	// to ensure they do not reference objects that become stale or that impose large chunks of memory
	// unknowningly and unneccesarily.
	//
	// Params:
	//   NotifyRequest - In addition to the identity name of the notification, this contains the stream
	//        to write events to.
	//
	// Return:
	//   NotifyCloser - Function to called to stop streaming events. This will be called for example when
	//      client closes the event stream socket either naturally or because of a network disconnect.
	//
	Notify(r NotifyRequest) (NotifyCloser, error)

	// Peek give an opportunity for API callers to gain access to the objects behind a node.
	//
	// Params:
	//   Selection - current selection that is handling the peeking
	//   interface{} - Who (or under what context) is attempting to peek. This can be used by
	//         node implementation to decide what is returned or if anything at all is returned.
	//         It may also be ingored.  This is up to
	//
	// Return:
	//   interface{} - Can be anything.
	//
	// This can be considered a violation of abstraction so implementation is not gauranteed and
	// can vary w/o any compiler warnings.  This can be useful in unit testing but uses outside this
	// should be used judicously.
	Peek(sel Selection, consumer interface{}) interface{}

	// Context provides an opportunity to add/change values in the request context that is passed to
	// operations on this node and operations to children of this node.  FreeCONF does not put or
	// require anything in the context, it is meant as a way to make application or request specific
	// data available to nodes.
	//
	// A popular use of context is to store the user and or user roles making the request so
	// operations can authorize or log user operations.  See RESTCONF request filters for one
	// way to implement that.
	Context(sel Selection) context.Context
}

// Used to pass values in/out of calls to Node.Field
type ValueHandle struct {

	// Readers do not set this, Writers will always have a valid value here
	Val val.Value
}
