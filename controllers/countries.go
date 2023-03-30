package controllers

import (
	"hotel-lister/database"
	"hotel-lister/entities"
	"hotel-lister/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	result    gin.H
	countries entities.Country
)

func GetCountries(c *gin.Context) {
	countries, err := repository.GetCountries(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": countries,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCountries(c *gin.Context) {

	err := c.ShouldBindJSON(&countries)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCountries(database.DbConnection, countries)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Countries!",
	})
}

func UpdateCountries(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&countries)
	if err != nil {
		panic(err)
	}

	countries.ID = id

	err = repository.UpdateCountries(database.DbConnection, countries)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Country",
	})
}

func DeleteCountries(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	countries.ID = id

	err = repository.DeleteCountries(database.DbConnection, countries)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Country",
	})
}

func GetHotelByCountries(c *gin.Context) {

	countryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	country := entities.Country{ID: countryID}

	hotels, err := repository.GetHotelByCountries(database.DbConnection, country)

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
