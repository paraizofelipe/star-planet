package repository

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/settings"
)

func TestRepositoryPlanetAdd(t *testing.T) {
	storage, err := sqlx.Open("postgres", settings.Storage)
	log.Println(settings.Storage)
	if err != nil {
		log.Panic(err)
	}

	planet := domain.Planet{
		Name: "Test",
	}

	repo := NewPostgreRepository(storage)
	repo.Add(planet)
}
