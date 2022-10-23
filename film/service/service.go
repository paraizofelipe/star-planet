package service

import "github.com/paraizofelipe/star-planet/film/domain"

type Reader interface {
	FindByID(int) (domain.Film, error)
}

type Writer interface {
	Add(domain.Film) error
}

type Service interface {
	Reader
	Writer
}
