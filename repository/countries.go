package repository

import (
	"database/sql"
	"errors"
	"hotel-lister/entities"
	"math"
)

func GetCountries(db *sql.DB) (results []entities.Country, err error) {
	sql := "SELECT * FROM country"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var country = entities.Country{}

		err = rows.Scan(&country.ID, &country.Name, &country.Created_at, &country.Updated_at)
		if err != nil {
			panic(err)
		}
		results = append(results, country)
	}
	return
}

func InsertCountries(db *sql.DB, country entities.Country) (err error) {
	sql := "SELECT COUNT(*) FROM country WHERE id = $1"
	var count int
	err = db.QueryRow(sql, country.ID).Scan(&count)
	if err != nil {
		return err
	}

	// duplicate ID validation
	if count > 0 {
		return errors.New("country with the given ID already exists")
	}

	sql = "INSERT INTO country (id, name, created_at, updated_at) VALUES ($1,$2, NOW(), NOW())"

	errs := db.QueryRow(sql, country.ID, country.Name)

	return errs.Err()
}

func UpdateCountries(db *sql.DB, country entities.Country) (err error) {
	sql := "UPDATE country SET name = $1, updated_at = NOW() WHERE id = $2"

	errs := db.QueryRow(sql, country.Name, country.ID)

	return errs.Err()
}

func DeleteCountries(db *sql.DB, country entities.Country) (err error) {
	sql := "DELETE FROM country WHERE id = $1"

	errs := db.QueryRow(sql, country.ID)

	return errs.Err()
}

func GetHotelByCountries(db *sql.DB, country entities.Country) (hotel []entities.Hotel, err error) {
	sql := `
	SELECT h.*, COALESCE(AVG(r.rating), 0.0) AS average_rating
		FROM hotel h
		LEFT JOIN review r ON h.id = r.hotel_id
		WHERE h.country_id = $1
		GROUP BY h.id
`

	rows, err := db.Query(sql, country.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []entities.Hotel

	for rows.Next() {
		var hotel entities.Hotel
		err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Image_url, &hotel.Phone, &hotel.Email, &hotel.Website, &hotel.Address, &hotel.AverageRating, &hotel.Country_id, &hotel.City_id, &hotel.Created_at, &hotel.Updated_at, &hotel.AverageRatingFloat64)
		if err != nil {
			panic(err)
		}
		hotel.AverageRatingFloat64 = math.Round(hotel.AverageRatingFloat64*10) / 10
		hotels = append(hotels, hotel)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return hotels, nil
}
