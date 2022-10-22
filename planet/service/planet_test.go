package service

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/paraizofelipe/star-planet/settings"
)

func TestPlanetServiceLoad(t *testing.T) {
	var (
		err     error
		service Service
	)

	storage, err := sqlx.Open("postgres", settings.Storage)
	if err != nil {
		t.Error(err)
	}

	service = NewService(storage)
	err = service.Load(8)
	if err != nil {
		t.Error(err)
	}
}
