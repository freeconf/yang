package node

import (
	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/val"
)

// Field level filter that let's you see the differences in values from the default values.
//
// For more information, see:
//   https://tools.ietf.org/html/draft-ietf-netconf-restconf-16#section-4.8.9
//
type WithDefaults int

const (
	// Show all values, default and otherwise.  This is really the same as not specifying
	// any constraint
	WithDefaultsAll WithDefaults = iota

	// Hide values that match the default whether the user has explicitly set or not
	WithDefaultsTrim

	// See https://tools.ietf.org/html/rfc6243#section-3.3
	// This is really about returning the values the "happen" to match the default. Not all
	// implementations will be able to distinguish this.
	// Example:  if leaf x has a default of 5 and user never sets the value to 5, then using
	// this default will NOT return a value for x.  If however the user set the value to exactly
	// and explicitly to 5, then using this paramaeter WILL return a value for x as 5.
	//
	// While a very useful flag, this not supported yet only because there are no data node
	// implementations that support this yet. In addition, user interfaces would have to allow the user
	// to "clear" a value to assume the default v.s. setting a value to the default to pin the value to
	// that value in case the model default value changes.
	WithDefaultsExplicit

	// Not supported
	// See https://tools.ietf.org/html/rfc6243#section-3.4
	// This mangles the JSON output so instead of
	//   "x" : 10
	//  you have soemthing like this
	//   "x" : 10,
	//   "ietf-netconf-with-defaults" : {
	//       "x" : 5
	//   }
	//  although i'm not 100% positive, either way, this is not supported yet
	WithDefaultsAllTagged
)

func NewWithDefaultsConstraint(expression string) (WithDefaults, error) {
	switch expression {
	case "trim":
		return WithDefaultsTrim, nil
	case "explicit":
		return WithDefaultsExplicit, c2.NewErrC("explicity parameter not supported yet", 501)
	case "report-all":
		return WithDefaultsAll, nil
	case "report-all-tagged":
		return WithDefaultsAllTagged, c2.NewErrC("report-all-tagged parameter not supported yet", 501)
	}
	return WithDefaultsAll, c2.NewErrC("Invalid with-defaults constraint: "+expression, 400)
}

func (self WithDefaults) CheckFieldPostConstraints(r FieldRequest, hnd *ValueHandle) (bool, error) {
	if r.IsNavigation() || self == WithDefaultsAll {
		return true, nil
	}
	if !r.Meta.HasDefault() {
		return true, nil
	}

	// Only way to get here is if we're in WithDefaultsTrim so we want to return nil if value
	// matches the default
	def, err := NewValue(r.Meta.Type(), r.Meta.Default())
	if err != nil {
		return false, err
	}
	if val.Equal(def, hnd.Val) {
		hnd.Val = nil
		return true, nil
	}
	return true, nil
}
