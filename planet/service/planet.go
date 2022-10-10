package service

import (
	"github.com/jmoiron/sqlx"
	moviesDomains "github.com/paraizofelipe/star-planet/movie/domain"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/planet/repository"
)

type PlanetService struct {
	repository repository.PlanetRepository
}

func NewService(db *sqlx.DB) Service {
	return &PlanetService{
		repository: repository.NewPostgreRepository(db),
	}
}

func (s PlanetService) Load(planetID int) (err error) {
	return
}

func (s PlanetService) Add(planet domain.Planet) (err error) {
	return s.repository.Add(planet)
}

func (s PlanetService) RemoveByID(id int) (err error) {
	return s.repository.RemoveByID(id)
}

func (s PlanetService) AddMovieToPlanet(planetID int, movieID int) (err error) {
	return
}

func (s PlanetService) FindByName(name string) (planet domain.Planet, err error) {
	return
}

func (s PlanetService) FindMovies(planetID int) ([]moviesDomains.Movie, error) {
	return s.repository.FindMovies(planetID)
}

func (s PlanetService) FindByID(id int) (planet domain.Planet, err error) {
	if planet, err = s.repository.FindByID(id); err != nil {
		return
	}
	if planet.Movies, err = s.FindMovies(id); err != nil {
		return
	}

	return
}

func (s PlanetService) FindAll() (plantes []domain.Planet, err error) {
	return
}
