package service

import "github.com/paraizofelipe/star-planet/planet/domain"

type Reader interface {
	FindByName(name string) (domain.Planet, error)
	FindByID(id int) (domain.Planet, error)
	FindAll() ([]domain.Planet, error)
}

type Writer interface {
	Load(id int) error
	Add(domain.Planet) error
	RemoveByID(id int) error
}

type Service interface {
	Reader
	Writer
}
