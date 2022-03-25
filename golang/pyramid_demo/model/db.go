package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB *gorm.DB
)

func ConnectToDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln("database connection error")
		return err
	}
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	return nil
}
