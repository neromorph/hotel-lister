package entities

import "time"

type City struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Country_id int       `json:"country_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
