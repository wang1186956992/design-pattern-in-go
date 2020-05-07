package decorator

import (
	"fmt"
)

type Decorator struct {
	Coffe ICoffe
	Desc  string
	Cos   float64
}

func (d *Decorator) Description() string {
	return fmt.Sprintf("%v %v %v\n", d.Desc, d.Cos, d.Coffe.Description())
}

func (d *Decorator) Cost() float64 {
	return d.Cos + d.Coffe.Cost()
}

type Milk struct {
	Decorator
}

func NewMilk(coffe ICoffe) *Milk {
	return &Milk{
		Decorator{
			Coffe: coffe,
			Desc:  "牛奶",
			Cos:   2.0,
		},
	}
}

type Soy struct {
	Decorator
}

func NewSoy(coffe ICoffe) *Soy {
	return &Soy{
		Decorator{
			Coffe: coffe,
			Desc:  "豆浆",
			Cos:   1.5,
		},
	}
}

type Chcolate struct {
	Decorator
}

func NewChcolate(coffe ICoffe) *Chcolate {
	return &Chcolate{
		Decorator{
			Coffe: coffe,
			Desc:  "巧克力",
			Cos:   3.0,
		},
	}
}
