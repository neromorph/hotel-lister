package repository

import (
	"database/sql"
	"hotel-lister/entities"
)

func GetRoomByHotels(db *sql.DB, hotel entities.Hotel) (room []entities.Room, err error) {
	sql := "SELECT * FROM room WHERE hotel_id = $1"

	rows, err := db.Query(sql, hotel.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []entities.Room

	for rows.Next() {
		var room entities.Room
		err = rows.Scan(&room.ID, &room.Name, &room.Description, &room.Image_url, &room.Price, &room.Hotel_id, &room.Created_at, &room.Updated_at)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func InsertRooms(db *sql.DB, room entities.Room) (err error) {

	sql := "INSERT INTO room (name, description, image_url, price, hotel_id, created_at, updated_at) VALUES ($1,$2,$3, $4, $5, NOW(), NOW())"

	errs := db.QueryRow(sql, &room.Name, &room.Description, &room.Image_url, &room.Price, &room.Hotel_id)

	return errs.Err()
}

func UpdateRooms(db *sql.DB, room entities.Room) (err error) {

	sql := "UPDATE room SET name = $1, description = $2, image_url = $3, price = $4, hotel_id = $5, updated_at = NOW() WHERE id = $6"

	errs := db.QueryRow(sql, &room.Name, &room.Description, &room.Image_url, &room.Price, &room.Hotel_id, &room.ID)

	return errs.Err()
}

func DeleteRooms(db *sql.DB, room entities.Room) (err error) {
	sql := "DELETE FROM room WHERE id = $1"

	errs := db.QueryRow(sql, room.ID)

	return errs.Err()
}
