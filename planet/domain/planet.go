package domain

import (
	"encoding/json"
	"time"

	"github.com/paraizofelipe/star-planet/movie/domain"
)

type Planet struct {
	ID        int            `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Climate   string         `json:"climate" db:"climate"`
	Terrain   string         `json:"terrain" db:"terrain"`
	Movies    []domain.Movie `json:"movies" db:"movies"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}

func (d Planet) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
