package board

// checkType takes a variable number of typeCheckables and checks them, returning the
// first error it discovers.
func checkType(types ...typeCheckable) error {
	for _, t := range types {
		if err := t.typeError(); err != nil {
			return err
		}
	}

	return nil
}
