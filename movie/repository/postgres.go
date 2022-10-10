package repository

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/movie/domain"
	"github.com/paraizofelipe/star-planet/storage"
)

type repository struct {
	storage storage.PostgresStorage
}

func NewPostgreRepository(db *sqlx.DB) MovieRepository {
	return &repository{
		storage: storage.NewPostgres(db),
	}
}

func (r repository) Add(movie domain.Movie) (err error) {
	statement := `
        INSERT INTO movies (
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
            $5
        );
    `
	err = r.storage.Exec(statement,
		movie.Title,
		movie.Director,
		movie.ReleaseDate,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) FindByID(id int) (movie domain.Movie, err error) {
	statement := `
        SELECT
            title,
            director,
            Release_date,
            created_at,
            updated_at
        FROM
            movies
        WHERE
            id = $1;
	`
	err = r.storage.Find(statement, &movie, id)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
