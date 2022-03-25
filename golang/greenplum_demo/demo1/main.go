package main

import (
	"fmt"
	"greenplum_demo/conn"
	"greenplum_demo/demo1/model"
)

func main() {
	dsn := `postgres://gpadmin:@localhost:15432/test?sslmode=disable`
	err := conn.ConnectToDB(dsn)
	if err != nil {
		fmt.Println("报错了")
	}
	//model.Automigrate()

	u := model.User{}
	//u.Add("John", 118)
	u.Get()
	//u.UpdateNameByAge("Tommy", 18)
	//u.DeleteUser(18)
	u.TransactionUser()
}
