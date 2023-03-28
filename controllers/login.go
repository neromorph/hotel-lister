package controllers

import (
	"database/sql"
	"hotel-lister/entities"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB, username string, password string) (string, error) {
	// Retrieve the user information from the database based on the username
	row := db.QueryRow("SELECT id, username, password, created_at, updated_at FROM users WHERE username=$1", username)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Created_at, &user.Updated_at)
	if err != nil {
		return "", err
	}
	// Verify the password against the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	// Generate a JWT token for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"username":   user.Username,
		"created_at": user.Created_at,
		"updated_at": user.Updated_at,
	})
	jwtToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
