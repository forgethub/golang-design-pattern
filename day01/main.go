package main

import (
    "sync"
)

type Singleton struct{
    aa int
    bb func(int) int
}

var singleton *Singleton



func GetInstance() *Singleton {
    /* return &Singleton{aa: 1, bb: func(i int) int {return i}} */
    return singleton
}

func init() {
    singleton = &Singleton{}
}

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

/*
// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
} */



func GetLazyInstance() *Singleton {
    if lazySingleton == nil {
        once.Do(func(){
            lazySingleton = &Singleton{}
        })
    }
    return lazySingleton
}



