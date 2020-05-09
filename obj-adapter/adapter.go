package main

import (
	"fmt"
)

type Votagle220V struct{}

func (v2 *Votagle220V) Output220V() int {
	fmt.Println("源电压，输出220v电压")
	return 220
}

type IVotagle5V interface {
	Output5V() int
}

type VoltageAdapter struct {
	V220V *Votagle220V
}

func (va *VoltageAdapter) Output5V() int {
	dst := 0
	if va.V220V != nil {
		fmt.Println("将220V电压转换成5V电压")
		dst = va.V220V.Output220V() / 44
	}
	return dst
}

type Phone struct {
	V5V IVotagle5V
}

func (p *Phone) Charge() {
	dst := p.V5V.Output5V()
	fmt.Printf("充电中，电压为%vV\n", dst)
}
