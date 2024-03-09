package main

import (
	"log"
	"net/http"

	"github.com/balada-raja/GET/controllers/authcontroller"
	"github.com/balada-raja/GET/controllers/kendaraancontroller"
	"github.com/balada-raja/GET/controllers/ordercontroller"
	"github.com/balada-raja/GET/controllers/penyedia_jasacontroller"
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
	rest.GET("/api/penyedia_jasa", penyedia_jasacontroller.Index)
	rest.GET("/api/penyedia_jasa/:id", penyedia_jasacontroller.Show)
	rest.POST("/api/penyedia_jasa", penyedia_jasacontroller.Create)
	rest.PUT("/api/penyedia_jasa/:id", penyedia_jasacontroller.Update)
	rest.DELETE("/api/penyedia_jasa", penyedia_jasacontroller.Delete)

	//route untuk kendaraan
	rest.GET("/api/kendaraan", kendaraancontroller.Index)
	rest.GET("/api/kendaraan/:id", kendaraancontroller.Show)
	rest.POST("/api/kendaraan", kendaraancontroller.Create)
	rest.PUT("/api/kendaraan/:id", kendaraancontroller.Update)
	rest.DELETE("/api/kendaraan", kendaraancontroller.Delete)

	//route untuk order
	rest.GET("/api/order", ordercontroller.Index)
	rest.GET("/api/order/:id", ordercontroller.Show)
	rest.POST("/api/order", ordercontroller.Create)
	rest.PUT("/api/order/:id", ordercontroller.Update)
	rest.DELETE("/api/order", ordercontroller.Delete)

	rest.Run()
}
