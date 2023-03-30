package entities

import "time"

type Review struct {
	ID          int       `json:"id"`
	Hotel_id    int       `json:"hotel_id"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
