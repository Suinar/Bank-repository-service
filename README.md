# Bank Repository Service

A gRPC-based repository microservice responsible for managing persistent storage for the banking system.

The service provides CRUD operations and business-specific repository actions for users, accounts, cards, credits, deposits, and currencies. It communicates with other microservices through Protocol Buffers and gRPC.

# Features

● gRPC API

● Protocol Buffers

● PostgreSQL

● Redis cache

● Repository Pattern

● Clean Architecture

● Unit Testing

● Integration Testing

● Docker & Docker Compose

● Environment configuration using Viper

● Automatic mock generation

● High test coverage

# Project Structure

<details>
<summary>Project Structure</summary>

  ```text

.
├───cmd
│   └───app
├───internal
│   ├───configs
│   ├───delivery
│   │   └───grps
│   │       └───handlers
│   │           ├───account
│   │           ├───card
│   │           ├───credit
│   │           ├───currency
│   │           ├───deposit
│   │           └───user
│   ├───mocks
│   │   ├───cache
│   │   ├───repository
│   │   └───service
│   ├───repository
│   │   ├───cache
│   │   │   └───currency
│   │   └───postgres_db
│   │       ├───account
│   │       ├───card
│   │       ├───credit
│   │       ├───currency
│   │       ├───deposit
│   │       └───user
│   ├───services
│   │   ├───account
│   │   ├───card
│   │   ├───credit
│   │   ├───currency
│   │   ├───deposit
│   │   └───user
│   └───test
│       ├───fixture
│       ├───repository
│       └───template
├───migration
├───pkg
│   ├───core
│   └───database
│       ├───cahce
│       └───postgres
└───proto
    └───repository
        ├───account
        ├───card
        ├───common
        ├───credit
        ├───currency
        ├───deposit
        └───user

 ```

</details>

# gRPC Services

The service exposes six repository APIs.

### UserRepository

##### Method	Description

##### GetAll	Returns all users

##### GetById	Returns user by ID

##### GetByEmail	Returns user by email

##### GetByPhoneNumber	Returns user by phone number

##### Create	Creates a new user

##### Update	Updates user information

##### Delete	Deletes a user

------------

### AccountRepository

##### Method	Description

##### GetAll	Returns all accounts

##### GetByUser	Returns accounts of a user

##### GetById	Returns account by ID

##### Create	Creates an account

##### Blocking	Blocks an account

##### Close	Closes an account

##### Update	Updates account

##### Delete	Deletes an account

------------

### CardRepository
 
##### Method	Description

##### GetAll	Returns all cards

##### GetByUser	Returns user cards

##### GetById	Returns card by ID

##### GetByNumber	Finds card by number

##### Blocking	Blocks card

##### Create	Creates a card

Delete	Deletes a card

------------

### CreditRepository

##### Method	Description

##### GetAll	Returns all credits

##### GetByUser	Returns user credits

##### GetById	Returns credit by ID

##### Create	Creates a credit

##### Repay	Repays credit

##### Delete	Deletes a credit

------------

#### DepositRepository

##### Method	Description

##### GetAll	Returns all deposits

##### GetByUser	Returns user deposits

##### GetById	Returns deposit by ID

##### Create	Creates a deposit

##### Replenish	Replenishes deposit

##### Delete	Deletes a deposit

------------

#### CurrencyRepository

##### Method	Description

##### GetAll	Returns all currencies

##### GetById	Returns currency by ID

##### GetByIso	Finds currency by ISO code

##### GetBySymbol	Finds currency by symbol

##### Create	Creates a currency

##### Update	Updates currency

##### Delete	Deletes a currency

# Make Commands
### Help
make help

Displays all available commands.

# Build
### make build

Builds the application.

# Run
### make run

Runs the service locally.

# Unit Tests
### make test

Runs all unit tests.
 
# Test Coverage
### make test-cover

Runs tests with coverage.

# Generate Mocks
### make generate-mocks

Generates GoMock mocks.

# Docker

### Start all services

make docker-up

### Stop all services

make docker-down

### Restart services

make docker-restart

### View logs

make docker-logs

### List running containers

make docker-ps

# Test Infrastructure

Start PostgreSQL for integration tests

make test-postgres-up

Stop PostgreSQL

make test-postgres-down

Start Redis

make test-redis-up

Stop Redis

make test-redis-down

# Testing

The project includes comprehensive testing for every application layer.

### Unit Tests

● gRPC handlers

● Services

● PostgreSQL repositories

● Redis cache

### Integration Tests

● PostgreSQL repositories

● Redis cache

### Test Tools

● GoMock

● Testify

# Configuration

Configuration is loaded using Viper.

Typical environment variables include:

APP_PORT=

POSTGRES_HOST=
POSTGRES_PORT=
POSTGRES_DB=
POSTGRES_USER=
POSTGRES_PASSWORD=

REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=

# Architecture

The service follows Clean Architecture principles.

  ```text

gRPC
    │
Handlers
    │
Services
    │
Repositories
    │  
PostgreSQL / Redis

 ```

# Dependencies

### Main libraries used in the project:

● google.golang.org/grpc
● google.golang.org/protobuf
● github.com/jmoiron/sqlx
● github.com/redis/go-redis/v9
● github.com/spf13/viper
● github.com/golang/mock
● github.com/stretchr/testify

# Highlights

● Production-ready project structure
● Clean Architecture
● gRPC communication
● Redis caching
● PostgreSQL persistence
● Repository Pattern
● Unit tests
● Integration tests
● Mock generation
● Docker support
● Protocol Buffers
● Easily extensible microservice

