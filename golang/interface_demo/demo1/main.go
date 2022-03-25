package main

import "fmt"

type ParentUSB struct {
	U USB
}

type USB interface {
	working()
	stop()
}

type USB1 interface {
	working()
	stop()
}

type Phone struct {}

type Cameron struct {}

type Computer struct {}

func (p Phone) working()  {
	fmt.Println("手机开始工作...")
}

func (p Phone) stop()  {
	fmt.Println("手机停止工作...")
}

func (p Cameron) working()  {
	fmt.Println("照相机开始工作...")
}

func (p Cameron) stop()  {
	fmt.Println("照相机停止工作...")
}

func (p Computer) work(usb USB)  {
	// 具有相同接口的对象(struct) 都可以作为函数参数(包含多态，高内聚低耦合的思想)
	usb.working()  // 根据上下文判断是Camera还是phone, 实现多态
	usb.stop()
}

func main() {
	c := Computer{}
	ca := Cameron{}
	ph := Phone{}
	c.work(ca) // 由于Cameron实现了USB接口，所以类型可以和USB进行匹配
	c.work(ph) // 由于Phone实现了USB接口，所以类型可以和USB进行匹配
}






