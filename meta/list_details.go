package meta

import "math"

// ListDetails about 'list' or 'leaf-list' as well as 'refine' or 'augment' of said elements
type ListDetails struct {
	minElementsPtr *int
	maxElementsPtr *int
	unboundedPtr   *bool
}

type HasListDetails interface {
	ListDetails() *ListDetails
}

func (y *ListDetails) SetMinElements(n int) {
	y.minElementsPtr = &n
}

func (y *ListDetails) SetMaxElements(n int) {
	y.maxElementsPtr = &n
}

func (y *ListDetails) SetUnbounded(unbounded bool) {
	y.unboundedPtr = &unbounded
}

func (y *ListDetails) HasMaxElements() bool {
	return y.maxElementsPtr != nil
}

func (y *ListDetails) HasMinElements() bool {
	return y.minElementsPtr != nil
}

func (y *ListDetails) MinElements() int {
	if y.minElementsPtr != nil {
		return *y.minElementsPtr
	}
	return 0
}

func (y *ListDetails) Unbounded() bool {
	if y.unboundedPtr != nil {
		return *y.unboundedPtr
	}
	return y.maxElementsPtr == nil
}

func (y *ListDetails) ExplicitlyUnbounded() bool {
	return y.unboundedPtr != nil
}

func (y *ListDetails) MaxElements() int {
	if y.maxElementsPtr != nil {
		return *y.maxElementsPtr
	}
	return math.MaxInt32
}
