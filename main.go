package main

import (
	"fmt"

	"github.com/ckenley941/rentals-demo-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(("App starting"))
	p := models.Rental{Make: "Rental Make 1", Model: "Rental Model 1", Location: models.Location{City: "Frisco"}}
	r := gin.Default()

	//Test API ping and that Rental property is populated
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": p.Location.City,
		})
	})
	r.Run()
}
