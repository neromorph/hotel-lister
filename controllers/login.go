package controllers

import (
	"hotel-lister/database"
	"hotel-lister/entities"
	"hotel-lister/repository"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginUser entities.User

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := repository.Login(database.DbConnection, loginUser.Username, loginUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Parse the token to get the claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Get the secret key from environment variable
		jwtKey := []byte(os.Getenv("JWT_SECRET"))
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Extract the user details from the claims
	userID := claims["id"].(float64)
	username := claims["username"].(string)
	email := claims["email"].(string)
	createdAt := claims["created_at"].(string)
	updatedAt := claims["updated_at"].(string)

	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"user_id":    userID,
		"username":   username,
		"email":      email,
		"created_at": createdAt,
		"updated_at": updatedAt,
	})
}
