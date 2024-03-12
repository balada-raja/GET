package initializers

import (
	"os"

	"github.com/balada-raja/GET/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

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
