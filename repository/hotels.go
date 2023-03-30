package repository

import (
	"database/sql"
	"hotel-lister/entities"
	"math"
)

func GetAllHotels(db *sql.DB) (hotels []entities.Hotel, err error) {
	sql := `
		SELECT h.*, COALESCE(AVG(r.rating), 0.0) AS average_rating
		FROM hotel h
		LEFT JOIN review r ON h.id = r.hotel_id
		GROUP BY h.id
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var hotel = entities.Hotel{}

		err = rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Image_url, &hotel.Phone, &hotel.Email, &hotel.Website, &hotel.Address, &hotel.AverageRating, &hotel.Country_id, &hotel.City_id, &hotel.Created_at, &hotel.Updated_at, &hotel.AverageRatingFloat64)
		if err != nil {
			panic(err)
		}
		hotel.AverageRatingFloat64 = math.Round(hotel.AverageRatingFloat64*10) / 10

		hotels = append(hotels, hotel)
	}
	return
}

func InsertHotels(db *sql.DB, hotel entities.Hotel) (err error) {

	sql := "INSERT INTO hotel (name, description, image_url, phone, email, website, address, city_id, country_id, created_at, updated_at) VALUES ($1,$2,$3, $4, $5, $6, $7, $8, $9, NOW(), NOW())"

	errs := db.QueryRow(sql, &hotel.Name, &hotel.Description, &hotel.Image_url, &hotel.Phone, &hotel.Email, &hotel.Website, &hotel.Address, &hotel.City_id, &hotel.Country_id)

	return errs.Err()
}

func UpdateHotels(db *sql.DB, hotel entities.Hotel) (err error) {

	sql := "UPDATE hotel SET name = $1, description = $2, image_url = $3, phone = $4, email = $5, website = $6, address = $7, city_id = $8, country_id = $9, updated_at = NOW() WHERE id = $10"

	errs := db.QueryRow(sql, &hotel.Name, &hotel.Description, &hotel.Image_url, &hotel.Phone, &hotel.Email, &hotel.Website, &hotel.Address, &hotel.City_id, &hotel.Country_id, &hotel.ID)

	return errs.Err()
}

func DeleteHotels(db *sql.DB, hotel entities.Hotel) (err error) {
	sql := "DELETE FROM hotel WHERE id = $1"

	errs := db.QueryRow(sql, hotel.ID)

	return errs.Err()
}
