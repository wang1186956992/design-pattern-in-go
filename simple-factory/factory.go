package main

import (
	"errors"
)

const (
	CHEESE = "cheese"
	GREEK  = "greek"
)

var (
	PizzaFac *PizzaFactory = &PizzaFactory{}
)

type PizzaFactory struct{}

func (f *PizzaFactory) GetPizza(Name string) (pizza IPizza, err error) {

	switch Name {
	case CHEESE:
		pizza = &CheesePizza{
			Pizza: Pizza{
				Name: "奶酪披萨",
			},
		}
	case GREEK:
		pizza = &GreekPizza{
			Pizza: Pizza{
				Name: "希腊披萨",
			},
		}
	default:
		err = errors.New("不支持的Pizza类型!!")
		return
	}

	return

}
