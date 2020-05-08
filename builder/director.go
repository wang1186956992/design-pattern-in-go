package main

const (
	COMMON = "common"
	HIGH   = "high"
)

type Director struct {
	Bu IBuilder
}

func NewDirector(Bu IBuilder) *Director {
	return &Director{
		Bu: Bu,
	}
}
func (d *Director) BuildHouse() *House {
	d.Bu.BuildBasic()
	d.Bu.BuildWalls()
	d.Bu.Roofed()
	return d.Bu.Build()
}
