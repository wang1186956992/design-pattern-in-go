package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	var phone IPhone = &FoldedPhone{
		Br: &Oppo{},
	}

	phone.Start()
	phone.Call()
	phone.Stop()

	fmt.Println("======================")

	var phone2 IPhone = &FoldedPhone{
		Br: &HuaWei{},
	}
	phone2.Start()
	phone2.Call()
	phone2.Stop()
}
