package model

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"cuttile_demo/internal/logger"
)

var (
	DB *gorm.DB
)

func ConnectToDB(driverName, dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Logger.Error("database connection error",
			zap.String("dsn", dsn),
			zap.Error(err))
		return err
	}
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	return nil
}
