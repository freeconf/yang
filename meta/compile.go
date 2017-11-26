package meta

func compile(m Meta, defs *defs) error {
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
