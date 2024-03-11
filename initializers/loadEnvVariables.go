package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error load env : %v", err)
	}
}