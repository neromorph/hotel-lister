package controllers

import (
	"hotel-lister/database"
	"hotel-lister/entities"
	"hotel-lister/helpers"
	"hotel-lister/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	errors []string
)

func GetAllHotels(c *gin.Context) {
	var (
		result gin.H
	)

	hotels, err := repository.GetAllHotels(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": hotels,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertHotels(c *gin.Context) {
	var (
		hotel entities.Hotel
	)

	err := c.ShouldBindJSON(&hotel)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if !helpers.IsValidURL(hotel.Image_url) {
		errors = append(errors, "Invalid image URL")
	}

	if !helpers.IsValidURL(hotel.Website) {
		errors = append(errors, "Invalid image URL")
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}

	err = repository.InsertHotels(database.DbConnection, hotel)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Hotel added",
	})
}

func UpdateHotels(c *gin.Context) {
	var (
		hotel  entities.Hotel
		errors []string
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&hotel)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if !helpers.IsValidURL(hotel.Image_url) {
		errors = append(errors, "Invalid image URL")
	}
	if !helpers.IsValidURL(hotel.Website) {
		errors = append(errors, "Invalid image URL")
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}

	hotel.ID = id

	err = repository.UpdateHotels(database.DbConnection, hotel)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Hotel detail updated",
	})
}

func DeleteHotels(c *gin.Context) {
	var hotel entities.Hotel
	id, err := strconv.Atoi(c.Param("id"))

	hotel.ID = id

	err = repository.DeleteHotels(database.DbConnection, hotel)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Hotel deleted",
	})
}
