package repository

import (
	"database/sql"
	"errors"
	"hotel-lister/entities"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB, user entities.User) error {
	row := db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username)
	var id int
	if err := row.Scan(&id); err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	user.Created_at = now
	user.Updated_at = now

	_, err = db.Exec("INSERT INTO users (username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", user.Username, hashedPassword, user.Email, user.Created_at, user.Updated_at)
	if err != nil {
		return err
	}
	return nil
}
