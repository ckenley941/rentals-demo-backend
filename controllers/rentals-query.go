package controllers

import (
	//"github.com/ckenley941/rentals-demo-backend/data"
	"net/http"

	"github.com/ckenley941/rentals-demo-backend/data"
	"github.com/ckenley941/rentals-demo-backend/filters"
	"github.com/ckenley941/rentals-demo-backend/initializers"
	"github.com/ckenley941/rentals-demo-backend/models"
	"github.com/gin-gonic/gin"
)

func GetRentalsById(c *gin.Context) {
	var db = initializers.ConnectToDB()

	var rental = data.Rentals{}

	db.Limit(1).Find(&rental)
	//Get id off url
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": rental,
	})
}

func GetRentalsByFilter(c *gin.Context) {
	var db = initializers.ConnectToDB()
	//Get filters from the url
	var rentalFilter = filters.RentalFilter{}

	//Not sure if this will bind directly to the object but that is the idea rather than setting each property individually
	c.ShouldBindUri(&rentalFilter)

	// rentalFilter := filters.RentalFilter{
	// 	PriceMin: c.Param("price_min"),
	// 	Limit:    c.Param("limit"),
	// 	Offset:   c.Param("offset"),
	// 	Ids:      c.Param("ids"),
	// 	Near:     c.Param("near"),
	// 	Sort:     c.Param("sort"),
	// }

	var rentals = []data.Rentals{}
	//Determine which filters to narrow down our data on
	//This method could be done more efficiently and/or dynamically but this was the simplest solution in shortest time given my limited knowledge of GoLang
	if len(rentalFilter.PriceMin) > 0 {
		db.Where("price_per_day > ?", rentalFilter.PriceMin).Find(&rentals)
	}

	if len(rentalFilter.PriceMax) > 0 {
		db.Where("price_per_day < ?", rentalFilter.PriceMax).Find(&rentals)
	}

	if rentalFilter.Limit > 0 {
		db.Limit(rentalFilter.Limit).Find(&rentals)
	}

	if rentalFilter.Offset > 0 {
		db.Offset(rentalFilter.Offset).Find(&rentals)
	}

	if len(rentalFilter.Ids) > 0 {
		db.Where("id IN ?", rentalFilter.Ids).Find(&rentals)
	}

	if len(rentalFilter.Near) > 0 {
		//Would need to write some sort of distance function using lng and lat coordinates to determine what is nearby
	}

	if len(rentalFilter.Sort) > 0 {
		//Would want this to be a little smarter and ensure user passes in a sortable field on the object
		db.Order(rentalFilter.Sort).Find(&rentals)
	}

	p := models.Rental{Make: "Rental Make 1", Model: "Rental Model 1", Location: models.Location{City: "Frisco21233"}}
	c.JSON(http.StatusOK, gin.H{
		"message": p.Location.City,
	})
}
