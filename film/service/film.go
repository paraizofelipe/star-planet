package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/film/domain"
	"github.com/paraizofelipe/star-planet/film/repository"
)

type FilmService struct {
	repository repository.FilmRepository
}

func NewService(db *sqlx.DB) Service {
	return &FilmService{
		repository: repository.NewPostgreRepository(db),
	}
}

func (s FilmService) FindByID(id int) (film domain.Film, err error) {
	return s.repository.FindByID(id)
}

func (s FilmService) Add(film domain.Film) (err error) {
	return s.repository.Add(film)
}

func (s FilmService) UpdateOrAdd(film domain.Film) (err error) {
	return s.repository.UpdateOrAdd(film)
}
