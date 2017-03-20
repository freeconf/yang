package node

// FastForwardingNode will always claim to have children for requests that follow
// a specific target path. Kinds like fast forwarding to a specific spot.  Useful
// in a variety of cases including proxying.
func FastForwardingNode() Node {
	e := &MyNode{}
	e.OnChild = func(r ChildRequest) (Node, error) {
		return e, nil
	}
	e.OnNext = func(r ListRequest) (Node, []*Value, error) {
		return e, r.Key, nil
	}
	return e
}
