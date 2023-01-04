.PHONY: clean lint build run list test testv 

.DEFAULT_GOAL := help
DOCKER_TAG := latest

clean:
	rm -f bin/*

lint:
	gofmt -l -s -w .

build:
	go build --race -o ./bin/main.exe cmd/server/main.go

run:
	go run cmd/server/main.go

list:
	go list ./...

test: ## Execute tests
	go test -shuffle=on ./...

testv: ## Execute tests (detail view)
	go test -shuffle=on -v ./...

container-build: ## Build docker image for deploy
	docker build -t symthy/go-app-study:${DOCKER_TAG} --target deploy ./

container-build-local: ## build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up

down: ## Do docker compose down
	docker compose down

logs: ## Show tail docker compose logs
	docker compose logs -f

ps:	## Check container status
	docker compose ps

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
