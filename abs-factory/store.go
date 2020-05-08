package main

import (
	"errors"
	"fmt"
)

type PizzaStore struct{}

func NewPizzaStore() *PizzaStore {
	return &PizzaStore{}
}

func (ps *PizzaStore) bookAnOrder() (err error) {
	loc := ""
	fmt.Println("请输入地址:")
	fmt.Scanln(&loc)
	var pf IFactory
	switch loc {
	case BJ:
		pf = new(BJFactory)
	case LD:
		pf = new(LDFactory)
	default:
		err = errors.New("不支持的地点类型!!")
		return
	}

	fmt.Println("请输入pizza类型:")
	pType := ""
	fmt.Scanln(&pType)
	pizza, err := pf.CreatePizza(pType)

	if err != nil {
		return
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return
}

func (ps *PizzaStore) BookOrders() {
	for {
		err := ps.bookAnOrder()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
