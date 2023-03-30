package controllers

import (
	"hotel-lister/database"
	"hotel-lister/entities"
	"hotel-lister/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user entities.User

	// Parse the JSON data from the request body into the user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Register function with the parsed user data
	if err := repository.Register(database.DbConnection, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
