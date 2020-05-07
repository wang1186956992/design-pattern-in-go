package bridge

import (
	"fmt"
)

type Brand interface {
	Start()
	Call()
	Stop()
}

type HuaWei struct{}

func (h *HuaWei) Start() {
	fmt.Println("华为手机开机")
}

func (h *HuaWei) Call() {
	fmt.Println("华为手机打电话")
}

func (h *HuaWei) Stop() {
	fmt.Println("华为手机关机")
}

type Oppo struct{}

func (o *Oppo) Start() {
	fmt.Println("Oppo手机开机")
}

func (o *Oppo) Call() {
	fmt.Println("Oppo手机打电话")
}

func (o *Oppo) Stop() {
	fmt.Println("Oppo手机关机")
}
