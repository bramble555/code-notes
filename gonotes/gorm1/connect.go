package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
func init() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "gorm_test"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&Local&timeout=%s",
		username, password, host, port, dbname, timeout,
	)
	var err error
	// 配置日志
	myLog := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: myLog,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}

