package main

import (
    "fmt"
)

type IObserver interface {
    Update(msg string)
}

type ISubject interface {
    Notify(msg string)
    Remove(observer IObserver)
    Add(observer IObserver)
}

type Subject struct {
    observer []IObserver
}

func (s *Subject) Notify(msg string) {
    for _, o := range s.observer {
        o.Update(msg)
    }
}

func (s *Subject) Remove(observer IObserver) {
    for i, o := range s.observer {
        if o == observer {
            s.observer = append(s.observer[:i], s.observer[i+1:]...)
        }
    }

}

func (s *Subject) Add(observer IObserver) {
    s.observer = append(s.observer, observer)
}

type Observer struct {}

func (o *Observer) Update(msg string) {
    fmt.Printf("收到通知 %s", msg)
}

