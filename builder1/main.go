package main

import "fmt"

func main() {
	var hb IHouseBuilder = NewHouseBuilder()
	hb.BuildBasic(5)
	hb.BuildWalls(10)
	hb.Roofed("石头顶")
	hou := hb.Build()
	fmt.Println(hou)

	var hbCropto IHouseBuilder = NewHouseBuilder()
	hbCropto.BuildBasic(10)
	hbCropto.BuildWalls(20)
	hbCropto.Roofed("水晶顶")
	houCropto := hbCropto.Build()
	fmt.Println(houCropto)
}
