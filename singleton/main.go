package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	value string
}

var obj *singleton
var once sync.Once

func getInstance(value string) *singleton {
	once.Do(func() {

		obj = &singleton{
			value: value,
		}
	})
	return obj
}

func main() {

	fmt.Println(getInstance("fist call init").value)
	fmt.Println(getInstance("send call").value) // out put will "fist call init"
}
