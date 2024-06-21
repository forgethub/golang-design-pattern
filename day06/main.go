package main

import (
    "fmt"
)

type Component interface {
    Weight() int
}

type Leaf struct{
    weight int
}

func (l *Leaf) Weight() int {
    return l.weight 
}

type Composite struct {
    children []Component
}

func (c *Composite) Weight() int {
    n := 0
	for _, children := range c.children {
		n += children.Weight()
	}
	return n
}

func (c *Composite) Add(i Component) {
    c.children = append(c.children, i)
}


func main() {
    composite := Composite{children: []Component{&Leaf{weight: 5}, &Leaf{weight: 7}}}
    fmt.Println(composite.Weight())
}

