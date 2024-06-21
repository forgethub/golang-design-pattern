package main

import "fmt"

// SubSystemOne 子系统一
type SubSystemOne struct{}

func (s *SubSystemOne) OperationOne() {
	fmt.Println("Operation One of SubSystemOne")
}

// SubSystemTwo 子系统二
type SubSystemTwo struct{}

func (s *SubSystemTwo) OperationTwo() {
	fmt.Println("Operation Two of SubSystemTwo")
}

// Facade 门面接口
type Facade struct {
	subSystemOne *SubSystemOne
	subSystemTwo *SubSystemTwo
}

// NewFacade 创建门面实例
func NewFacade() *Facade {
	return &Facade{
		subSystemOne: &SubSystemOne{},
		subSystemTwo: &SubSystemTwo{},
	}
}

// Operation 通过门面进行操作
func (f *Facade) Operation() {
	fmt.Println("Facade Operation")
	f.subSystemOne.OperationOne()
	f.subSystemTwo.OperationTwo()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
