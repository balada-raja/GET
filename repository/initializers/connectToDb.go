package initializers

import (
	"os"

	"github.com/balada-raja/GET/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Get database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Connect to the database
	database, err := gorm.Open(mysql.Open(dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName))
	if err != nil {
		panic("Failed to connect database")
	}

	// Auto migrate models
	database.AutoMigrate(&models.Users{})
	database.AutoMigrate(&models.Vendor{})
	database.AutoMigrate(&models.Vehicle{})
	database.AutoMigrate(&models.Order{})
	database.AutoMigrate(&models.DetailOrder{})

	// Assign the database connection to global variable
	DB = database
}
