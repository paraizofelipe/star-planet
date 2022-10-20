package swapi

import (
	"fmt"
)

type RespFilm struct {
	Title         string   `json:"title"`
	EpisodeID     int      `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	StarshipURLs  []string `json:"starships"`
	VehicleURLs   []string `json:"vehicles"`
	SpeciesURLs   []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	ReleaseDate   string   `json:"release_date"`
	URL           string   `json:"url"`
}

type RespPlanet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

func (c *Client) Planet(id int) (RespPlanet, error) {
	req, err := c.newRequest(fmt.Sprintf("planets/%d", id))
	if err != nil {
		return RespPlanet{}, err
	}

	var planet RespPlanet

	if _, err = c.do(req, &planet); err != nil {
		return RespPlanet{}, err
	}

	return planet, nil
}

func (c *Client) Film(id int) (RespFilm, error) {
	req, err := c.newRequest(fmt.Sprintf("films/%d", id))
	if err != nil {
		return RespFilm{}, err
	}

	var film RespFilm

	if _, err = c.do(req, &film); err != nil {
		return RespFilm{}, err
	}

	return film, nil
}
