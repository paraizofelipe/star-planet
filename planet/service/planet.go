package service

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	filmService "github.com/paraizofelipe/star-planet/film/service"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/planet/repository"
	"github.com/paraizofelipe/star-planet/swapi"
)

type PlanetService struct {
	filmService filmService.Service
	repository  repository.PlanetRepository
	swapiClient swapi.SWAPI
}

func NewService(db *sqlx.DB) Service {
	return &PlanetService{
		filmService: filmService.NewService(db),
		repository:  repository.NewPostgreRepository(db),
		swapiClient: swapi.NewClient(),
	}
}

func (s PlanetService) Load(planetID int) (err error) {
	var (
		planet     domain.Planet
		respPlanet swapi.RespPlanet
		wg         sync.WaitGroup
	)

	if planet, err = s.repository.FindByID(planetID); err != nil {
		return
	}

	if planet.ID != 0 {
		return
	}

	if respPlanet, err = s.swapiClient.Planet(planetID); err != nil {
		return
	}

	planet = domain.Planet{
		ID:      planetID,
		Name:    respPlanet.Name,
		Climate: respPlanet.Climate,
		Terrain: respPlanet.Terrain,
	}

	if err = s.repository.Add(planet); err != nil {
		return
	}

	for _, filmURL := range respPlanet.FilmURLs {
		wg.Add(1)

		go func(planetID int, filmURL string) {
			defer wg.Done()

			if err := s.filmService.LoadFilms(planetID, filmURL); err != nil {
				log.Println(err)
			}
		}(planetID, filmURL)
	}

	wg.Wait()

	return
}

func (s PlanetService) Add(planet domain.Planet) (err error) {
	return s.repository.Add(planet)
}

func (s PlanetService) RemoveByID(id int) (err error) {
	return s.repository.RemoveByID(id)
}

func (s PlanetService) FindByName(name string) (planet domain.Planet, err error) {
	return s.repository.FindByName(name)
}

func (s PlanetService) FindByID(id int) (planet domain.Planet, err error) {
	return s.repository.FindByID(id)
}

func (s PlanetService) FindAll() (planets []domain.Planet, err error) {
	return s.repository.FindAll()
}
