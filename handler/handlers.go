package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayalova/kodix-test/model"
	"github.com/kayalova/kodix-test/repository"
)

// GetAll return all cars
func GetAll(c *gin.Context) {
	// params := c.Request.URL.Query()
	params, _ := c.Get("searchParams")

	var cars []model.Car
	// fmt.Println(isExists)
	var r = params.(map[string]interface{})

	fmt.Println(params)
	cars = repository.GetAll(r)

	c.JSON(http.StatusOK, gin.H{"cars": cars})
	return
}

// CreateOne  creates a car
func CreateOne(c *gin.Context) {
	var car model.Car
	c.Bind(&car)

	repository.Create(car)

	c.JSON(http.StatusCreated, gin.H{"message": "Car was successfully created"})
}

//UpdateOne updates a car
func UpdateOne(c *gin.Context) {
	// id := c.PostForm("id")

	var car model.Car

	c.Bind(&car)

	fmt.Println(car)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
	return
}

// DeleteOne deletes a car
func DeleteOne(c *gin.Context) {
	id, _ := c.GetQuery("id")

	fmt.Println(id)
	repository.DeleteOne(id)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
	return
}
