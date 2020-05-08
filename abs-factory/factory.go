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

type IFactory interface {
	CreatePizza(pizzaType string) (pizza IPizza, err error)
}

type BJFactory struct{}

type LDFactory struct{}

func (bf *BJFactory) CreatePizza(pizzaType string) (pizza IPizza, err error) {
	switch pizzaType {
	case CHEESE:
		pizza = &BJCheesePizza{
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

func (lf *LDFactory) CreatePizza(pizzaType string) (pizza IPizza, err error) {
	switch pizzaType {
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
