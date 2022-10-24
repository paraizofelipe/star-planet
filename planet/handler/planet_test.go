package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	filmDomain "github.com/paraizofelipe/star-planet/film/domain"
	"github.com/paraizofelipe/star-planet/planet/domain"
	"github.com/paraizofelipe/star-planet/planet/service"
	"github.com/paraizofelipe/star-planet/router"
)

func TestHandlerLoadSucces(t *testing.T) {
	logger := log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		requestID          int
		expectedStatusCode int
		setupMock          func(int) *service.MockService
		context            router.Context
	}{
		{
			description:        "when run load return success",
			requestID:          666,
			expectedStatusCode: 201,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().Load(id).Return(nil).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
		{
			description:        "when run load return error",
			requestID:          666,
			expectedStatusCode: 500,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().Load(id).Return(errors.New("Error")).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
		{
			description:        "when run load return invalid ID",
			expectedStatusCode: 500,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": "AAAA",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				mockService = test.setupMock(test.requestID)
				w           = httptest.NewRecorder()
			)

			handler := Planet{
				Logger:  logger,
				Service: mockService,
			}

			ctx := test.context
			ctx.ResponseWriter = w

			handler.load(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestHandlerList(t *testing.T) {
	logger := log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	t.Run("when run list return success", func(t *testing.T) {
		w := httptest.NewRecorder()
		expectedStatusCode := 200
		expectedBody := []domain.Planet{
			{
				Name:    "Paradise",
				Climate: "hot-summer",
				Terrain: "florest",
				Films: []filmDomain.Film{
					{
						PlanetID:    1,
						Title:       "Return of Paradise",
						Director:    "Paraizo",
						ReleaseDate: time.Now(),
					},
				},
			},
		}

		ctx := router.Context{
			ResponseWriter: w,
		}

		ctrl := gomock.NewController(t)
		service := service.NewMockService(ctrl)
		service.EXPECT().FindAll().Return(expectedBody, nil).Times(1)

		handler := Planet{
			Logger:  logger,
			Service: service,
		}

		handler.list(&ctx)
		resp := w.Result()

		if resp.StatusCode != expectedStatusCode {
			t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, expectedStatusCode)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		j, err := json.Marshal(expectedBody)
		if err != nil {
			t.Error(err)
		}

		if string(b) != string(j) {
			t.Errorf("cuttent: %s ---> expected: %s", string(b), string(j))
		}
	})

	t.Run("when run list return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		expectedStatusCode := 500
		expectedBody := Response{Message: "Error when fetching the planets!"}

		ctx := router.Context{
			ResponseWriter: w,
		}

		ctrl := gomock.NewController(t)
		service := service.NewMockService(ctrl)
		service.EXPECT().FindAll().Return(nil, errors.New("Error")).Times(1)

		handler := Planet{
			Logger:  logger,
			Service: service,
		}

		handler.list(&ctx)
		resp := w.Result()

		if resp.StatusCode != expectedStatusCode {
			t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, expectedStatusCode)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		j, err := json.Marshal(expectedBody)
		if err != nil {
			t.Error(err)
		}

		if string(b) != string(j) {
			t.Errorf("cuttent: %s ---> expected: %s", string(b), string(j))
		}
	})
}

func TestHandlerRemove(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.Lshortfile)

	tests := []struct {
		description        string
		requestID          int
		expectedStatusCode int
		setupMock          func(int) *service.MockService
		context            router.Context
	}{
		{
			description:        "when run remove return success",
			requestID:          666,
			expectedStatusCode: 200,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().RemoveByID(id).Return(nil).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
		{
			description:        "when run list return error",
			requestID:          666,
			expectedStatusCode: 500,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().RemoveByID(id).Return(errors.New("Error")).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
		{
			description:        "when run list return invaid ID",
			requestID:          666,
			expectedStatusCode: 500,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": "AAAA",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				mockService = test.setupMock(test.requestID)
				w           = httptest.NewRecorder()
			)

			handler := Planet{
				Logger:  logger,
				Service: mockService,
			}

			ctx := test.context
			ctx.ResponseWriter = w

			handler.remove(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestHandlerFindByID(t *testing.T) {
	logger := log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		requestID          int
		expectedStatusCode int
		setupMock          func(int) *service.MockService
		context            router.Context
	}{
		{
			description:        "when run findByID return success",
			requestID:          666,
			expectedStatusCode: 200,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().FindByID(id).Return(domain.Planet{ID: id}, nil).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
		{
			description:        "when run findByID return error",
			requestID:          666,
			expectedStatusCode: 500,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().FindByID(id).Return(domain.Planet{}, errors.New("error")).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
		{
			description:        "when run findByID return invalid ID",
			expectedStatusCode: 500,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": "AAAA",
				},
			},
		},
		{
			description:        "when run findByID return not found",
			requestID:          666,
			expectedStatusCode: 404,
			setupMock: func(id int) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().FindByID(id).Return(domain.Planet{}, nil).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"id": strconv.Itoa(666),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				mockService = test.setupMock(test.requestID)
				w           = httptest.NewRecorder()
			)

			handler := Planet{
				Logger:  logger,
				Service: mockService,
			}

			ctx := test.context
			ctx.ResponseWriter = w

			handler.findByID(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestHandlerName(t *testing.T) {
	logger := log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		requestName        string
		expectedStatusCode int
		setupMock          func(string) *service.MockService
		context            router.Context
	}{
		{
			description:        "when run findByName return success",
			requestName:        "Paradise",
			expectedStatusCode: 200,
			setupMock: func(name string) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().FindByName(name).Return(domain.Planet{ID: 1}, nil).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"name": "Paradise",
				},
			},
		},
		{
			description:        "when run findByName return error",
			requestName:        "Paradise",
			expectedStatusCode: 500,
			setupMock: func(name string) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().FindByName(name).Return(domain.Planet{}, errors.New("error")).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"name": "Paradise",
				},
			},
		},
		{
			description:        "when run findByName return not found",
			requestName:        "Paradise",
			expectedStatusCode: 404,
			setupMock: func(name string) *service.MockService {
				ctrl := gomock.NewController(t)
				mockService := service.NewMockService(ctrl)
				mockService.EXPECT().FindByName(name).Return(domain.Planet{}, nil).Times(1)
				return mockService
			},
			context: router.Context{
				Params: map[string]string{
					"name": "Paradise",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				mockService = test.setupMock(test.requestName)
				w           = httptest.NewRecorder()
			)

			handler := Planet{
				Logger:  logger,
				Service: mockService,
			}

			ctx := test.context
			ctx.ResponseWriter = w

			handler.findByName(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}
