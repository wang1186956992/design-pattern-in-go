package main

import (
	"fmt"
)

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	Name string
}

func (p *Pizza) Bake() {
	fmt.Printf("%s 开始烘焙\n", p.Name)
}

func (p *Pizza) Cut() {
	fmt.Printf("%s 开始切割\n", p.Name)
}

func (p *Pizza) Box() {
	fmt.Printf("%s 开始装箱\n", p.Name)
}

type BJCheesePizaa struct {
	Pizza
}

func (bp *BJCheesePizaa) Prepare() {
	fmt.Printf("%s 开始准备，需要北京的奶酪\n", bp.Name)
}

type BJPepperPizza struct {
	Pizza
}

func (bpp *BJPepperPizza) Prepare() {
	fmt.Printf("%s 开始准备，需要北京的胡椒\n", bpp.Name)
}

type LDCheesePizza struct {
	Pizza
}

func (lp *LDCheesePizza) Prepare() {
	fmt.Printf("%s 开始准备，需要伦敦的奶酪\n", lp.Name)
}

type LDPepperPizza struct {
	Pizza
}

func (lpp *LDPepperPizza) Prepare() {
	fmt.Printf("%s 开始准备，需要伦敦的胡椒\n", lpp.Name)
}
