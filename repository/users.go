package repository

import (
	"database/sql"
	"fmt"
	"hotel-lister/entities"
)

func GetAllUsers(db *sql.DB) ([]entities.User, error) {
	rows, err := db.Query("SELECT id, username, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Created_at, &user.Updated_at); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteUser(db *sql.DB, userID int64) error {
	res, err := db.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", userID)
	}

	return nil
}
