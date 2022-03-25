package main

import (
	"fmt"
	"pyramid_demo/model"
	"pyramid_demo/v1/pyramid"
	"time"
)

// v1 版本

// 建立影像金字塔

// 总耗时：1343s

var DSN = `postgresql://postgres:123456@localhost:5432/vector_test?sslmode=disable`

func main() {
	start := time.Now()
	err := model.ConnectToDB(DSN)
	if err != nil {
		return
	}
	pyramid.Pyramid()
	dur := time.Since(start)
	fmt.Println("共耗时: ", dur.Seconds())
}
