package bridge

import (
	"fmt"
)

type IPhone interface {
	Start()
	Stop()
	Call()
}

type FoldedPhone struct {
	Br Brand
}

func (f *FoldedPhone) Start() {
	f.Br.Start()
	fmt.Println("折叠手机开机")
}

func (f *FoldedPhone) Call() {
	f.Br.Call()
	fmt.Println("折叠手机打电话")
}

func (f *FoldedPhone) Stop() {
	f.Br.Stop()
	fmt.Println("折叠手机关机")
}
