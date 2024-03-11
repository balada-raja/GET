package main

import (
	"log"
	"net/http"

	"github.com/balada-raja/GET/controllers/authcontroller"
	"github.com/balada-raja/GET/controllers/vehiclecontroller"
	"github.com/balada-raja/GET/controllers/ordercontroller"
	"github.com/balada-raja/GET/controllers/vendorcontroller"
	"github.com/balada-raja/GET/controllers/userscontroller"
	"github.com/balada-raja/GET/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	rest := gin.Default()
	models.ConnectDatabase()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

	//route untuk user
	rest.GET("/api/users", userscontroller.Index)
	rest.GET("/api/users/:id", userscontroller.Show)
	// rest.POST("/api/users", userscontroller.Create)
	rest.PUT("/api/users/:id", userscontroller.Update)
	rest.DELETE("/api/users/:id", userscontroller.Delete)

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

	rest.Run()
}
