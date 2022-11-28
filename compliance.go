package yang

// Compliance is the global variable that sets the default behavior if the
// FreeCONF YANG library.
//
// By default this is for strict IETF compliance!
//
// This sets just the default behavior of data structures, each individual
// instance should allow for controlling the compliance of that instance should
// you need to have instances in different modes at the same time.
//
// If you wish change default compliance, be sure to do it at the beginning of your
// to application before any objects are constructed.
var Compliance = ComplianceOptions{}

// Simplified are the settings pre 2023 before true IETF compliance was
// attempted. To use this:
//
//  restconf.Compliance = restconf.Simplified
//
// or you can just set individual settings on restconf.Compliance global variable
var Simplified = ComplianceOptions{

	QualifyNamespaceDisabled: true,
}

// ComplianceOptions hold all the compliance settings. If you need to set this to true
// then you will not be compatiable with any NETCONF RFCs
type ComplianceOptions struct {

	// QualifyNamespaceDisabled when true then all JSON object keys will not
	// include YANG module according to RFC7952.
	QualifyNamespaceDisabled bool
}
