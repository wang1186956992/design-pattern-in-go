package main

import (
	"fmt"
)

type IButton interface {
	Click()
	DblClick()
	MouseOn()
	MouseUp()
}

type ButtonAdapter struct{}

func (ba *ButtonAdapter) Click()    {}
func (ba *ButtonAdapter) DblClick() {}
func (ba *ButtonAdapter) MouseOn()  {}
func (ba *ButtonAdapter) MouseUp()  {}

type MyBtnAdapter struct {
	ButtonAdapter
}

func (mba *MyBtnAdapter) Click() {
	fmt.Println("MyBtnAdapter 点击了...")
}
