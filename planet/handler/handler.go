package handler

import (
	"net/http"

	"github.com/paraizofelipe/star-planet/router"
)

type Handler interface {
	create(*router.Context)
	remove(*router.Context)
	findByID(*router.Context)
	findByName(*router.Context)

	Router(http.ResponseWriter, *http.Request)
}
