package repository

import (
	"database/sql"
	"hotel-lister/entities"
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
	sql := "INSERT INTO country (id, name, created_at, updated_at) VALUES ($1,$2, NOW(), NOW())"

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
	sql := "SELECT * FROM hotel WHERE country_id = $1"

	rows, err := db.Query(sql, country.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []entities.Hotel

	for rows.Next() {
		var hotel entities.Hotel
		err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Image_url, &hotel.Phone, &hotel.Email, &hotel.Website, &hotel.Address, &hotel.AverageRating, &hotel.City_id, &hotel.Created_at, &hotel.Updated_at)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return hotels, nil
}
