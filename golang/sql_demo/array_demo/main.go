package main

import (
	"fmt"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// 结论： gorm pq不支持多维数组，需要自行实现。
// pgsql中多维数组，需保证插入数组长度一致，例如：插入二维数组 {{1,2}, {3}} --> 无法插入; {{1,2}, {3, 4}} --> 可以插入


var DB *gorm.DB
var DSN = "postgres://postgres:123456@localhost:5432/test?sslmode=disable"

type TArray struct {
	ID uint            `gorm:"primary_key" json:"Id,omitempty"`
	A  []pq.Int64Array `gorm:"type:integer[][]" json:"a"`
}

func main() {
	ConnectToDB(DSN)
	//创建
	aa := []pq.Int64Array{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(aa)
	a := TArray{A: aa}
	DB.Create(&a)

	//查询
	b := TArray{}
	DB.Debug().First(&b)
	fmt.Println(b)
}

func ConnectToDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("database connection error ", err)
	}
	sqlDB, _ := DB.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("ping err: ", err)
		return nil
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = DB.Debug().AutoMigrate(&TArray{})
	if err != nil {
		log.Fatal("AutoMigrate err: ", err)
		return nil
	}
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	return nil
}
