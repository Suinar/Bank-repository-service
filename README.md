# Bank Repository Service

Go microservice that owns persistence operations for the banking platform. It exposes repository APIs over gRPC for users, accounts, cards, credits, deposits, and currencies.

## Technologies

- Go 1.25
- gRPC and Protocol Buffers
- PostgreSQL 16
- Redis 7
- Apache Kafka 4.3 in KRaft mode (local infrastructure and configuration prepared)
- Docker and Docker Compose
- Kubernetes-ready configuration, graceful shutdown, and native gRPC health probes
- Viper
- GoMock
- Testify

> Kafka producer and consumer logic is not implemented yet. The project currently contains Kafka configuration and an optional local broker profile for the next integration step.

## Features

- Six gRPC repository APIs
- PostgreSQL repositories
- Redis currency cache
- Service and handler layers
- Environment-first configuration
- Optional local `.env` file
- Standard gRPC health service
- Graceful termination for containers and Kubernetes
- Unit and integration tests
- Reproducible GoMock generation
- Non-root distroless runtime image

## Project structure

```text
.
├── cmd/app/                         # Application entry point
├── internal/
│   ├── configs/                     # Environment and .env configuration
│   ├── delivery/grps/               # gRPC server and handlers
│   ├── mocks/                       # Generated GoMock implementations
│   ├── repository/
│   │   ├── cache/                   # Redis repositories
│   │   └── postgres_db/             # PostgreSQL repositories
│   ├── services/                    # Application services
│   └── test/                        # Fixtures and integration-test helpers
├── pkg/
│   ├── core/                        # Domain models
│   └── database/                    # PostgreSQL and Redis connections
├── docker/postgres/migration/       # Development and test database SQL
├── compose.yaml                     # PostgreSQL, Redis, Kafka, and app profiles
├── Dockerfile                       # Multi-stage non-root image
├── .env.example                     # Configuration template
└── Makefile                         # Development commands
```

## Configuration

Copy the template for local development:

```sh
cp .env.example .env
```

The `.env` file is optional. Environment variables override values from the file, which allows the same image to run in Docker, Kubernetes, or CI without embedding configuration.

Main variables:

| Variable | Default | Purpose |
|---|---:|---|
| `APP_ENV` | `local` | Environment name |
| `GRPC_HOST` | `0.0.0.0` | gRPC bind address |
| `GRPC_PORT` | `50052` | gRPC port |
| `NETWORK` | `tcp` | Listener network |
| `SHUTDOWN_TIMEOUT` | `10s` | Graceful shutdown timeout |
| `DB_HOST` | — | PostgreSQL host |
| `DB_PORT` | — | PostgreSQL port |
| `DB_USER` | — | PostgreSQL user |
| `DB_PASSWORD` | — | PostgreSQL password |
| `DB_NAME` | — | PostgreSQL database |
| `DB_SSL_MODE` | `disable` | PostgreSQL SSL mode |
| `REDIS_ADDR` | — | Redis address in `host:port` format |
| `KAFKA_ENABLED` | `false` | Enables future Kafka integration |
| `KAFKA_BROKERS` | — | Comma-separated broker addresses |
| `KAFKA_CLIENT_ID` | `bank-repository-service` | Kafka client identifier |
| `KAFKA_CONSUMER_GROUP` | `bank-repository-service` | Future consumer group |

Keep passwords and future Kafka credentials in secrets, never in Git. `.env` and `.env.test` are ignored.

## Quick start

### Local infrastructure

```sh
docker compose up -d postgres redis
go run ./cmd/app
```

### Application in Docker

```sh
docker compose --profile app up -d --build
```

### Kafka

Kafka is optional and runs only when its profile is selected:

```sh
docker compose --profile kafka up -d kafka
```

Run the application and Kafka together:

```sh
docker compose --profile app --profile kafka up -d --build
```

## Make commands

Run `make help` to display the complete list.

| Command | Description |
|---|---|
| `make build` | Build the binary into `bin/` |
| `make run` | Run the service locally |
| `make test` | Run all tests without cache |
| `make test-cover` | Generate `coverage.out` |
| `make vet` | Run `go vet` |
| `make check` | Run vet and all tests |
| `make generate-mocks` | Regenerate GoMock files |
| `make docker-build` | Build `bank-repository-service:local` |
| `make docker-up` | Start PostgreSQL and Redis |
| `make docker-down` | Remove the Compose stack |
| `make app-up` | Build and start the application stack |
| `make kafka-up` | Start the Kafka profile |
| `make test-infra-up` | Start integration-test infrastructure |

## Testing

Integration tests use:

- PostgreSQL database `bank_test` on `localhost:5432`
- Redis database `0` on `localhost:6379`

These defaults are built into the test configuration, so `.env.test` is not required. Override them in CI when necessary:

```text
TEST_DB_HOST
TEST_DB_PORT
TEST_DB_USER
TEST_DB_PASSWORD
TEST_DB_NAME
TEST_DB_SSL_MODE
TEST_REDIS_ADDR
TEST_REDIS_PASSWORD
TEST_REDIS_DB
```

Start dependencies and run all checks:

```sh
make test-infra-up
make check
```

Repository test packages share one test database and are serialized through a PostgreSQL advisory lock to avoid cross-package data conflicts.

## Test coverage

The current statement-coverage baseline is measured with running PostgreSQL and Redis dependencies:

```sh
go test ./internal/services/... ./internal/delivery/grps/handlers/... ./internal/repository/postgres_db/... ./internal/repository/cache/currency -cover
```

| Layer | Package | Coverage |
|---|---|---:|
| Service | Account | 100.0% |
| Service | Card | 100.0% |
| Service | Credit | 100.0% |
| Service | Currency | 100.0% |
| Service | Deposit | 100.0% |
| Service | User | 97.4% |
| gRPC handler | Account | 100.0% |
| gRPC handler | Card | 100.0% |
| gRPC handler | Credit | 100.0% |
| gRPC handler | Currency | 100.0% |
| gRPC handler | Deposit | 100.0% |
| gRPC handler | User | 100.0% |
| Redis cache | Currency | 100.0% |
| PostgreSQL repository | Account | 70.3% |
| PostgreSQL repository | Card | 66.0% |
| PostgreSQL repository | Credit | 67.4% |
| PostgreSQL repository | Currency | 75.0% |
| PostgreSQL repository | Deposit | 67.4% |
| PostgreSQL repository | User | 73.6% |

Generated mocks, fixtures, bootstrap packages, and test helpers are excluded from this package-level baseline because their statement percentages do not represent application behavior.
## gRPC APIs

The service provides repository operations for:

- `UserRepository`
- `AccountRepository`
- `CardRepository`
- `CreditRepository`
- `DepositRepository`
- `CurrencyRepository`

Protocol definitions are provided by the `github.com/Suinar/Bank-proto` module.

## Architecture

```text
gRPC handlers
      │
      ▼
application services
      │
      ├── PostgreSQL repositories
      └── Redis cache

Future event path:
application services ──► Kafka producer/consumer adapters
```

## Kubernetes

The binary reads configuration from environment variables, handles `SIGTERM`, stops accepting traffic before shutdown, and exposes the standard gRPC health service on port `50052`.

Example probes:

```yaml
readinessProbe:
  grpc:
    port: 50052
livenessProbe:
  grpc:
    port: 50052
```

Place non-sensitive configuration in a `ConfigMap` and credentials in a `Secret`.

## Docker image

Build manually:

```sh
docker build -t bank-repository-service:local .
```

The final image uses a non-root distroless runtime and does not contain `.env`, source files, build tools, or test data.
