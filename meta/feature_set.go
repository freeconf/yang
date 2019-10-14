package meta

type FeatureSet interface {
	Initialize(m *Module) error
	Resolve(*IfFeature) (bool, error)
}

func AllFeaturesOn() FeatureSet {
	return FeaturesOff([]string{})
}

// All other features will be off. Effectively a blacklist.
func FeaturesOn(features []string) FeatureSet {
	return &supportedFeatures{specified: features}
}

// All other features will be on. Effectively a whitelist.
func FeaturesOff(features []string) FeatureSet {
	return &supportedFeatures{specified: features, otherwiseOn: true}
}

type supportedFeatures struct {
	specified   []string
	otherwiseOn bool
	enabled     map[string]*Feature
	cache       map[string]bool
}

func (self *supportedFeatures) Initialize(m *Module) error {
	self.cache = make(map[string]bool)
	enabled := make(map[string]*Feature)
	if self.otherwiseOn {
		// copy all in
		for id, f := range m.Features() {
			enabled[id] = f
		}
		// remove blacklisted
		for _, id := range self.specified {
			delete(enabled, id)
		}
	} else {
		// copy in only items in whitelist
		for _, id := range self.specified {
			if f, found := m.Features()[id]; found {
				enabled[id] = f
			}
		}
	}
	if self.enabled == nil {
		self.enabled = enabled
	} else {
		for id, f := range enabled {
			self.enabled[id] = f
		}
	}
	return nil
}

func (self *supportedFeatures) Resolve(f *IfFeature) (bool, error) {
	if on, found := self.cache[f.Expression()]; found {
		return on, nil
	}
	on, err := f.Evaluate(self.enabled)
	if err != nil {
		return false, err
	}
	self.cache[f.Expression()] = on
	return on, err
}

func checkFeature(m HasIfFeatures) (bool, error) {
	if len(m.IfFeatures()) > 0 {
		mod := RootModule(m)
		for _, iff := range m.IfFeatures() {
			if on, err := mod.featureSet.Resolve(iff); err != nil || !on {
				return false, err
			}
		}
	}
	return true, nil
}
