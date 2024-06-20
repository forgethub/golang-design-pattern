package main

import (
    "fmt"
)

type IWomen interface {
    MakeEyeswithMan() string
}

type JinLian struct {}

func (j *JinLian) MakeEyeswithMan() string {
    return "金莲对你笑"
}

type JinLianAge struct {
    jinLian IWomen
    age int
}

func NewJinLianAge(w IWomen, age int) *JinLianAge {
    return &JinLianAge{jinLian: w, age: age}

}

func (j *JinLianAge) MakeEyeswithMan() string {
    return fmt.Sprintf("%d岁的%s", j.age, j.jinLian.MakeEyeswithMan())
}

