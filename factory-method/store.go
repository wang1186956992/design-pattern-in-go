package main

import (
	"errors"
	"fmt"
)

type PizzaStore struct{}

func NewStore() *PizzaStore {
	return new(PizzaStore)
}

func (ps *PizzaStore) CreateOrders() (err error) {
	fmt.Println("请输入地址信息[bj,ld]:")
	loc := ""
	fmt.Scanln(&loc)
	var order IOrderPizza
	switch loc {
	case BJ:
		order = &BJOrderPizza{}
	case LD:
		order = &LDOrderPizza{}
	default:
		err = errors.New("不支持的地址类型!!")
		return
	}

	fmt.Println("请输入Pizza类型[cheese, pepper]:")
	pType := ""
	fmt.Scanln(&pType)
	pizza, err := order.CreatePizza(pType)
	if err != nil {
		return
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return
}
