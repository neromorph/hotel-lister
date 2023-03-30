package entities

import (
	"database/sql"
	"time"
)

type Hotel struct {
	ID                   int             `json:"id"`
	Name                 string          `json:"name"`
	Description          string          `json:"description"`
	Image_url            string          `json:"image_url"`
	Phone                string          `json:"phone"`
	Email                string          `json:"email"`
	Website              string          `json:"website"`
	Address              string          `json:"address"`
	AverageRating        sql.NullFloat64 `json:"average_rating"`
	AverageRatingFloat64 *float64        `json:"average_rating_float64"`
	City_id              int             `json:"city_id"`
	Country_id           int             `json:"country_id"`
	Created_at           time.Time       `json:"created_at"`
	Updated_at           time.Time       `json:"updated_at"`
}
