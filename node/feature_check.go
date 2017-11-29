package node

import (
	"github.com/freeconf/c2g/meta"
)

type FeatureCheck struct {
}

func (y FeatureCheck) CheckContainerPreConstraints(r *ChildRequest) (bool, error) {
	return y.checkFeature(r.Selection, r.Meta)
}

func (y FeatureCheck) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	return y.checkFeature(r.Selection, r.Meta)
}

func (y FeatureCheck) CheckListPreConstraints(r *ListRequest) (bool, error) {
	return y.checkFeature(r.Selection, r.Meta)
}

func (y FeatureCheck) CheckActionPreConstraints(r *ActionRequest) (bool, error) {
	return y.checkFeature(r.Selection, r.Meta)
}

func (y FeatureCheck) checkFeature(s Selection, m meta.Meta) (bool, error) {
	if s.Browser.Features != nil {
		for _, f := range m.(meta.HasIfFeatures).IfFeatures() {
			return s.Browser.Features.FeatureOn(f)
		}
	}
	return true, nil
}
