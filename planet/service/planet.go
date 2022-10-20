package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	filmDomains "github.com/paraizofelipe/star-planet/film/domain"
	filmService "github.com/paraizofelipe/star-planet/film/service"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/planet/repository"
	"github.com/paraizofelipe/star-planet/swapi"
)

type PlanetService struct {
	filmService filmService.Service
	repository  repository.PlanetRepository
}

func NewService(db *sqlx.DB) Service {
	return &PlanetService{
		filmService: filmService.NewService(db),
		repository:  repository.NewPostgreRepository(db),
	}
}

func (s PlanetService) alreadyExists(planetID int) (exists bool, err error) {
	var planet domain.Planet

	if planet, err = s.FindByID(planetID); err != nil {
		return
	}

	exists = (planet.Name != "")
	if exists {
		return
	}

	return
}

func (s PlanetService) extractIDfromURL(url string) (id int, err error) {
	chunks := strings.Split(url, "/")
	lastPath := chunks[len(chunks)-2]
	if id, err = strconv.Atoi(lastPath); err != nil {
		return
	}
	return
}

func (s PlanetService) Load(planetID int) (err error) {
	var (
		filmID       int
		planet       domain.Planet
		respFilm     swapi.RespFilm
		filmToPlanet filmDomains.Film
		respPlanet   swapi.RespPlanet
		client       *swapi.Client = swapi.NewClient()
		releaseDate  time.Time
	)

	if respPlanet, err = client.Planet(planetID); err != nil {
		return
	}

	planet = domain.Planet{
		ID:      planetID,
		Name:    respPlanet.Name,
		Climate: respPlanet.Climate,
		Terrain: respPlanet.Terrain,
	}

	if err = s.UpdateOrAdd(planet); err != nil {
		return
	}

	for _, filmURL := range respPlanet.FilmURLs {
		if filmID, err = s.extractIDfromURL(filmURL); err != nil {
			return
		}
		if respFilm, err = client.Film(filmID); err != nil {
			return
		}

		if releaseDate, err = time.Parse("2006-01-02", respFilm.ReleaseDate); err != nil {
			return
		}

		filmToPlanet = filmDomains.Film{
			ID:          filmID,
			ReleaseDate: releaseDate,
			Title:       respFilm.Title,
			Director:    respFilm.Director,
		}

		if err = s.filmService.UpdateOrAdd(filmToPlanet); err != nil {
			return
		}

		if err = s.AddFilmsToPlanet(planetID, filmID); err != nil {
			return
		}
	}

	return
}

func (s PlanetService) Add(planet domain.Planet) (err error) {
	return s.repository.Add(planet)
}

func (s PlanetService) RemoveByID(id int) (err error) {
	return s.repository.RemoveByID(id)
}

func (s PlanetService) AddFilmsToPlanet(planetID int, filmID int) (err error) {
	return s.repository.AddFilmToPlanet(planetID, filmID)
}

func (s PlanetService) FindByName(name string) (planet domain.Planet, err error) {
	return
}

func (s PlanetService) UpdateOrAdd(planet domain.Planet) error {
	return s.repository.UpdateOrAdd(planet)
}

func (s PlanetService) FindByID(id int) (planet domain.Planet, err error) {
	if planet, err = s.repository.FindByID(id); err != nil {
		return
	}
	// if planet.Films, err = s.filmService.FindByID(id); err != nil {
	// 	return
	// }

	return
}

func (s PlanetService) FindAll() (plantes []domain.Planet, err error) {
	return
}
