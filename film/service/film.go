package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/film/domain"
	"github.com/paraizofelipe/star-planet/film/repository"
	"github.com/paraizofelipe/star-planet/swapi"
)

type FilmService struct {
	swapiClient swapi.SWAPI
	repository  repository.FilmRepository
}

func NewService(db *sqlx.DB) Service {
	return &FilmService{
		swapiClient: swapi.NewClient(),
		repository:  repository.NewPostgreRepository(db),
	}
}

func (s FilmService) FindByID(id int) (film domain.Film, err error) {
	return s.repository.FindByID(id)
}

func (s FilmService) Add(film domain.Film) (err error) {
	return s.repository.Add(film)
}

func (s FilmService) extractIDfromURL(url string) (id int, err error) {
	chunks := strings.Split(url, "/")
	lastPath := chunks[len(chunks)-2]
	if id, err = strconv.Atoi(lastPath); err != nil {
		return
	}
	return
}

func (s FilmService) LoadFilms(planetID int, filmURL string) (err error) {
	var (
		filmID       int
		respFilm     swapi.RespFilm
		releaseDate  time.Time
		filmToPlanet domain.Film
	)

	if filmID, err = s.extractIDfromURL(filmURL); err != nil {
		return
	}
	if respFilm, err = s.swapiClient.Film(filmID); err != nil {
		return
	}

	if releaseDate, err = time.Parse("2006-01-02", respFilm.ReleaseDate); err != nil {
		return
	}

	filmToPlanet = domain.Film{
		ID:          filmID,
		PlanetID:    planetID,
		ReleaseDate: releaseDate,
		Title:       respFilm.Title,
		Director:    respFilm.Director,
	}

	if err = s.repository.Add(filmToPlanet); err != nil {
		return
	}

	return
}
