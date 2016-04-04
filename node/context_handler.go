package node

type ContextHandler struct {
	NoFail      bool
	NewLocation string
	Violations  []error
	//OnIncomplete func(location *Path)
	//OnViolation func(violation error) error
	//OnNewLocation func(location *Path)
}

func (self *ContextHandler) IncompleteResponse(location *Path) {
	//c2.Err.Println("Incomplete response served at " + location.String())
}

func (self *ContextHandler) LocatableNode(location *Path) {
	self.NewLocation = location.String()
}

func (self *ContextHandler) ConstraintViolation(violation error) error {
	if !self.NoFail {
		return violation
	}
	if self.Violations == nil {
		self.Violations = []error{violation}
	} else {
		self.Violations = append(self.Violations, violation)
	}
	return nil
}


