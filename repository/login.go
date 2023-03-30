package repository

import (
	"database/sql"
	"hotel-lister/entities"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB, username string, password string) (string, error) {

	row := db.QueryRow("SELECT id, username, password, email, created_at, updated_at FROM users WHERE username=$1", username)

	var user entities.User

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"created_at": user.Created_at,
		"updated_at": user.Updated_at,
		"exp":        expirationTime.Unix(),
	})

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	jwtToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
