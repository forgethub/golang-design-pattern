package main

import "fmt"

type Shape interface {
    draw()
}

type Rectangle struct{}

func (r *Rectangle) draw() {
    fmt.Println("Drawing a rectangle")
}

type Circle struct{}

func (c *Circle) draw() {
    fmt.Println("Drawing a circle")
}

type Diamond struct{}

func (d *Diamond) draw() {
    fmt.Println("Drawing a diamond")
}

type ShapeFactory interface {
    createShape() Shape
}

type ConcreteShapeFactory struct{}

func (f *ConcreteShapeFactory) createShape(shapeType string) Shape {
    switch shapeType {
    case "rectangle":
        return &Rectangle{}
    case "circle":
        return &Circle{}
    case "diamond":
        return &Diamond{}
    default:
        fmt.Println("Shape type not supported")
        return nil
    }
}

func main() {
    factory := &ConcreteShapeFactory{}
    shape1 := factory.createShape("rectangle")
    shape1.draw()
}
