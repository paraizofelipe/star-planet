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

type ErrorResponse struct {
	Error string `json:"errors"`
}

type SuccessResponse struct {
	Success string `json:"success"`
}

func (r ErrorResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r SuccessResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
