package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	filmService "github.com/paraizofelipe/star-planet/film/service"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/planet/repository"
	"github.com/paraizofelipe/star-planet/swapi"
)

func TestPlanetServiceLoad(t *testing.T) {
	var (
		ctrl    = gomock.NewController(t)
		service PlanetService
	)

	tests := []struct {
		description string
		in          int
		setupMock   func(int)
		hasError    bool
	}{
		{
			description: "planet already exists in the database",
			in:          1,
			setupMock: func(id int) {
				repo := repository.NewMockPlanetRepository(ctrl)
				repo.EXPECT().FindByID(id).Return(domain.Planet{ID: id}, nil).AnyTimes()
				service = PlanetService{
					repository: repo,
				}
			},
			hasError: false,
		},
		{
			description: "planet not exists in the database",
			in:          1,
			setupMock: func(id int) {
				mockRepo := repository.NewMockPlanetRepository(ctrl)
				mockFilmSvc := filmService.NewMockService(ctrl)
				mockClient := swapi.NewMockSWAPI(ctrl)

				filmID := 1
				filmURL := fmt.Sprintf("http://test.com/%d/", filmID)

				mockRepo.EXPECT().FindByID(id).Return(domain.Planet{}, nil).AnyTimes()
				mockClient.EXPECT().Planet(id).Return(swapi.RespPlanet{
					Name:    "Paradise",
					Climate: "hot-summer",
					Terrain: "florest",
					FilmURLs: []string{
						filmURL,
					},
				}, nil).AnyTimes()
				mockFilmSvc.EXPECT().LoadFilms(id, filmURL).Return(nil).AnyTimes()

				mockRepo.EXPECT().Add(gomock.Any()).Return(nil).AnyTimes()

				service = PlanetService{
					swapiClient: mockClient,
					repository:  mockRepo,
					filmService: mockFilmSvc,
				}
			},
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.setupMock(test.in)

			err := service.Load(test.in)
			if test.hasError && err == nil {
				t.Error(err)
			}
		})
	}
}
