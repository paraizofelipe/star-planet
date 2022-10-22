package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	planet "github.com/paraizofelipe/star-planet/planet/handler"
	"github.com/paraizofelipe/star-planet/settings"
)

func main() {
	var err error

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	storage, err := sqlx.Open("postgres", settings.Storage)
	if err != nil {
		log.Panic(err)
	}

	planetHandler := planet.NewHandler(storage, logger)
	http.HandleFunc("/api/planets/", planetHandler.Router)

	url := fmt.Sprintf("%s:%d", settings.Host, settings.Port)

	log.Printf("Server listening in %s", url)

	if err = http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
