package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	account "github.com/Suinar/Bank-proto/repository/account"
	card "github.com/Suinar/Bank-proto/repository/card"
	credit "github.com/Suinar/Bank-proto/repository/credit"
	currency "github.com/Suinar/Bank-proto/repository/currency"
	deposit "github.com/Suinar/Bank-proto/repository/deposit"
	user "github.com/Suinar/Bank-proto/repository/user"
	kafkaBroker "github.com/kVinsom/Bank-repository-service/internal/brokers/kafka"
	config "github.com/kVinsom/Bank-repository-service/internal/configs"
	deliveryGRPC "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc"
	handler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler"
	appLog "github.com/kVinsom/Bank-repository-service/internal/logging/app"
	cache "github.com/kVinsom/Bank-repository-service/internal/repositories/cache"
	repository "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres"
	service "github.com/kVinsom/Bank-repository-service/internal/services"
	connectToCache "github.com/kVinsom/Bank-repository-service/pkg/database/cache"
	connectToDB "github.com/kVinsom/Bank-repository-service/pkg/database/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// main runs the service and reports a single fatal startup or serving error.
func main() {
	appLog.Configure()
	appLog.Started()
	defer appLog.Stopped()

	if err := run(); err != nil {
		appLog.GRPCServerFailed(err)
	}
}

// run owns the application resources so they are closed on every return path.
func run() error {
	appLog.ConfigurationLoading()
	cfg, err := config.LoadConfig()
	if err != nil {
		appLog.ConfigurationFailed(err)
		return fmt.Errorf("load config: %w", err)
	}
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		return checkHealth(cfg.GRPCPort)
	}
	appLog.ConfigurationLoaded(
		net.JoinHostPort(cfg.GRPCHost, cfg.GRPCPort),
		net.JoinHostPort(cfg.Postgres.Host, fmt.Sprint(cfg.Postgres.Port)),
		cfg.Redis.Addr,
		cfg.Kafka.Enabled,
		cfg.Kafka.Brokers,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	startupCtx, cancelStartup := context.WithTimeout(ctx, 5*time.Second)
	defer cancelStartup()

	appLog.DependencyConnecting("postgres", net.JoinHostPort(cfg.Postgres.Host, fmt.Sprint(cfg.Postgres.Port)))
	db, err := connectToDB.NewPostgresDB(startupCtx, cfg)
	if err != nil {
		appLog.DependencyFailed("postgres", cfg.Postgres.Host, err)
		return err
	}
	defer db.Close()
	appLog.DependencyConnected("postgres", net.JoinHostPort(cfg.Postgres.Host, fmt.Sprint(cfg.Postgres.Port)))

	appLog.DependencyConnecting("redis", cfg.Redis.Addr)
	rdb := connectToCache.NewRedisDB(cfg)
	defer rdb.Close()
	if err := rdb.Ping(startupCtx).Err(); err != nil {
		appLog.DependencyFailed("redis", cfg.Redis.Addr, err)
		return fmt.Errorf("ping Redis: %w", err)
	}
	appLog.DependencyConnected("redis", cfg.Redis.Addr)
	// Connect to the configured Kafka cluster and idempotently create all
	// repository-owned topics before accepting gRPC requests.
	if err := kafkaBroker.EnsureKafkaTopics(ctx, cfg); err != nil {
		appLog.DependencyFailed("kafka", fmt.Sprint(cfg.Kafka.Brokers), err)
		return fmt.Errorf("initialize Kafka: %w", err)
	}

	appLog.ComponentsInitializing()
	repositories := repository.InitRepositories(db)
	caches := cache.InitCaches(rdb)
	services := service.InitServices(repositories, caches)
	handlers := handler.InitHandlers(services)
	appLog.ComponentsInitialized()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(deliveryGRPC.LoggingUnaryInterceptor))
	account.RegisterAccountRepositoryServer(grpcServer, handlers.AccountHandler)
	card.RegisterCardRepositoryServer(grpcServer, handlers.CardHandler)
	currency.RegisterCurrencyRepositoryServer(grpcServer, handlers.CurrencyHandler)
	credit.RegisterCreditRepositoryServer(grpcServer, handlers.CreditHandler)
	deposit.RegisterDepositRepositoryServer(grpcServer, handlers.DepositHandler)
	user.RegisterUserRepositoryServer(grpcServer, handlers.UserHandler)

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	address := net.JoinHostPort(cfg.GRPCHost, cfg.GRPCPort)
	listener, err := net.Listen(cfg.Network, address)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", address, err)
	}
	defer listener.Close()

	go shutdown(ctx, grpcServer, healthServer, cfg.ShutdownTimeout)

	appLog.GRPCServerStarting(address)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("serve gRPC: %w", err)
	}
	return nil
}

// checkHealth lets Docker and Kubernetes probe the gRPC server without
// requiring an additional utility in the distroless runtime image.
func checkHealth(port string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	connection, err := grpc.DialContext(
		ctx,
		net.JoinHostPort("127.0.0.1", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}
	defer connection.Close()

	response, err := healthpb.NewHealthClient(connection).Check(ctx, &healthpb.HealthCheckRequest{})
	if err != nil {
		return err
	}
	if response.Status != healthpb.HealthCheckResponse_SERVING {
		return fmt.Errorf("service is not healthy: %s", response.Status)
	}
	return nil
}

// shutdown marks the service unhealthy and stops gRPC within the configured deadline.
func shutdown(ctx context.Context, server *grpc.Server, healthServer *health.Server, timeout time.Duration) {
	<-ctx.Done()
	appLog.ShutdownSignalReceived()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)

	stopped := make(chan struct{})
	go func() {
		server.GracefulStop()
		close(stopped)
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-stopped:
		appLog.GRPCServerStopped()
	case <-timer.C:
		appLog.ShutdownForced(timeout)
		server.Stop()
	}
}
