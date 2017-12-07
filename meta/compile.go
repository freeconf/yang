package meta

func compile(m Meta, defs *defs) error {
	if defs != nil {
		if err := defs.compile(m); err != nil {
			return err
		}
	}

	if x, ok := m.(HasDataType); ok {
		if err := x.DataType().compile(m); err != nil {
			return err
		}
	}

	return nil
}
