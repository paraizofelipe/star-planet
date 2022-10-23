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

func (r repository) RemoveByID(id int) (err error) {
	return r.storage.Exec(`DELETE FROM planets WHERE id = $1`, id)
}

func (r repository) FindByName(name string) (planet domain.Planet, err error) {
	statement := `
        SELECT
            id,
            name,
            climate,
            terrain,
            created_at,
            updated_at
        FROM
            planets
        WHERE
            name = $1;
	`
	err = r.storage.Find(statement, &planet, name)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) FindByID(id int) (planet domain.Planet, err error) {
	var films []filmsDomain.Film

	statement := `
	       SELECT
               id,
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
	if films, err = r.FindFilms(id); err != nil {
		return
	}
	planet.Films = films
	return
}

func (r repository) FindFilms(planetID int) (films []filmsDomain.Film, err error) {
	statement := `
		SELECT f.id, 
			f.title,
            f.planet_id,
			f.director,
			f.release_date,
			f.created_at,
			f.updated_at
			FROM  films as f
		INNER JOIN planets as p 
			ON p.id = f.planet_id
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
            id,
            name,
            climate,
            terrain,
            created_at,
            updated_at
        FROM
            planets;
	`
	err = r.storage.FindAll(statement, &planets)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
