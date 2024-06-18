package main

import (
	"testing"


	_ "github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	/* assert.Equal(t, GetInstance(), GetInstance()) */
    if GetInstance() != GetInstance() {
        t.Errorf("test fail")
    }
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestLazyGetInstance(t *testing.T) {
    if GetLazyInstance() != GetLazyInstance() {
        t.Errorf("test fail")
    }
}

func BenchmarkLazyGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
