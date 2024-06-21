package main

import (
	"testing"
)

func TestAdapterRequest(t *testing.T) {
	// 确保 adapter 实现了目标接口
    adapter := &Adapter{Adaptee: &Adaptee{}}
    adapter.Request()
}

