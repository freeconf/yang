package meta

// Iterator iterates over meta. Use meta.Children for most common way to
// iterate.
type Iterator interface {
	Next() Definition
	HasNext() bool
}

// Children of a meta list returning only containers, lists and leafs
func Children(m HasDataDefs) Iterator {
	return Iterate(m.DataDefs())
}

func Iterate(dataDefs []Definition) Iterator {
	return &iterator{dataDefs: dataDefs}
}

type iterator struct {
	position int
	dataDefs []Definition
}

func (i *iterator) Next() Definition {
	i.position++
	return i.dataDefs[i.position-1]
}

func (i *iterator) HasNext() bool {
	return i.position < len(i.dataDefs)
}
