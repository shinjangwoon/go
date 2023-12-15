package database

import (
	"log"

	"github.com/shinjangwoon/golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "root:woon@tcp(127.0.0.1:3306)/fiber_blog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	log.Println("Connection successful")

	db.AutoMigrate(new(model.Blog))

	DBConn = db
}
