package main

import (
	"fmt"
	"strconv"
	"time"
)

// 结构体嵌套接口
// 结构体里嵌套接口的目的：
// 当前结构体实例可以用所有实现了该接口的其他结构体来初始化（即使他们的属性不完全一致）

// 接口：一组方法的集合
// OpenCloser 接口定义两个方法 返回 error
type OpenCloser interface {
	Open() error
	Close() error
}

//type Locker interface {
//	Lock() error
//	Unlock() error
//}

type Door struct {
	open bool // 门的状态是否开启
	lock bool // 门的状态是否上锁
}

func (d *Door) Open() error {
	fmt.Println("door open...")
	d.open = true
	return nil
}

func (d *Door) Close() error {
	fmt.Println("door close...")
	d.open = false
	return nil
}

type AutoDoor struct {
	OpenCloser        // 匿名接口
	delay      int    // 延迟多长时间开启
	msg        string // 自动开启时的警报
}

func (a *AutoDoor) Open() error {
	fmt.Println("Open after " + strconv.Itoa(a.delay) + " seconds")
	time.Sleep(time.Duration(a.delay) * time.Second)
	fmt.Println("Door is opening:" + a.msg)
	return nil
}

func main() {
	door := &AutoDoor{&Door{false, false}, 3, "warning"}
	door.Open()
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println(v)
	}

	door.OpenCloser.Open()
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println(v)
	}

	door.Close()
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println(v)
	}
}
