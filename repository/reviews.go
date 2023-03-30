package repository

import (
	"database/sql"
	"errors"
	"hotel-lister/entities"
	"hotel-lister/helpers"
)

func GetAllReviews(db *sql.DB) (reviews []entities.Review, err error) {
	sql := "SELECT * FROM review"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var review = entities.Review{}

		err = rows.Scan(&review.ID, &review.Hotel_id, &review.Rating, &review.Description, &review.Created_at, &review.Updated_at)
		if err != nil {
			panic(err)
		}
		reviews = append(reviews, review)
	}
	return
}

func InsertReviews(db *sql.DB, reviews entities.Review) (err error) {

	if !helpers.RatingLimit(reviews.Rating) {
		return errors.New("rating value must be between 1 and 5")
	}

	sql := "INSERT INTO review (hotel_id, rating, description, created_at, updated_at) VALUES ($1,$2,$3, NOW(), NOW())"

	errs := db.QueryRow(sql, &reviews.Hotel_id, &reviews.Rating, &reviews.Description)

	return errs.Err()
}

func UpdateReviews(db *sql.DB, reviews entities.Review) (err error) {
	if !helpers.RatingLimit(reviews.Rating) {
		return errors.New("rating value must be between 1 and 5")
	}

	sql := "UPDATE review SET rating = $1, description = $2, updated_at = NOW() WHERE id = $3"

	errs := db.QueryRow(sql, reviews.Rating, &reviews.Description, reviews.ID)

	return errs.Err()
}

func DeleteReviews(db *sql.DB, reviews entities.Review) (err error) {
	sql := "DELETE FROM review WHERE id = $1"

	_, err = db.Exec(sql, reviews.ID)

	if err != nil {
		return err
	}

	return err
}

func GetReviewByHotel(db *sql.DB, hotel entities.Hotel) (review []entities.Review, err error) {
	sql := "SELECT * FROM review WHERE hotel_id = $1"

	rows, err := db.Query(sql, &hotel.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []entities.Review

	for rows.Next() {
		var review entities.Review
		err := rows.Scan(&review.ID, &review.Hotel_id, &review.Rating, &review.Description, &review.Created_at, &review.Updated_at)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
