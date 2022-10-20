package repository

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	filmsDomain "github.com/paraizofelipe/star-planet/film/domain"
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
        INSERT INTO planets (
            id,
            name,
            climate,
            terrain,
            created_at,
            updated_at
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6
        );
    `
	err = r.storage.Exec(statement,
		planet.ID,
		planet.Name,
		planet.Climate,
		planet.Terrain,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) AddFilmToPlanet(planetID int, filmID int) (err error) {
	statement := `
			INSERT INTO films_to_planets (
				planet_id,
				film_id,
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
		filmID,
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
            terrain,
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
            terrain,
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

func (r repository) FindFilms(planetID int) (films []filmsDomain.Film, err error) {
	statement := `
		SELECT p.id,
			m.title, 
			m.director,
			m.release_date,
			m.created_at,
			m.updated_at
			FROM  films_to_planets as mp
		INNER JOIN planets as p 
			ON p.id = mp.planet_id
		INNER JOIN films as m
			ON m.id = mp.film_id
		WHERE p.id = $1;
	`
	err = r.storage.FindAll(statement, &films, planetID)
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
            terrain,
            created_at,
            updated_at
        FROM
            planets;
	`
	err = r.storage.Find(statement, &planets, nil)
	return
}

func (r repository) UpdateOrAdd(planet domain.Planet) (err error) {
	statement := `
        INSERT INTO planets (
            id,
            name,
            climate,
            terrain,
            created_at,
            updated_at
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6
        ) ON CONFLICT (id)
        DO UPDATE SET 
            name = $2,
            climate = $3,
            terrain = $4,
            created_at = $5,
            updated_at = $6
        WHERE
            planets.id = $1;
    `
	err = r.storage.Exec(statement,
		planet.ID,
		planet.Name,
		planet.Climate,
		planet.Terrain,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}
