package val

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func Test_Conv(t *testing.T) {
	tests := []struct {
		F       Format
		In      interface{}
		Out     interface{}
		Invalid bool
	}{
		////////////
		{
			F:   FmtBool,
			In:  "true",
			Out: true,
		},
		{
			F:   FmtBool,
			In:  "1",
			Out: true,
		},
		{
			F:   FmtBool,
			In:  "0",
			Out: false,
		},
		{
			F:   FmtBool,
			In:  "false",
			Out: false,
		},
		{
			F:       FmtBool,
			In:      "bleep",
			Invalid: true,
		},
		////////////
		{
			F:   FmtBoolList,
			In:  "true",
			Out: []bool{true},
		},
		{
			F:   FmtBoolList,
			In:  []string{"1", "0"},
			Out: []bool{true, false},
		},
		{
			F:   FmtBoolList,
			In:  []interface{}{"1", false},
			Out: []bool{true, false},
		},
		{
			F:   FmtBoolList,
			In:  []interface{}{"1", false},
			Out: []bool{true, false},
		},
		{
			F:       FmtBoolList,
			In:      []interface{}{true, "bleep"},
			Invalid: true,
		},
		////////////
		{
			F:   FmtInt32,
			In:  0,
			Out: 0,
		},
		{
			F:   FmtInt32,
			In:  float64(99),
			Out: 99,
		},
		{
			F:   FmtInt32,
			In:  "99",
			Out: 99,
		},
		////////////
		{
			F:   FmtInt32List,
			In:  0,
			Out: []int{0},
		},
		{
			F:   FmtInt32List,
			In:  []float64{99, 98},
			Out: []int{99, 98},
		},
		{
			F:   FmtInt32List,
			In:  []string{"99", "98"},
			Out: []int{99, 98},
		},
		{
			F:   FmtInt32List,
			In:  []interface{}{"99", 98},
			Out: []int{99, 98},
		},
		////////////
		{
			F:   FmtDecimal64,
			In:  0,
			Out: float64(0),
		},
		{
			F:   FmtDecimal64,
			In:  float64(99),
			Out: float64(99),
		},
		{
			F:   FmtDecimal64,
			In:  "99",
			Out: float64(99),
		},
		////////////
		{
			F:   FmtDecimal64List,
			In:  0,
			Out: []float64{0},
		},
		{
			F:   FmtDecimal64List,
			In:  []float64{99, 98},
			Out: []float64{99, 98},
		},
		{
			F:   FmtDecimal64List,
			In:  []string{"99", "98"},
			Out: []float64{99, 98},
		},
		{
			F:   FmtDecimal64List,
			In:  []interface{}{"99", 98},
			Out: []float64{99, 98},
		},
	}
	for _, test := range tests {
		v, err := Conv(test.F, test.In)
		if test.Invalid {
			if err == nil {
				t.Errorf("test=%v. expected invalid, got %v", test, v)
			}
		} else if !test.Invalid && err != nil {
			t.Errorf("test=%v. err=%v", test, err)
		} else if v == nil {
			t.Errorf("not value returned for %v", test)
		} else {
			if !fc.AssertEqual(t, v.Value(), test.Out) {
				t.Log(test)
			}
		}
	}
}
