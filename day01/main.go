package main

import (
    "sync"
)

type Singleton struct{}

var singleton *Singleton



func GetInstance() *Singleton {
	return singleton
}

func init() {
    singleton = &Singleton{}
}

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}


