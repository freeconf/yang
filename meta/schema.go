package meta

type schemaError struct {
	s string
}

func (err *schemaError) Error() string {
	return err.s
}
