package controllers

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"hotel-lister/entities"
	"time"
)

func Register(db *sql.DB, user entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	user.Created_at = now
	user.Updated_at = now

	_, err = db.Exec("INSERT INTO user (username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", user.Username, hashedPassword, user.Email, user.Created_at, user.Updated_at)
	if err != nil {
		return err
	}
	return nil
}
