package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/twpayne/go-geom/encoding/wkb"
	"github.com/twpayne/go-geom/encoding/wkbhex"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB
var DSN = "postgres://postgres:123456@localhost:5432/test?sslmode=disable"

type GeoJson string

type JD struct {
	ID   uint    `gorm:"primary_key" json:"Id,omitempty"`
	Geom GeoJson `gorm:"type:geometry" json:"geom"`
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *GeoJson) Scan(value interface{}) error {
	var data []byte
	gT, err := wkbhex.Decode(value.(string))
	if err != nil {
		panic(fmt.Errorf("wkbhex.Decode ERROR: %v", err))
	}
	data, err = geojson.Marshal(gT)
	if err != nil {
		panic(fmt.Errorf("geojson.Encode ERROR: %v", err))
	}
	*j = GeoJson(data)
	return nil
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j GeoJson) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	var (
		gT geom.T
	)
	err := geojson.Unmarshal([]byte(j), &gT)
	if err != nil {
		panic(fmt.Errorf("geojson.Unmarshal ERROR: %v", err))
	}
	gStr, err := wkbhex.Encode(gT, wkb.NDR)
	if err != nil {
		panic(fmt.Errorf("wkbhex.Encode ERROR: %v", err))
	}
	return gStr, nil
}

func main() {
	ConnectToDB(DSN)
	//创建
	a := JD{
		Geom: `{
	 "type": "LineString",
	 "coordinates": [
	   [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0]
	 ]
	}`,
	}
	DB.Create(&a)

	//查询
	b := JD{}
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
	DB.AutoMigrate(&JD{})
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	return nil
}
