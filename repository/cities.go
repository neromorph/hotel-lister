package repository

import (
	"database/sql"
	"errors"
	"hotel-lister/entities"
	"math"
)

func GetCities(db *sql.DB) (cities []entities.City, err error) {
	sql := "SELECT * FROM city"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var city = entities.City{}

		err = rows.Scan(&city.ID, &city.Name, &city.Country_id, &city.Created_at, &city.Updated_at)
		if err != nil {
			panic(err)
		}
		cities = append(cities, city)
	}
	return
}

func InsertCities(db *sql.DB, cities entities.City) (err error) {
	sql := "SELECT COUNT(*) FROM city WHERE id = $1"
	var count int
	err = db.QueryRow(sql, cities.ID).Scan(&count)
	if err != nil {
		return err
	}

	// If there is already a record with the same ID, return an error.
	if count > 0 {
		return errors.New("city with the given ID already exists")
	}

	sql = "INSERT INTO city (id, name, country_id, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())"

	errs := db.QueryRow(sql, cities.ID, cities.Name, &cities.Country_id)

	return errs.Err()
}

func UpdateCities(db *sql.DB, cities entities.City) (err error) {
	sql := "UPDATE city SET name = $1, updated_at = NOW() WHERE id = $2"

	errs := db.QueryRow(sql, cities.Name, cities.ID)

	return errs.Err()
}

func DeleteCities(db *sql.DB, cities entities.City) (err error) {
	sql := "DELETE FROM city WHERE id = $1"

	errs := db.QueryRow(sql, cities.ID)

	return errs.Err()
}

func GetHotelByCities(db *sql.DB, cities entities.City) (hotel []entities.Hotel, err error) {
	sql := `
	SELECT h.*, AVG(r.rating) AS average_rating
		FROM hotel h
		LEFT JOIN review r ON h.id = r.hotel_id
		WHERE h.city_id = $1
		GROUP BY h.id
`

	rows, err := db.Query(sql, cities.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []entities.Hotel

	for rows.Next() {
		var hotel entities.Hotel
		err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Image_url, &hotel.Phone, &hotel.Email, &hotel.Website, &hotel.Address, &hotel.AverageRating, &hotel.Country_id, &hotel.City_id, &hotel.Created_at, &hotel.Updated_at, &hotel.AverageRatingFloat64)
		if err != nil {
			return nil, err
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
