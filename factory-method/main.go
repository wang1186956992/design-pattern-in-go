package main

import (
	"fmt"
)

func main() {

	for {
		err := NewStore().CreateOrders()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
