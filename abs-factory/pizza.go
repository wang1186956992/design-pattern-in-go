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

type BJCheesePizza struct {
	Pizza
}

func (bcp *BJCheesePizza) Prepare() {
	fmt.Printf("%s 开始准备，需要北京奶酪\n", bcp.Name)
}

type BJPepperPizza struct {
	Pizza
}

func (bpp *BJPepperPizza) Prepare() {
	fmt.Printf("%s 开始准备，需要北京胡椒\n", bpp.Name)
}

type LDCheesePizza struct {
	Pizza
}

func (lcp *LDCheesePizza) Prepare() {
	fmt.Printf("%s 开始准备，需要伦敦奶酪\n", lcp.Name)
}

type LDPepperPizza struct {
	Pizza
}

func (lpp *LDPepperPizza) Prepare() {
	fmt.Printf("%s 开始准备，需要伦敦胡椒\n", lpp.Name)
}
