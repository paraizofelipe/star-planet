package repository

import (
	"github.com/paraizofelipe/star-planet/planet/domain"
)

type Reader interface {
	FindByName(name string) (domain.Planet, error)
	FindByID(id int) (domain.Planet, error)
	FindAll() ([]domain.Planet, error)
}

type Writer interface {
	Add(domain.Planet) error
	RemoveByID(id int) error
	UpdateOrAdd(domain.Planet) error
}

type PlanetRepository interface {
	Reader
	Writer
}
