package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/planet/service"
	"github.com/paraizofelipe/star-planet/router"
)

type Planet struct {
	Logger  *log.Logger
	Service service.Service
}

func NewHandler(db *sqlx.DB, logger *log.Logger) Handler {
	return &Planet{
		Logger:  logger,
		Service: service.NewService(db),
	}
}

func (h Planet) load(ctx *router.Context) {
	var (
		err      error
		planetID int
		paramID  string = ctx.Params["id"]
	)

	if planetID, err = strconv.Atoi(paramID); err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{"Invalid ID!"})
		return
	}

	if err = h.Service.Load(planetID); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{"Error when load the planet!"})
		return
	}

	ctx.JSON(http.StatusCreated, Response{"Planet loaded with success"})
}

func (h Planet) list(ctx *router.Context) {
	var (
		planets []domain.Planet
		err     error
	)
	if planets, err = h.Service.FindAll(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{"Error when fetching the planets!"})
		h.Logger.Print(err)
		return
	}

	ctx.JSON(http.StatusOK, planets)
}

func (h Planet) remove(ctx *router.Context) {
	var (
		planetID int
		paramID  string = ctx.Params["id"]
		err      error
	)

	if planetID, err = strconv.Atoi(paramID); err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{"Invalid ID!"})
		return
	}

	if err = h.Service.RemoveByID(planetID); err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{"Error when removing the planet!"})
		return
	}

	ctx.JSON(http.StatusOK, Response{"Planet removed."})
}

func (h Planet) findByID(ctx *router.Context) {
	var (
		err      error
		planetID int
		paramID  string = ctx.Params["id"]
		planet   domain.Planet
	)

	if planetID, err = strconv.Atoi(paramID); err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{"Invalid ID!"})
		return
	}

	if planet, err = h.Service.FindByID(planetID); err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{"Error when fetching the planet!"})
		return
	}

	if planet.ID == 0 {
		ctx.JSON(http.StatusNotFound, Response{"Planet not found!"})
		return
	}

	ctx.JSON(http.StatusOK, planet)
}

func (h Planet) findByName(ctx *router.Context) {
	var (
		err        error
		planetName string = ctx.Params["name"]
		planet     domain.Planet
	)

	if planet, err = h.Service.FindByName(planetName); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{"Error when fetching the planet!"})
		return
	}

	if planet.ID == 0 {
		ctx.JSON(http.StatusNotFound, Response{"Planet not found!"})
		return
	}

	ctx.JSON(http.StatusOK, planet)
}

func (h Planet) Router(w http.ResponseWriter, r *http.Request) {
	newRouter := router.NewRouter(h.Logger)

	newRouter.AddRoute(`planets/load/(?P<id>[\d|-]+)/?`, http.MethodPost, h.load)
	newRouter.AddRoute(`planets/id/(?P<id>[\d|-]+)/?`, http.MethodGet, h.findByID)
	newRouter.AddRoute(`planets/?`, http.MethodGet, h.list)
	newRouter.AddRoute(`planets/id/(?P<id>[\d|-]+)/?`, http.MethodDelete, h.remove)
	newRouter.AddRoute(`planets/name/(?P<name>[\w|-]+)/?`, http.MethodGet, h.findByName)

	newRouter.ServeHTTP(w, r)
}
