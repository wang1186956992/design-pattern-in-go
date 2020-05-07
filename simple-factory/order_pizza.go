package main

import (
	"fmt"
)

type OrderPizza struct {
}

func NewOrderPizza() *OrderPizza {
	op := new(OrderPizza)
	return op
}

func (op *OrderPizza) bookAnOrder() (err error) {
	pizzaType := ""
	fmt.Println("请输入Pizza类型:")
	fmt.Scanln(&pizzaType)
	pizza, err := PizzaFac.GetPizza(pizzaType)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return

}

func (op *OrderPizza) BookOrders() {
	for {
		err := op.bookAnOrder()
		if err != nil {
			break
		}
	}
}
