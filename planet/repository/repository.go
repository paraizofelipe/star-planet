package repository

import (
	movieDomain "github.com/paraizofelipe/star-planet/movie/domain"
	"github.com/paraizofelipe/star-planet/planet/domain"
)

type Reader interface {
	FindByName(name string) (domain.Planet, error)
	FindByID(id int) (domain.Planet, error)
	FindAll() ([]domain.Planet, error)
}

type Writer interface {
	Add(domain.Planet) error
	RemoveByID(id int) error
}

type Relater interface {
	FindMovies(planetID int) ([]movieDomain.Movie, error)
	AddMovieToPlanet(planetID int, movieID int) error
}

type PlanetRepository interface {
	Reader
	Writer
	Relater
}
