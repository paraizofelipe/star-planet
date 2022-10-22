package domain

import (
	"encoding/json"
	"time"
)

type Film struct {
	ID          int       `json:"id" db:"id"`
	PlanetID    int       `json:"planet_id" db:"planet_id"`
	Title       string    `json:"title" db:"title"`
	Director    string    `json:"director" db:"director"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (d Film) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
