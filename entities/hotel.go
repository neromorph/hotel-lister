package entities

import "time"

type Hotel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image_url   string    `json:"image_url"`
	Phone       int       `json:"phone"`
	Email       string    `json:"email"`
	Website     string    `json:"website"`
	Address     string    `json:"address"`
	Rating      int       `json:"rating"`
	City_id     int       `json:"city_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
