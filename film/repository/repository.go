package repository

import "github.com/paraizofelipe/star-planet/film/domain"

type Reader interface {
	FindByID(id int) (domain.Film, error)
}

type Writer interface {
	Add(domain.Film) error
}

type FilmRepository interface {
	Reader
	Writer
}
