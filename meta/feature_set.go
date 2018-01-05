package meta

import (
	"github.com/freeconf/gconf/c2"
)

type FeatureSet interface {
	Resolve(m *Module) error
	FeatureOn(*IfFeature) (bool, error)
}

func AllFeatures() FeatureSet {
	return &supportedFeatures{specified: []string{}, blacklist: true}
}

func BlacklistFeatures(features []string) FeatureSet {
	return &supportedFeatures{specified: features, blacklist: true}
}

func StrictBlacklistFeatures(features []string) FeatureSet {
	return &supportedFeatures{specified: features, blacklist: true, strict: true}
}

func WhitelistFeatures(features []string) FeatureSet {
	return &supportedFeatures{specified: features}
}

func StrictWhitelistFeatures(features []string) FeatureSet {
	return &supportedFeatures{specified: features, strict: true}
}

type supportedFeatures struct {
	specified []string
	blacklist bool
	strict    bool
	features  map[string]*Feature
	cache     map[string]bool
}

func (self *supportedFeatures) Resolve(m *Module) error {
	self.cache = make(map[string]bool)
	self.features = make(map[string]*Feature)
	if self.blacklist {
		// copy all in
		for id, f := range m.Features() {
			self.features[id] = f
		}
		// remove blacklisted
		for _, id := range self.specified {
			if self.strict {
				if _, found := self.features[id]; !found {
					return c2.NewErr(id + " feature not found")
				}
			}
			delete(self.features, id)
		}
	} else {
		// copy in only items in whitelist
		for _, id := range self.specified {
			if f, found := m.Features()[id]; found {
				self.features[id] = f
			} else if self.strict {
				return c2.NewErr(id + " feature not found")
			}
		}
	}
	return nil
}

func (self *supportedFeatures) FeatureOn(f *IfFeature) (bool, error) {
	if on, found := self.cache[f.Expression()]; found {
		return on, nil
	}
	on, err := f.Evaluate(self.features)
	if err != nil {
		return false, err
	}
	self.cache[f.Expression()] = on
	return on, err
}

func checkFeature(m HasIfFeatures) (bool, error) {
	if len(m.IfFeatures()) > 0 {
		mod := Root(m)
		for _, iff := range m.IfFeatures() {
			if on, err := mod.featureSet.FeatureOn(iff); err != nil || !on {
				return false, err
			}
		}
	}
	return true, nil
}
