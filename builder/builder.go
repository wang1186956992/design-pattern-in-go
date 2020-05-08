package main

type House struct {
	Basic string
	Walls string
	Roof  string
}

type IBuilder interface {
	BuildBasic()
	BuildWalls()
	Roofed()
	Build() *House
}

type CommonBuilder struct {
	Hou *House
}

func (cb *CommonBuilder) BuildBasic() {
	cb.Hou.Basic = "普通房子地基 5 米"
}

func (cb *CommonBuilder) BuildWalls() {
	cb.Hou.Walls = "普通房子砌墙 10 cm"
}

func (cb *CommonBuilder) Roofed() {
	cb.Hou.Roof = "普通房子屋顶"
}

func (cb *CommonBuilder) Build() *House {
	return cb.Hou
}

type HighBuilder struct {
	Hou *House
}

func (hb *HighBuilder) BuildBasic() {
	hb.Hou.Basic = "高楼地基 100 米"
}

func (hb *HighBuilder) BuildWalls() {
	hb.Hou.Walls = "高楼砌墙 20 cm"
}

func (hb *HighBuilder) Roofed() {
	hb.Hou.Roof = "高楼透明屋顶"
}

func (hb *HighBuilder) Build() *House {
	return hb.Hou
}
