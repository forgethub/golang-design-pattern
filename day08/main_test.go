package main

import "testing"

func TestSubject_Notify(t *testing.T) {
        sub := &Subject{observer: []IObserver{&Observer{}}}
        sub.Notify("hi")
}
