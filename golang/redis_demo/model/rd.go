package model

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RDB *redis.Client

// 连接reids
func ConnectToRedis(dsn string) error {
	RDB = redis.NewClient(&redis.Options{Addr: dsn, DB: 0, PoolSize: 100})
	_, err := RDB.Ping().Result()
	if err != nil {
		return fmt.Errorf("connect redis error, err: ", err)
	}
	return nil
}

