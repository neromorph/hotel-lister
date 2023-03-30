package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Verify the signing key
		jwtKey := []byte(os.Getenv("JWT_SECRET"))
		return jwtKey, nil
	})

	if err != nil {
		log.Println("Error while parsing token:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Salah"})
		return
	}

	// Verify the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("user_id", claims["id"])
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Salah lagi"})
		return
	}
}
