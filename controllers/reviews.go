package controllers

import (
	"hotel-lister/database"
	"hotel-lister/entities"
	"hotel-lister/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllReviews(c *gin.Context) {
	var (
		result gin.H
	)

	reviews, err := repository.GetAllReviews(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": reviews,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertReviews(c *gin.Context) {
	var reviews entities.Review

	err := c.ShouldBindJSON(&reviews)
	if err != nil {
		panic(err)
	}

	err = repository.InsertReviews(database.DbConnection, reviews)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Review!",
	})
}

func UpdateReviews(c *gin.Context) {
	var reviews entities.Review
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&reviews)
	if err != nil {
		panic(err)
	}

	reviews.ID = id

	err = repository.UpdateReviews(database.DbConnection, reviews)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Review",
	})
}

func DeleteReviews(c *gin.Context) {
	var reviews entities.Review
	id, err := strconv.Atoi(c.Param("id"))

	reviews.ID = id

	err = repository.DeleteReviews(database.DbConnection, reviews)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Review",
	})
}

func GetReviewByHotel(c *gin.Context) {
	var (
		result gin.H
	)

	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	hotel := entities.Hotel{ID: hotelID}

	hotels, err := repository.GetReviewByHotel(database.DbConnection, hotel)

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
