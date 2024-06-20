package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	jinLian := &JinLian{}
	j := NewJinLianAge(jinLian, 18)
	got := j.MakeEyeswithMan()
	assert.Equal(t, "18岁的金莲对你笑", got)
}
