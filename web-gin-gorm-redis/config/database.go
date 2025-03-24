package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql" // 替换为mysql驱动
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	// MySQL的DSN格式："user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName)

	// dsn := "root:123456@tcp(127.0.0.1:3306)/web-gin-gorm-redis?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	log.Println("Database connected")
}
