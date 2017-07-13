package gen

 

/*
  Car - Vehicle of sorts
*/
type Car struct {
	Tire []*Tire
	Miles int64
	LastRotation int64
	Running bool
	Speed int
	Engine *Engine
}

 

/*
  Tire - rubber circular part that makes contact with road
*/
type Tire struct {
	Pos int
	Size string
	Worn bool
	Wear float64
	Flat bool
}

 

/*
  Engine - 
*/
type Engine struct {
	Specs *Specs
}

 

/*
  Specs - 
*/
type Specs struct {
	Horsepower int
}

