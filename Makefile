APP_NAME := repository-service
IMAGE := bank-repository-service:local
COMPOSE := docker compose

.PHONY: help build run test test-cover vet check generate-mocks \
        docker-build docker-up docker-down docker-restart docker-logs docker-ps \
        app-up app-down kafka-up kafka-down \
        test-infra-up test-infra-down test-postgres-up test-postgres-down \
        test-redis-up test-redis-down

help:
	@echo "Available commands:"
	@echo ""
	@echo "Go:"
	@echo "  make build                Build the application into bin/"
	@echo "  make run                  Run the application locally"
	@echo "  make test                 Run all tests without cache"
	@echo "  make test-cover           Write test coverage to coverage.out"
	@echo "  make vet                  Run go vet"
	@echo "  make check                Run vet and tests"
	@echo "  make generate-mocks       Regenerate GoMock files"
	@echo ""
	@echo "Containers:"
	@echo "  make docker-build         Build the application image"
	@echo "  make docker-up            Start PostgreSQL and Redis"
	@echo "  make docker-down          Stop and remove the Compose stack"
	@echo "  make docker-restart       Restart PostgreSQL and Redis"
	@echo "  make docker-logs          Follow Compose logs"
	@echo "  make docker-ps            Show Compose services"
	@echo "  make app-up               Build and start the application stack"
	@echo "  make app-down             Stop the application container"
	@echo "  make kafka-up             Start Kafka (KRaft mode)"
	@echo "  make kafka-down           Stop Kafka"
	@echo ""
	@echo "Test infrastructure:"
	@echo "  make test-infra-up        Start PostgreSQL and Redis"
	@echo "  make test-infra-down      Stop PostgreSQL and Redis"
	@echo "  make test-postgres-up     Start PostgreSQL"
	@echo "  make test-postgres-down   Stop PostgreSQL"
	@echo "  make test-redis-up        Start Redis"
	@echo "  make test-redis-down      Stop Redis"

build:
	go build -o bin/$(APP_NAME) ./cmd/app
run:
	go run ./cmd/app
test:
	go clean -testcache
	go test ./...
test-cover:
	go test -coverprofile=coverage.out ./...
vet:
	go vet ./...
check: vet test
generate-mocks:
	go generate ./...
docker-build:
	docker build -t $(IMAGE) .
docker-up:
	$(COMPOSE) up -d postgres redis
docker-down:
	$(COMPOSE) down --remove-orphans
docker-restart:
	$(COMPOSE) restart postgres redis
docker-logs:
	$(COMPOSE) logs -f
docker-ps:
	$(COMPOSE) ps
app-up:
	$(COMPOSE) --profile app up -d --build
app-down:
	$(COMPOSE) --profile app stop repository-service
kafka-up:
	$(COMPOSE) --profile kafka up -d kafka
kafka-down:
	$(COMPOSE) --profile kafka stop kafka
test-infra-up: docker-up
test-infra-down:
	$(COMPOSE) stop postgres redis
test-postgres-up:
	$(COMPOSE) up -d postgres
test-postgres-down:
	$(COMPOSE) stop postgres
test-redis-up:
	$(COMPOSE) up -d redis
test-redis-down:
	$(COMPOSE) stop redis
