package main

import "fmt"

// 实现部分接口定义
type Printer interface {
    PrintFile()
}

type Epson struct {}

func (e *Epson) PrintFile() {
    fmt.Println("Printing by a EPSON Printer")
}

type HP struct {}

func (e *HP) PrintFile() {
    fmt.Println("Printing by a HP Printer")
}

// 抽象部分接口定义

type Computer interface {
    Print()
}

type Mac struct {
    printer Printer 
}

type Linux struct {
    printer Printer 
}

func (m *Mac) Print() {
    fmt.Println("mac 电脑")
    m.printer.PrintFile()
}

func (m *Linux) Print() {
    fmt.Println("Linux 电脑")
    m.printer.PrintFile()
}
