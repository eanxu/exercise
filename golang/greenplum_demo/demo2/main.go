package main

// 使用数组类型

import (
	"fmt"
	"github.com/lib/pq"
	"greenplum_demo/conn"
	"greenplum_demo/demo2/model"
)

func main() {
	dsn := `postgres://gpadmin:@localhost:15432/test?sslmode=disable`
	err := conn.ConnectToDB(dsn)
	if err != nil {
		fmt.Println("报错了")
	}
	model.Automigrate()
	model.Get(pq.Int64Array{1, 6})
}
