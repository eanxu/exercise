package main

import (
	"encoding/json"
	"fmt"
	"log"
	"redis_demo/model"
	"time"
)

var DSN = `localhost:6379`

func main() {
	err := model.ConnectToRedis(DSN)
	if err != nil {
		return
	}
	for {
		s, err := model.RDB.BRPop(0, "job").Result()
		if err != nil {
			log.Fatalln("brpop error, err: ", err)
		}
		d := make(map[string]string, 0)
		err = json.Unmarshal([]byte(s[1]), &d)
		if err != nil {
			log.Fatalln("json unmarshal error, err: ", err)
		}
		fmt.Println("我从redis中取出数据: ", d)
		time.Sleep(2 * time.Second)
	}
}
