package meta

type FeatureSet interface {
	FeatureOn(*IfFeature) (bool, error)
}

type SupportedFeatures struct {
	features map[string]*Feature
	cache    map[string]bool
}

func Whitelist(m *Module, features []string) *SupportedFeatures {
	enabled := make(map[string]*Feature)
	for _, id := range features {
		if f, found := m.Features()[id]; found {
			enabled[id] = f
		}
	}
	return NewSupportedFeatures(enabled)
}

func Backlist(m *Module, features []string) *SupportedFeatures {
	enabled := make(map[string]*Feature)
	for id, f := range m.Features() {
		enabled[id] = f
	}
	for _, j := range features {
		delete(enabled, j)
	}
	return NewSupportedFeatures(enabled)
}

func NewSupportedFeatures(features map[string]*Feature) *SupportedFeatures {
	return &SupportedFeatures{
		features: features,
		cache:    make(map[string]bool),
	}
}

func (self *SupportedFeatures) FeatureOn(f *IfFeature) (bool, error) {
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
