package main

import (
    "fmt"
)

// Request 定义请求的结构
type Request struct {
    Message string
}

// Handler 定义处理请求的接口
type Handler interface {
    HandleRequest(request *Request)
    SetNext(handler Handler)
}

// ConcreteHandler 实现具体的处理器
type ConcreteHandler struct {
    nextHandler Handler
    name string
}


func (c *ConcreteHandler) HandleRequest(request *Request) {
    if c.nextHandler != nil {
        fmt.Printf("%s pass to the next handler\n", c.name)
        c.nextHandler.HandleRequest(request)
    }
    fmt.Printf("call to the %s handler\n", c.name)
}

func (c *ConcreteHandler) SetNext(handler Handler) {
    c.nextHandler = handler
}

func main() {
    // 创建处理器链
    handlerA := &ConcreteHandler{name: "Handler A"}
    handlerB := &ConcreteHandler{name: "Handler B"}
    handlerC := &ConcreteHandler{name: "Handler C"}

    handlerA.SetNext(handlerB)
    handlerB.SetNext(handlerC)

    // 创建请求并传递
    request := &Request{Message: "Request 1"}
    handlerA.HandleRequest(request)
}
