package main

import "fmt"

// 主题接口
type Subject interface {
    Request()
}

// 真实主题
type RealSubject struct{}

func (r *RealSubject) Request() {
    fmt.Println("RealSubject: Handling request.")
}

// 代理主题
type Proxy struct {
    realSubject *RealSubject
}

func NewProxy(realSubject *RealSubject) *Proxy {
    return &Proxy{realSubject: realSubject}
}

func (p *Proxy) Request() {
    if p.realSubject == nil {
        p.realSubject = &RealSubject{}
    }
    // 代理可以在这里添加额外的逻辑
    fmt.Println("Proxy: Logging request.")
    p.realSubject.Request()
}

func main() {
    realSubject := &RealSubject{}
    proxy := NewProxy(realSubject)

    fmt.Println("RealSubject直接调用:")
    realSubject.Request()

    fmt.Println("\nProxy间接调用:")
    proxy.Request()
}
