package node

import (
	"strconv"
	"strings"

	"github.com/freeconf/gconf/c2"
)

type ListRange struct {
	Selector   PathMatcher
	CurrentRow int64
	StartRow   int64
	EndRow     int64
}

var listRangeUsage = "Range expression format {selector}!{startRow}-[{endRow}]"

func NewListRange(expression string) (lr *ListRange, err error) {
	lr = &ListRange{}
	bang := strings.IndexRune(expression, '!')
	if bang < 0 {
		return nil, c2.NewErrC(listRangeUsage, 400)
	}
	if lr.Selector, err = ParsePathExpression(expression[:bang]); err != nil {
		return nil, err
	}
	rowsExpression := expression[bang+1:]
	startEndStr := strings.Split(rowsExpression, "-")
	if lr.StartRow, err = strconv.ParseInt(startEndStr[0], 10, 64); err != nil {
		return nil, c2.NewErrC(listRangeUsage, 400)
	}
	if len(startEndStr) > 1 && len(startEndStr[1]) > 0 {
		if lr.EndRow, err = strconv.ParseInt(startEndStr[1], 10, 64); err != nil {
			return nil, c2.NewErrC(listRangeUsage, 400)
		}
	} else {
		lr.EndRow = -1
	}
	return
}

func (self *ListRange) CheckListPreConstraints(r *ListRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	if self.Selector.PathMatches(r.Base, r.Selection.Path) {
		if r.First {
			r.SetStartRow(self.StartRow)
			r.SetRow(self.StartRow)
		} else if r.Row64 >= self.EndRow && self.EndRow != -1 {
			return false, nil
		}
	}
	return true, nil
}
