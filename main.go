package main

import (
	"github.com/ckenley941/rentals-demo-backend/controllers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	server = gin.Default()
}

func main() {
	router := server.Group("/api")
	// p := models.Rental{Make: "Rental Make 1", Model: "Rental Model 1", Location: models.Location{City: "Frisco"}}
	// //Test API ping and that Rental property is populated
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": p.Location.City,
	// 	})
	// })
	router.GET("/rentals", controllers.GetRentalsById)

	server.Run()
}
