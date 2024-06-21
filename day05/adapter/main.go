package main

type Target interface {
    Request() string
}

type Adaptee struct {}

func (a *Adaptee) SpecificRequest() string {
    return "SpecificRequest"
}

type Adapter struct {
    *Adaptee
}

func (a *Adapter) Request() string {
    return a.Adaptee.SpecificRequest() + "强行适配"
}
