package repository

import "github.com/paraizofelipe/star-planet/movie/domain"

type Reader interface {
	FindByID(id int) (domain.Movie, error)
}

type Writer interface {
	Add(domain.Movie) error
}

type MovieRepository interface {
	Reader
	Writer
}
