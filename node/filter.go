package node

import (
	"github.com/c2stack/c2g/xpath"
)

func NewFilterConstraint(filter string) (NotifyFilterConstraint, error) {
	p, err := xpath.Parse(filter)
	if err != nil {
		return nil, err
	}
	return xpathFilter{p: p}, nil
}

type xpathFilter struct {
	p xpath.Path
}

func (f xpathFilter) CheckNotifyFilterConstraints(msg Selection) (bool, error) {
	return msg.XPredicate(f.p)
}
