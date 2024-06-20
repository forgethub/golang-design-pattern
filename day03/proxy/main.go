package main

import "fmt"

type Isubject interface {
    Request()
}

type RealSubject struct{}

func (r *RealSubject) Request() {
    fmt.Println("真实实现")
}

type Proxy struct {
    realSubject Isubject
}

func (p *Proxy) Request() {
    if p.realSubject == nil {
        p.realSubject = &RealSubject{}
    }
    fmt.Println("代理逻辑")
    p.realSubject.Request()
}

func NewProxy(i Isubject) *Proxy {
    return &Proxy{realSubject: i}
}

func main() {
    p := NewProxy(&RealSubject{})
    p.Request()
}
