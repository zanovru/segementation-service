-include .env
export $(shell sed 's/=.*//' .env)

DB_MIGRATIONS_SOURCE=migrations
DATABASE_URL='postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable'

DOCKER_COMPOSE=docker compose

.PHONY: build
build:
	go build -o ./build/segment-service ./cmd/main.go

run: build
	./build/segment-service

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

docker-rebuild:
	$(DOCKER_COMPOSE) up --no-deps --build -d

migration:
	migrate create -ext sql -dir $(DB_MIGRATIONS_SOURCE) -seq $(NAME)

migrate:
	migrate -path $(DB_MIGRATIONS_SOURCE) -database $(DATABASE_URL) up

migrate-drop:
	migrate -path $(DB_MIGRATIONS_SOURCE) -database $(DATABASE_URL) drop

migrate-down:
	migrate -path $(DB_MIGRATIONS_SOURCE) -database $(DATABASE_URL) down
