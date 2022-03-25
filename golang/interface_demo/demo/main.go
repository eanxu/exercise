package main

import (
	"interface_demo/demo/a"
	"interface_demo/demo/root"
)

func main() {
	t := a.A{&root.Root{}}
	t.R.Save("Tom", 18)
}
