package meta

// Details manages the specifics of things like leaf that can be marked as config or not
// or mandatory or not.  Essentialy meta about the meta.
// You'll find more information here:
//   https://tools.ietf.org/html/rfc6020#section-7.6
type Details struct {
	// Tri-state boolean - true, false, and undeclared
	ConfigPtr    *bool
	MandatoryPtr *bool
}

// HasDetails for meta that has 'config' or 'mandatory' contraints
type HasDetails interface {
	Details() *Details
}

// Config return true if this item is configurable, otherwise it's operational.
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

// Mandatory return true if this item is mandatory in datastores (not in edits).
func (d *Details) Mandatory() bool {
	if d.MandatoryPtr != nil {
		return *d.MandatoryPtr
	}
	return false
}

func (d *Details) SetMandatory(mandatory bool) {
	d.MandatoryPtr = &mandatory
}
