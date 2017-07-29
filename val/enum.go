package val

type Enumerable interface {
	EnumById(label string) (EnumVal, bool)
	EnumLabel(id int) (EnumVal, bool)
	EnumLen() int
	Enums() []EnumVal
}
