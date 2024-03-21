package delivery

import (
	"time"

	"github.com/balada-raja/GET/delivery/controllers/authcontroller"
	"github.com/balada-raja/GET/delivery/controllers/ordercontroller"
	"github.com/balada-raja/GET/delivery/controllers/vehiclecontroller"
	"github.com/balada-raja/GET/delivery/controllers/vendorcontroller"
	"github.com/balada-raja/GET/delivery/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute API
func SetupRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))
	
	router.POST("/login", authcontroller.Login)
	router.POST("/register", authcontroller.Register)
	router.GET("/logout", authcontroller.Logout)
	router.GET("/validate", middlewares.RequireAuth, authcontroller.Validate)

	// Rute untuk penyedia jasa
	router.GET("/api/vendor", vendorcontroller.Index)
	router.GET("/api/vendor/:id", vendorcontroller.Show)
	router.POST("/api/vendor", vendorcontroller.Create)
	router.PUT("/api/vendor/:id", vendorcontroller.Update)
	router.DELETE("/api/vendor", vendorcontroller.Delete)

	// Rute untuk kendaraan
	router.GET("/api/vehicle", vehiclecontroller.Index)
	router.GET("/api/vehicle/:id", vehiclecontroller.Show)
	router.GET("/api/vehicle/name", vehiclecontroller.FindVehicleByName)
	router.GET("/api/vehicle/vehicle_type", vehiclecontroller.FindVehicleByVehicleType)
	router.GET("/api/vehicle/transmission", vehiclecontroller.FindVehicleByTransmission)
	router.GET("/api/vehicle/price", vehiclecontroller.FindVehicleByPriceRange)
	router.POST("/api/vehicle", vehiclecontroller.Create)
	router.PUT("/api/vehicle/:id", vehiclecontroller.Update)
	router.DELETE("/api/vehicle", vehiclecontroller.Delete)

	// Rute untuk order
	router.GET("/api/order", ordercontroller.Index)
	router.GET("/api/order/:id", ordercontroller.Show)
	router.POST("/api/order", ordercontroller.Create)
	router.PUT("/api/order/:id", ordercontroller.Update)
	router.DELETE("/api/order", ordercontroller.Delete)
}
