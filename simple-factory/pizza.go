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

type CheesePizza struct {
	Pizza
}

func (cp *CheesePizza) Prepare() {
	fmt.Printf("%s 准备，需要奶酪\n", cp.Name)
}

type GreekPizza struct {
	Pizza
}

func (gp *GreekPizza) Prepare() {
	fmt.Printf("%s 准备，需要牛肉\n", gp.Name)
}
