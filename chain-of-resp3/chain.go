package main

import "fmt"

type Request struct {
	Price int
	Name  string
}

type IHandler interface {
	Handle(req Request)
}

type Handler struct {
	Next IHandler
}

/**
* 系主任，处理 0- 5000 之内的订单
 */

type DepartmentHandler struct {
	Handler
}

func (dh *DepartmentHandler) Handle(req Request) {
	if req.Price >= 0 && req.Price <= 5000 {
		fmt.Printf("请求被 系主任 处理, 金额为 %v\n", req.Price)
	} else {
		if dh.Next != nil {
			dh.Next.Handle(req)
		}
	}
}

type CollegeHandler struct {
	Handler
}

func (ch *CollegeHandler) Handle(req Request) {
	if req.Price > 5000 && req.Price <= 10000 {
		fmt.Printf("请求被 院长 处理，金额为 %v\n", req.Price)
	} else {
		if ch.Next != nil {
			ch.Next.Handle(req)
		}
	}
}

type ViceSchoolHandler struct {
	Handler
}

func (vh *ViceSchoolHandler) Handle(req Request) {
	if req.Price > 10000 && req.Price <= 30000 {
		fmt.Printf("请求被 副校长 处理，金额为 %v\n", req.Price)
	} else {
		if vh.Next != nil {
			vh.Next.Handle(req)
		}
	}
}

type SchoolHandler struct {
	Handler
}

func (sh *SchoolHandler) Handle(req Request) {
	if req.Price > 30000 {
		fmt.Printf("请求被 校长处理，金额为 %v\n", req.Price)
	} else {
		if sh.Next != nil {
			sh.Next.Handle(req)
		}
	}
}
