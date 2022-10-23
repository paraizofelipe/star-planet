package service

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/paraizofelipe/star-planet/film/repository"
	"github.com/paraizofelipe/star-planet/swapi"
)

func TestPlanetServiceLoadFilm(t *testing.T) {
	var (
		ctrl    = gomock.NewController(t)
		service FilmService
	)

	type Params struct {
		filmURL  string
		planetID int
	}

	tests := []struct {
		description string
		in          Params
		setupMock   func(Params)
		hasError    bool
	}{
		{
			description: "when loadfilm return success",
			in: Params{
				planetID: 1,
				filmURL:  "https://test.com/film/666/",
			},
			setupMock: func(params Params) {
				filmID := 666
				mockRepo := repository.NewMockFilmRepository(ctrl)
				mockClient := swapi.NewMockSWAPI(ctrl)

				mockClient.EXPECT().Film(filmID).Return(swapi.RespFilm{
					ReleaseDate: "2022-10-24",
				}, nil).AnyTimes()
				mockRepo.EXPECT().Add(gomock.Any()).Return(nil).AnyTimes()

				service = FilmService{
					repository:  mockRepo,
					swapiClient: mockClient,
				}
			},
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.setupMock(test.in)

			err := service.LoadFilms(test.in.planetID, test.in.filmURL)
			if test.hasError && err == nil {
				fmt.Println(err)
				t.Error(err)
			}
		})
	}
}
