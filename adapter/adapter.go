package main

import (
	"fmt"
)

type Voltage220V struct{}

func (v2 *Voltage220V) Output220V() int {
	fmt.Println("220v电源，输出220v电压...")
	return 220
}

type Voltage5V interface {
	Output5V() int
}

type VoltageAdapter struct {
	Voltage220V
}

func (va *VoltageAdapter) Output5V() int {
	srcV := va.Output220V()
	dstV := srcV / 44
	return dstV
}
