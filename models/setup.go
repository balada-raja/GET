package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_get"))
	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(&Users{})
	database.AutoMigrate(&Vendor{})
	database.AutoMigrate(&Vehicle{})
	database.AutoMigrate(&Order{})
	database.AutoMigrate(&DetailOrder{})

	DB = database
}
