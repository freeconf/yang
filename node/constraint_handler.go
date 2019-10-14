package node

type ConstraintHandler struct {
	NoFail      bool
	NewLocation string
	Violations  []error
}

func (self *ConstraintHandler) IncompleteResponse(location *Path) {
	//fc.Err.Println("Incomplete response served at " + location.String())
}

func (self *ConstraintHandler) LocatableNode(location *Path) {
	self.NewLocation = location.String()
}

func (self *ConstraintHandler) ConstraintViolation(violation error) error {
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


