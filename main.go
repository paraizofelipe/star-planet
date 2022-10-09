package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paraizofelipe/star-planet/settings"
)

func main() {
	var err error

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	url := fmt.Sprintf("%s:%d", settings.Host, settings.Port)

	if err = http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
