package main

import "fmt"

type IHouseBuilder interface {
	BuildBasic(basic int)
	BuildWalls(walls int)
	Roofed(roof string)
	Build() *House
}

type House struct {
	Basic int
	Walls int
	Roof  string
}

func (h *House) String() string {
	hInfo := fmt.Sprintf("房子地基: %v 米, 房子墙 %v 厘米，房子吊顶: %v",
		h.Basic, h.Walls, h.Roof)
	return hInfo
}

type HouseBuilder struct {
	Hou *House
}

func NewHouseBuilder() *HouseBuilder {
	hb := &HouseBuilder{
		Hou: new(House),
	}
	return hb
}

func (hb *HouseBuilder) BuildBasic(basic int) {
	hb.Hou.Basic = basic
}

func (hb *HouseBuilder) BuildWalls(walls int) {
	hb.Hou.Walls = walls
}

func (hb *HouseBuilder) Roofed(roof string) {
	hb.Hou.Roof = roof
}

func (hb *HouseBuilder) Build() *House {
	return hb.Hou
}
