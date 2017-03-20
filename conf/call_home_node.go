package conf

import "github.com/c2stack/c2g/node"

func CallHomeNode(ch *CallHome) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(ch),
	}
}
