package repository

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/film/domain"
	"github.com/paraizofelipe/star-planet/storage"
)

type repository struct {
	storage storage.PostgresStorage
}

func NewPostgreRepository(db *sqlx.DB) FilmRepository {
	return &repository{
		storage: storage.NewPostgres(db),
	}
}

func (r repository) Add(film domain.Film) (err error) {
	statement := `
        INSERT INTO films (
            id,
            planet_id,
            title,
            director,
            Release_date,
            created_at,
            updated_at
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6,
            $7
        );
    `
	err = r.storage.Exec(statement,
		film.ID,
		film.PlanetID,
		film.Title,
		film.Director,
		film.ReleaseDate,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) FindByID(id int) (film domain.Film, err error) {
	statement := `
        SELECT
            title,
            director,
            Release_date,
            created_at,
            updated_at
        FROM
            films
        WHERE
            id = $1;
	`
	err = r.storage.Find(statement, &film, id)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
