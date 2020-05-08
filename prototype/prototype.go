package main

import (
	"encoding/json"
	"fmt"
)

type IProto interface {
	Clone() (ret interface{}, err error)
}

type Proto struct {
	Name string
	Age  int32
}

func (p *Proto) Clone() (ret interface{}, err error) {

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Json序列化失败, %s\n", err.Error())
		return
	}
	ret = &Proto{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		fmt.Printf("Json反序列化失败, %s\n", err.Error())
		return
	}
	return
}
