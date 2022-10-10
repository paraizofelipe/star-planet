package repository

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	moviesDomain "github.com/paraizofelipe/star-planet/movie/domain"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/storage"
)

type repository struct {
	storage storage.PostgresStorage
}

func NewPostgreRepository(db *sqlx.DB) PlanetRepository {
	return &repository{
		storage: storage.NewPostgres(db),
	}
}

func (r repository) Add(planet domain.Planet) (err error) {
	statement := `
        INSERT INTO planet (
            name,
            climate,
            terranin,
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
		planet.Name,
		planet.Climate,
		planet.Terrain,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) AddMovieToPlanet(planetID int, movieID int) (err error) {
	statement := `
			INSERT INTO movie_to_planet (
				planet_id,
				movie_id,
				created_at,
				updated_at
			) VALUES (
				$1,
				$2,
				$3,
				$4
			);
	`
	err = r.storage.Exec(statement,
		planetID,
		movieID,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) RemoveByID(id int) (err error) {
	return r.storage.Exec(`DELETE FROM planets WHERE id = $1`, id)
}

func (r repository) FindByName(name string) (planet domain.Planet, err error) {
	statement := `
        SELECT
            name,
            climate,
            terranin,
            created_at,
            updated_at
        FROM
            planets
        WHERE
            email = $1;
	`
	err = r.storage.Find(statement, &planet, name)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) FindByID(id int) (planet domain.Planet, err error) {
	statement := `
        SELECT
            name,
            climate,
            terranin,
            created_at,
            updated_at
        FROM
            planets
        WHERE
            id = $1;
	`
	err = r.storage.Find(statement, &planet, id)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) FindMovies(planetID int) (movies []moviesDomain.Movie, err error) {
	statement := `
		SELECT p.id,
			m.title, 
			m.director,
			m.release_date,
			m.created_at,
			m.updated_at
			FROM  movies_to_planets as mp
		INNER JOIN planets as p 
			ON p.id = mp.planet_id
		INNER JOIN movies as m
			ON m.id = np.movie_id
		WHERE p.id = $1;
	`
	err = r.storage.FindAll(statement, &movies, planetID)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) FindAll() (planets []domain.Planet, err error) {
	statement := `
        SELECT
            name,
            climate,
            terranin,
            created_at,
            updated_at
        FROM
            planets;
	`
	err = r.storage.Find(statement, &planets, nil)
	return
}
