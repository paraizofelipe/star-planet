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
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: mockgen
mockgen:
	@go get github.com/golang/mock/mockgen@v1.5.0

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
