package c2

func AppendUrlSegment(a string, b string) string {
	if a == "" || b == "" {
		return a + b
	}
	slashA := a[len(a)-1] == '/'
	slashB := b[0] == '/'
	if slashA != slashB {
		return a + b
	}
	if slashA && slashB {
		return a + b[1:]
	}
	return a + "/" + b
}
