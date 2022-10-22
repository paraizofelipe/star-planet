package service

import (
	"log"
	"strconv"
	"strings"
	"sync"
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

func (s PlanetService) extractIDfromURL(url string) (id int, err error) {
	chunks := strings.Split(url, "/")
	lastPath := chunks[len(chunks)-2]
	if id, err = strconv.Atoi(lastPath); err != nil {
		return
	}
	return
}

func (s PlanetService) LoadFilms(wg *sync.WaitGroup, planetID int, filmURL string) {
	var (
		err          error
		filmID       int
		respFilm     swapi.RespFilm
		releaseDate  time.Time
		client       *swapi.Client = swapi.NewClient()
		filmToPlanet filmDomains.Film
	)

	defer wg.Done()

	if filmID, err = s.extractIDfromURL(filmURL); err != nil {
		log.Println(err)
	}
	if respFilm, err = client.Film(filmID); err != nil {
		log.Println(err)
	}

	if releaseDate, err = time.Parse("2006-01-02", respFilm.ReleaseDate); err != nil {
		log.Println(err)
	}

	filmToPlanet = filmDomains.Film{
		ID:          filmID,
		PlanetID:    planetID,
		ReleaseDate: releaseDate,
		Title:       respFilm.Title,
		Director:    respFilm.Director,
	}

	if err = s.filmService.Add(filmToPlanet); err != nil {
		log.Println(err)
	}
}

func (s PlanetService) Load(planetID int) (err error) {
	var (
		planet     domain.Planet
		respPlanet swapi.RespPlanet
		client     *swapi.Client = swapi.NewClient()
		wg         sync.WaitGroup
	)

	if planet, err = s.repository.FindByID(planetID); err != nil {
		return
	}

	if planet.ID != 0 {
		return
	}

	if respPlanet, err = client.Planet(planetID); err != nil {
		return
	}

	planet = domain.Planet{
		ID:      planetID,
		Name:    respPlanet.Name,
		Climate: respPlanet.Climate,
		Terrain: respPlanet.Terrain,
	}

	if err = s.Add(planet); err != nil {
		return
	}

	for _, filmURL := range respPlanet.FilmURLs {
		wg.Add(1)
		go s.LoadFilms(&wg, planetID, filmURL)
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
	if planet, err = s.repository.FindByID(id); err != nil {
		return
	}
	return
}

func (s PlanetService) FindAll() (plantes []domain.Planet, err error) {
	return s.repository.FindAll()
}
