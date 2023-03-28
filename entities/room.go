package entities

import "time"

type Room struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image_url   string    `json:"image_url"`
	Price       string    `json:"price"`
	Hotel_id    int       `json:"hotel_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
