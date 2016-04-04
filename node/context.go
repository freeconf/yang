package node

import "github.com/c2g/meta"

type Context interface {
	Handler() *ContextHandler
	Constraints() *Constraints
	Select(m meta.MetaList, node Node) Selector
	Selector(s *Selection) Selector
}
