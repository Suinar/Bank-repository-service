.PHONY: help build run test test-cover generate-mocks \
        docker-up docker-down docker-restart docker-logs docker-ps \
        test-postgres-up test-postgres-down \
        test-redis-up test-redis-down

## ========================
## HELP
## ========================
help:
	@echo "Available commands:"
	@echo ""
	@echo "Go:"
	@echo "  make build            Build application"
	@echo "  make run              Run application"
	@echo "  make test             Run unit tests"
	@echo "  make test-cover       Run tests with coverage"
	@echo "  make generate-mocks   Generate gomock mocks"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up        Start all services"
	@echo "  make docker-down      Stop all services"
	@echo "  make docker-restart   Restart all services"
	@echo "  make docker-logs      Show logs"
	@echo "  make docker-ps        Show containers"
	@echo ""
	@echo "Test infra:"
	@echo "  make test-postgres-up"
	@echo "  make test-postgres-down"
	@echo "  make test-redis-up"
	@echo "  make test-redis-down"

## ========================
## GO
## ========================
build:
	go build -o bin/app ./cmd/app

run:
	go run ./cmd/app

test:
	go clean -testcache
	go test -v ./...

test-cover:
	go test -cover ./...

generate-mocks:
	go generate ./...

## ========================
## DOCKER
## ========================
docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-restart:
	docker compose down
	docker compose up -d

docker-logs:
	docker compose logs -f

docker-ps:
	docker compose ps

## ========================
## TEST INFRA
## ========================
test-postgres-up:
	docker compose up -d postgres_test

test-postgres-down:
	docker compose stop postgres_test

test-redis-up:
	docker compose up -d redis

test-redis-down:
	docker compose stop redis