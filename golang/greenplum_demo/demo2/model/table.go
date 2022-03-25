package model

import (
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"greenplum_demo/conn"
	"log"
)

type User struct {
	ID     uint          `gorm:"primary_key" json:"id"`
	Name   string        `gorm:"type:varchar(100)" json:"name"`
	Age    uint          `gorm:"type:integer" json:"age"`
	LogIds pq.Int64Array `gorm:"type:integer[]" json:"log_ids"`
}

type Log struct {
	ID  uint   `gorm:"primary_key" json:"id"`
	Log string `gorm:"type:text" json:"log"`
}

//This migrate all tables
func Automigrate() {
	err := conn.DB.AutoMigrate(&User{}, &Log{})
	if err != nil {
		panic(fmt.Errorf("Automigrate table error, err: %v", err))
	}
}

func Add() {
	err := conn.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&User{Name: "Tom", Age: 18, LogIds: pq.Int64Array{1, 2, 3, 4, 5}}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Log{Log: "早上起床"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Log{Log: "穿衣服"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Log{Log: "叠被子"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Log{Log: "洗漱"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Log{Log: "吃饭"}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Get(array pq.Int64Array)  {
	us := []User{}
	err := conn.DB.Model(&User{}).Where("log_ids && ?", array).Find(&us).Error
	if err != nil {
		log.Fatalf("get logs error, err:", err)
	}
	fmt.Println(us)
}
