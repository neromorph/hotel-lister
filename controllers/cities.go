package controllers

import (
	"hotel-lister/database"
	"hotel-lister/entities"
	"hotel-lister/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	var (
		result gin.H
	)

	cities, err := repository.GetCities(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": cities,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCities(c *gin.Context) {
	var cities entities.City

	err := c.ShouldBindJSON(&cities)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCities(database.DbConnection, cities)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert City!",
	})
}

func UpdateCities(c *gin.Context) {
	var cities entities.City
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&cities)
	if err != nil {
		panic(err)
	}

	cities.ID = id

	err = repository.UpdateCities(database.DbConnection, cities)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update City",
	})
}

func DeleteCities(c *gin.Context) {
	var cities entities.City
	id, err := strconv.Atoi(c.Param("id"))

	cities.ID = id

	err = repository.DeleteCities(database.DbConnection, cities)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete City",
	})
}

func GetHotelByCities(c *gin.Context) {
	var (
		result gin.H
	)

	cityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid City ID"})
		return
	}

	city := entities.City{ID: cityID}

	hotels, err := repository.GetHotelByCities(database.DbConnection, city)

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
