package a

import "interface_demo/demo/root"

type A struct {
	R root.RootInterface
}

type AInterface interface {
	Save(name string, age int)
}
