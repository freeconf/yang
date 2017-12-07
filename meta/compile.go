package meta

func compile(m Meta, defs *defs) error {
	if defs != nil {
		if err := defs.compile(m); err != nil {
			return err
		}
	}

	if x, ok := m.(HasType); ok {
		if err := x.Type().compile(m); err != nil {
			return err
		}
	}

	return nil
}
