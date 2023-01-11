package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&ProducModel{})

	DB = database

}
