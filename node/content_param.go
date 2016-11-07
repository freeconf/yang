package node

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

type ContentConstraint int

const (
	ContentAll ContentConstraint = iota
	ContentOperational
	ContentConfig
)

func NewContentConstraint(initialPath *Path, expression string) (c ContentConstraint, err error) {
	switch expression {
	case "config":
		return ContentConfig, nil
	case "nonconfig":
		return ContentOperational, nil
	case "all":
		return ContentAll, nil
	}
	return ContentAll, c2.NewErrC("Invalid content contraint: "+expression, 400)
}

func (self ContentConstraint) CheckContainerPreConstraints(r *ContainerRequest) (bool, error) {
	// config containers may have operational fields so always pass on operational
	if r.IsNavigation() || self == ContentAll || self == ContentOperational {
		return true, nil
	}

	var isConfig bool
	// meta.Module does not implement HasDetails, but spec implies yes
	if d, hasDets := r.Meta.(meta.HasDetails); !hasDets {
		isConfig = true
	} else {
		isConfig = d.Details().Config(r.Selection.Path)
	}
	return isConfig, nil
}

func (self ContentConstraint) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	if r.IsNavigation() || self == ContentAll {
		return true, nil
	}
	isConfig := r.Meta.(meta.HasDetails).Details().Config(r.Selection.Path)
	return (isConfig && self == ContentConfig) || (!isConfig && self == ContentOperational), nil
}
