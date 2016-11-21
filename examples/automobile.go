package examples

import "github.com/c2stack/c2g/node"

type auto struct {
	FrontBrake brake
	RearBrake  brake
	Tires      []tire
	CupHolders map[string]*cupHolder
	Engine     *engine
}

type quadrant int

const (
	frontPassenger quadrant = iota
	frontDriver
	rearPassenger
	rearDriver
)

type cupHolder struct {
	Capacity int
}

type engine struct {
	Specs map[string]interface{}
}

type tire struct {
	Size string
}

func (self tire) pressure() int {
	return 32
}

type brake struct {
	Abs   bool
	Style brakeStyle
}

type brakeStyle int

const (
	drum brakeStyle = 1 + iota
	disc
)

func engineNode(engine *engine) node.Node {
	return nil
}

func cupHolderNode(*cupHolder) node.Node {
	return nil
}
