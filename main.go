package main

import (
	// "log"
	// "net/http"

	"github.com/balada-raja/GET/controllers/authcontroller"
	"github.com/balada-raja/GET/controllers/ordercontroller"
	"github.com/balada-raja/GET/controllers/vehiclecontroller"
	"github.com/balada-raja/GET/controllers/vendorcontroller"
	"github.com/balada-raja/GET/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	rest := gin.Default()

	rest.POST("/login", authcontroller.Login)
	rest.POST("/register", authcontroller.Register)
	rest.GET("/validate", authcontroller.Validate)
	rest.GET("/logout", authcontroller.Logout)

	//route untuk penyedia jasa
	rest.GET("/api/vendor", vendorcontroller.Index)
	rest.GET("/api/vendor/:id", vendorcontroller.Show)
	rest.POST("/api/vendor", vendorcontroller.Create)
	rest.PUT("/api/vendor/:id", vendorcontroller.Update)
	rest.DELETE("/api/vendor", vendorcontroller.Delete)

	//route untuk kendaraan
	rest.GET("/api/vehicle", vehiclecontroller.Index)
	rest.GET("/api/vehicle/:id", vehiclecontroller.Show)
	rest.POST("/api/vehicle", vehiclecontroller.Create)
	rest.PUT("/api/vehicle/:id", vehiclecontroller.Update)
	rest.DELETE("/api/vehicle", vehiclecontroller.Delete)

	//route untuk order
	rest.GET("/api/order", ordercontroller.Index)
	rest.GET("/api/order/:id", ordercontroller.Show)
	rest.POST("/api/order", ordercontroller.Create)
	rest.PUT("/api/order/:id", ordercontroller.Update)
	rest.DELETE("/api/order", ordercontroller.Delete)

	rest.Run(":8080")
}
