PROFILE ?= local
APP_NAME ?= bootstrap
include .meta/.env
-include .meta/.env.$(PROFILE)
export

.PHONY: lint
lint:
	golangci-lint run -c .meta/.golangci.yaml --fix ./...

.PHONY: dev
dev:
	docker compose -p ${APP_NAME} -f .meta/docker-compose.yaml up --build

.PHONY: down
down:
	docker compose -p ${APP_NAME} -f .meta/docker-compose.yaml down

.PHONY: dev-run
dev-run:
	go run cmd/app/main.go

.PHONY: build
build:
	go build -o bin/app cmd/app/main.go

.PHONY: run
run:
	./bin/app

.PHONY: deploy
bin-deploy: build run