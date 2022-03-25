package main

import "fmt"

// 测试接口指针与接口的区别

type person struct {
	name string
	age  int
}

type father struct {
	s Save
}

type fatherInterface interface {
	Get()
}

func (f *father) Get() {
	f.s.save("TTTTT", 118)
	fmt.Println("father get")
}

type grandfather struct {
	g fatherInterface
}

type Save interface {
	save(name string, age int)
}

func (p *person) save(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("save: ", p)
}

func main() {
	f := father{s: &person{}}
	f.s.save("Tom", 18)

	g := grandfather{g: &f}
	g.g.Get()
}

