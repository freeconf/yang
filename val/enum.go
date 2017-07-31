package val

type Enumerable interface {
	EnumById(label string) (Enum, bool)
	EnumLabel(id int) (Enum, bool)
	EnumLen() int
	Enums() []Enum
}
