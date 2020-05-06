package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	carMaker := NewMaker()
	makeRes := carMaker.MakeCar()
	if makeRes != MAKE_CAR_OK {
		t.Fatalf("期望返回%s,实际返回 %s", MAKE_CAR_OK, makeRes)
	}
	t.Log("Make Car成功!!")
}
