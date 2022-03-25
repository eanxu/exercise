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
	// 链接redis
	err := model.ConnectToRedis(DSN)
	if err != nil {
		return
	}
	// 将多组json转存入列表队列中
	a := []map[string]string{
		{"a":"qwer"}, {"b":"qwer"}, {"c":"qwer"}, {"d":"qwer"}, {"e":"qwer"}, {"f":"qwer"},
		{"g":"qwer"}, {"h":"qwer"}, {"i":"qwer"}, {"j":"qwer"}, {"k":"qwer"}, {"l":"qwer"},
		{"m":"qwer"}, {"n":"qwer"}, {"o":"qwer"}, {"p":"qwer"}, {"q":"qwer"}, {"r":"qwer"},
	}
	for _, i := range a {
		b, err := json.Marshal(i)
		if err != nil {
			log.Fatalln("json marsh error, err: ", err)
		}
		err = model.RDB.LPush("job", string(b)).Err()
		if err != nil {
			log.Fatalln("redis rpush error, err: ", err)
		}
		// 每0.5s塞一个去redis
		fmt.Println("我已将数据放入redis中, 数据: ", string(b))
		time.Sleep(500 * time.Millisecond)
	}
}
