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

func GetRoomByHotels(c *gin.Context) {
	var (
		result gin.H
	)

	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hotel ID"})
		return
	}

	hotel := entities.Hotel{ID: hotelID}

	rooms, err := repository.GetRoomByHotels(database.DbConnection, hotel)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": rooms,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertRooms(c *gin.Context) {
	var rooms entities.Room

	err := c.ShouldBindJSON(&rooms)
	if err != nil {
		panic(err)
	}

	if !helpers.IsValidURL(rooms.Image_url) {
		errors = append(errors, "Invalid image URL")
	}

	err = repository.InsertRooms(database.DbConnection, rooms)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Room added",
	})
}

func UpdateRooms(c *gin.Context) {
	var rooms entities.Room
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&rooms)
	if err != nil {
		panic(err)
	}

	rooms.ID = id

	if !helpers.IsValidURL(rooms.Image_url) {
		errors = append(errors, "Invalid image URL")
	}

	err = repository.UpdateRooms(database.DbConnection, rooms)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Room detail updated",
	})
}

func DeleteRooms(c *gin.Context) {
	var rooms entities.Room
	id, err := strconv.Atoi(c.Param("id"))

	rooms.ID = id

	err = repository.DeleteRooms(database.DbConnection, rooms)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Room deleted",
	})
}
