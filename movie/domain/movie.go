package domain

import (
	"encoding/json"
	"time"
)

type Movie struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Director    string    `json:"director" db:"director"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (d Movie) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
