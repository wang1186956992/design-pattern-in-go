package main

import (
	"errors"
)

const (
	BJ     = "bj"
	LD     = "ld"
	CHEESE = "cheese"
	PEPPER = "pepper"
)

type IOrderPizza interface {
	CreatePizza(typeName string) (pizza IPizza, err error)
}

type BJOrderPizza struct{}

type LDOrderPizza struct{}

func (bo *BJOrderPizza) CreatePizza(typeName string) (pizza IPizza, err error) {
	switch typeName {
	case CHEESE:
		pizza = &BJCheesePizaa{
			Pizza: Pizza{
				Name: "北京奶酪披萨",
			},
		}
	case PEPPER:
		pizza = &BJPepperPizza{
			Pizza: Pizza{
				Name: "北京胡椒披萨",
			},
		}
	default:
		err = errors.New("不支持的北京披萨类型!!")
	}

	return
}

func (lo *LDOrderPizza) CreatePizza(typeName string) (pizza IPizza, err error) {
	switch typeName {
	case CHEESE:
		pizza = &LDCheesePizza{
			Pizza: Pizza{
				Name: "伦敦奶酪披萨",
			},
		}
	case PEPPER:
		pizza = &LDPepperPizza{
			Pizza: Pizza{
				Name: "伦敦胡椒披萨",
			},
		}
	default:
		err = errors.New("不支持的伦敦披萨类型!!")
	}
	return
}
