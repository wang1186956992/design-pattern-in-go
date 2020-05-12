package main

import "fmt"

type IAdapter interface {
	Process(n int)
}

type Adapter struct {
	Ada IAdaptee
}

func (ad *Adapter) Process(n int) {
	ad.Ada.Foo()
	ad.Ada.Bar(n)
}

type IAdaptee interface {
	Foo()
	Bar(n int)
}

type Adaptee struct{}

func (tee *Adaptee) Foo() {
	fmt.Printf("被适配对象的Foo方法\n")
}

func (tee *Adaptee) Bar(n int) {
	fmt.Printf("被适配对象的Bar方法, n的值 %v\n", n)
}
