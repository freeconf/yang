package meta

func compile(m Meta, defs *defs) error {
	// order is somewhat important here:
	// 1. want typedefs resolved so datadefs can find them.
	// 2. want augments to change items after datadefs are compiled
	if x, ok := m.(HasTypedefs); ok {
		for _, y := range x.Typedefs() {
			if err := y.compile(); err != nil {
				return err
			}
		}
	}

	if x, ok := m.(HasGroupings); ok {
		for _, y := range x.Groupings() {
			if err := y.compile(); err != nil {
				return err
			}
		}
	}

	if x, ok := m.(HasConditions); ok {
		for _, y := range x.Conditions() {
			if err := y.compile(); err != nil {
				return err
			}
		}
	}

	if defs != nil {
		if err := defs.compile(m); err != nil {
			return err
		}
	}

	if x, ok := m.(HasDataType); ok {
		if err := x.DataType().compile(); err != nil {
			return err
		}
	}

	if x, ok := m.(HasAugments); ok {
		for _, y := range x.Augments() {
			if err := y.compile(); err != nil {
				return err
			}
		}
	}

	return nil
}
