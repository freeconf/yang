package meta

// Things that need clean-up.  Often this used in driver contexts that release objects
// in other environments so GCs can cleanup
type Resource interface {
	Close() error
}

func CloseResource(o interface{}) error {
	if res, isRes := o.(Resource); isRes {
		return res.Close()
	}
	return nil
}
