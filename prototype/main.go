package main

import (
	"fmt"
)

func main() {
	p := &Proto{
		Name: "张三",
		Age:  30,
	}

	if p1, err := p.Clone(); err == nil {
		fmt.Printf("p1的值%v\n", *(p1.(*Proto)))
		fmt.Printf("p1的地址:%p\n", p1)
	}

	if p2, err := p.Clone(); err == nil {
		fmt.Printf("p2的值%v\n", *(p2.(*Proto)))
		fmt.Printf("p2的地址:%p\n", p2)
	}

}
