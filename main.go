package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/balada-raja/GET/repository/initializers"
	"github.com/balada-raja/GET/delivery"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	RunServer()
}

// RunServer memulai server HTTP
func RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := gin.Default()
	delivery.SetupRoutes(router)
	router.Run(fmt.Sprintf(":%s", port))
}

