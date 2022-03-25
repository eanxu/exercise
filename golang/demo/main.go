package main

import (
	"errors"
	"fmt"
)

type a struct {}

type n interface {
	name()
	ttt()
}

func (receiver *a) name() {
	fmt.Println("啦啦啦了")
}

func (receiver *a) ttt() {
	fmt.Println("aaaaaa")
}

func errorTest() error {
	err := errors.New("我是一个错误")
	return fmt.Errorf("lalalal err: %v", err)
}


func main() {

	//var t a
	//var c n
	//c = t
	//t.ttt()
	//t.name()
	err := errorTest()
	if err !=nil {
		panic(err)
	}
}
