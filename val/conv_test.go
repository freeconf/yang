package val

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func Test_Conv(t *testing.T) {
	type CustomUint32 uint32
	type CustomInt32 int32
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
		{
			F:   FmtInt32,
			In:  CustomInt32(10),
			Out: 10,
		},
		////////////
		{
			F:   FmtUInt16,
			In:  float64(6000),
			Out: uint16(6000),
		},
		{
			F:       FmtUInt16,
			In:      0x0001ffff,
			Out:     uint16(65535),
			Invalid: true,
		},
		////////////
		{
			F:   FmtInt32,
			In:  int32(1000),
			Out: int(1000),
		},
		////////////
		{
			F:   FmtUInt32,
			In:  0xffffffff,
			Out: uint(0xffffffff),
		},
		{
			F:   FmtUInt32,
			In:  CustomUint32(10),
			Out: uint(10),
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
