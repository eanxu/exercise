package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB
var DSN = "postgres://postgres:123456@localhost:5432/test?sslmode=disable"

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
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	return nil
}
