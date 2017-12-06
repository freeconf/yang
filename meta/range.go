package meta

import (
	"strconv"
	"strings"
)

type Range struct {
	Min    string
	Max    string
	notNil bool
	max    int
	min    int
}

func (r Range) Empty() bool {
	return !r.notNil
}

func NewRange(encoded string) (r Range, err error) {
	r.notNil = true
	// TODO: Support multiple ranges with '|'
	segments := strings.Split(string(encoded), "..")
	if len(segments) == 2 {
		r.Min = segments[0]
		if r.Min != "min" {
			if r.min, err = strconv.Atoi(r.Min); err != nil {
				return
			}
		}
		r.Max = segments[1]
	} else {
		r.Max = segments[0]
	}
	if r.Max != "max" {
		if r.max, err = strconv.Atoi(segments[0]); err != nil {
			return
		}
	}
	return
}

func (r Range) String() string {
	return r.Min + ".." + r.Max
}

// This is a start but I think the ideal solution collapses a list of
// ranges by looking at overlapping areas while validating each range
// is more restrictive as per RFC7950 Sec. 9.2.4 the "range" statement
//
// func MergeRanges(l []Range) []Range {
// 	return append(l, r)
// }
//
// func (a Range) Merge(b Range) Range {
// 	r := Range{
// 		notNil: true,
// 	}
// 	if a.Min == "" || a.min < b.min {
// 		r.min = b.min
// 		r.Min = b.Min
// 	} else {
// 		r.min = a.min
// 		r.Min = a.Min
// 	}
// 	if a.Max == "" || a.max > b.max {
// 		r.max = b.max
// 		r.Max = b.Max
// 	} else {
// 		r.max = a.max
// 		r.Max = a.Max
// 	}
// 	return r
// }
