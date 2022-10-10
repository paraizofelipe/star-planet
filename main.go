package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/settings"
)

func main() {
	var err error

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	storage, err := sqlx.Open("postgres", settings.Storage)
	if err != nil {
		log.Panic(err)
	}

	url := fmt.Sprintf("%s:%s", settings.Host, settings.Port)

	if err = http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
