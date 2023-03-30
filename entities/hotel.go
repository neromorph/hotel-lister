package entities

import (
	"database/sql"
	"encoding/json"
	"time"
)

type FloatNull struct {
	sql.NullFloat64
}

func (f FloatNull) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	}
	return json.Marshal(nil)
}

type Hotel struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Image_url            string    `json:"image_url"`
	Phone                string    `json:"phone"`
	Email                string    `json:"email"`
	Website              string    `json:"website"`
	Address              string    `json:"address"`
	AverageRating        FloatNull `json:"-"`
	AverageRatingFloat64 float64   `json:"average_rating"`
	City_id              int       `json:"city_id"`
	Country_id           int       `json:"country_id"`
	Created_at           time.Time `json:"created_at"`
	Updated_at           time.Time `json:"updated_at"`
}
