APP_NAME := repository-service
IMAGE := bank-repository-service:local
COMPOSE := docker compose -f docker/docker-compose.yml
K8S_NAMESPACE := bank
K8S_DIR := docker/kubernetes

.PHONY: help build run test test-cover vet check generate-mocks \
        docker-build docker-up docker-down docker-restart docker-logs docker-ps \
        app-up app-down kafka-up kafka-down \
        k8s-apply k8s-delete k8s-status k8s-logs k8s-port-forward \
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
	@echo "  make k8s-apply            Deploy repository resources to Kubernetes"
	@echo "  make k8s-delete           Remove repository resources from Kubernetes"
	@echo "  make k8s-status           Show repository Kubernetes resources"
	@echo "  make k8s-logs             Follow repository pod logs"
	@echo "  make k8s-port-forward     Forward localhost:50052 to the gRPC service"
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
	docker build -f docker/app/Dockerfile -t $(IMAGE) .
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
k8s-apply:
	kubectl apply -f $(K8S_DIR)
k8s-delete:
	kubectl delete -f $(K8S_DIR) --ignore-not-found
k8s-status:
	kubectl get deployment,pod,service,pvc,hpa,pdb -n $(K8S_NAMESPACE) -l app.kubernetes.io/part-of=bank
k8s-logs:
	kubectl logs -f deployment/bank-repository -n $(K8S_NAMESPACE)
k8s-port-forward:
	kubectl port-forward service/bank-repository 50052:50052 -n $(K8S_NAMESPACE)
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
