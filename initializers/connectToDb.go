package initializers

import (
	"github.com/balada-raja/GET/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_get"))
	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(&models.Users{})
	database.AutoMigrate(&models.Vendor{})
	database.AutoMigrate(&models.Vehicle{})
	database.AutoMigrate(&models.Order{})
	database.AutoMigrate(&models.DetailOrder{})

	DB = database
}
