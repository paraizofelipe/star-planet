package handler

import (
	"encoding/json"
	"net/http"

	"github.com/paraizofelipe/star-planet/router"
)

type Handler interface {
	load(*router.Context)
	list(*router.Context)
	remove(*router.Context)
	findByID(*router.Context)
	findByName(*router.Context)

	Router(http.ResponseWriter, *http.Request)
}

type Response struct {
	Message string `json:"message"`
}

func (r Response) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
