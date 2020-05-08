package main

import (
	"fmt"
)

type Phone struct {
	V5V Voltage5V
}

func (p *Phone) Charing() {
	output := p.V5V.Output5V()
	if output != 5 {
		fmt.Printf("充电失败，期待输出 5v，实际输出%v\n", output)
	}

	fmt.Printf("输出5v，充电成功!!\n")
}
