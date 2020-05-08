package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("请输入想要建造的房子类型:")
		var buildType = ""
		fmt.Scanln(&buildType)
		var builder IBuilder
		var house *House
		switch buildType {
		case COMMON:
			builder = &CommonBuilder{
				Hou: &House{},
			}
			house = NewDirector(builder).BuildHouse()
		case HIGH:
			builder = &HighBuilder{
				Hou: &House{},
			}
			house = NewDirector(builder).BuildHouse()
		default:
			fmt.Println("不支持的房子类型!!")
			return
		}
		fmt.Println(*house)
	}
}
