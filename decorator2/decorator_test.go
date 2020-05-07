package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	var coffe ICoffe = &ShortBlack{
		Desc: "黑咖啡",
		Cos:  20.0,
	}
	t.Logf("%v %v", coffe.Description(), coffe.Cost())
	coffe = NewMilk(coffe)
	t.Logf("%v %v", coffe.Description(), coffe.Cost())
	coffe = NewChcolate(coffe)
	t.Logf("%v %v", coffe.Description(), coffe.Cost())
	coffe = NewChcolate(coffe)
	t.Logf("%v %v", coffe.Description(), coffe.Cost())

}
