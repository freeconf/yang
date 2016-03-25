package meta

type Details struct {
	// Tri-state boolean - true, false, and undeclared
	ConfigPtr *bool
	MandatoryPtr *bool
}

func (d *Details) Config(p Path) bool {
	if d.ConfigPtr != nil {
		return *d.ConfigPtr
	}

	// if details are on leaf, then p is parent container, otherwise
	// p is what we're supposed to check config for
	if p != nil {
		if hasDetails, ok := p.Meta().(HasDetails); ok {
			if parentDetails := hasDetails.Details(); parentDetails == d {
				return hasDetails.Details().Config(p.MetaParent())
			}
			return hasDetails.Details().Config(p)
		}
	}
	return true
}

func (d *Details) SetConfig(config bool) {
	d.ConfigPtr = &config
}

func (d *Details) Mandatory() bool {
	if d.MandatoryPtr != nil {
		return *d.MandatoryPtr
	}
	return false
}

func (d *Details) SetMandatory(mandatory bool) {
	d.MandatoryPtr = &mandatory
}

