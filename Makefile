LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: deps
deps:
	@go mod tidy
	@go mod download


.PHONY: build
build:
	$(LINUX_AMD64) go build -o star-planet-api main.go

.PHONY: run
run:
	@go run ./main.go

.PHONY: install-lint
install-lint:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: mockgen 
mockgen:
	@go install github.com/golang/mock/mockgen@v1.6.0
	@mockgen -source ./film/repository/repository.go -destination ./film/repository/repository_mock.go -package repository
	@mockgen -source ./planet/repository/repository.go -destination ./planet/repository/repository_mock.go -package repository
	@mockgen -source ./film/service/service.go -destination ./film/service/service_mock.go -package service
	@mockgen -source ./planet/service/service.go -destination ./planet/service/service_mock.go -package service
	@mockgen -source ./swapi/swapi.go -destination ./swapi/swapi_mock.go -package swapi

.PHONY: build
start: build
	@go run main.go

.PHONY: test
test:
	@go test ./... -covermode=count -count 1

.PHONY: dk-start
dk-start:
	@docker run -p 3000:3000 star-planet:latest

.PHONY: dk-build
dk-build: build
	@docker build -t star-planet:latest .


.PHONY: dk-deploy
dk-deploy:
	@docker-compose up -d --build

