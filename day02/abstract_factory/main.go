package main

import "fmt"

// Shape 定义了形状接口
type Shape interface {
    Draw()
}

// Rectangle 是长方形的具体实现
type Rectangle struct{}

func (r *Rectangle) Draw() {
    fmt.Println("Drawing a rectangle")
}

// Circle 是圆形的具体实现
type Circle struct{}

func (c *Circle) Draw() {
    fmt.Println("Drawing a circle")
}

// Diamond 是菱形的具体实现
type Diamond struct{}

func (d *Diamond) Draw() {
    fmt.Println("Drawing a diamond")
}

// ShapeFactory 定义了形状工厂接口
type ShapeFactory interface {
    CreateRectangle() Shape
    CreateCircle() Shape
    CreateDiamond() Shape
}

// ConcreteShapeFactory 是具体的形状工厂
type ConcreteShapeFactory struct{}

func (f *ConcreteShapeFactory) CreateRectangle() Shape {
    return &Rectangle{}
}

func (f *ConcreteShapeFactory) CreateCircle() Shape {
    return &Circle{}
}

func (f *ConcreteShapeFactory) CreateDiamond() Shape {
    return &Diamond{}
}

func main() {
    // 创建形状工厂实例
    factory := &ConcreteShapeFactory{}

    // 使用工厂来创建形状
    rectangle := factory.CreateRectangle()
    circle := factory.CreateCircle()
    diamond := factory.CreateDiamond()

    // 使用形状
    rectangle.Draw()  // 输出: Drawing a rectangle
    circle.Draw()    // 输出: Drawing a circle
    diamond.Draw()   // 输出: Drawing a diamond
}
