package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	account "github.com/Suinar/Bank-proto/repository/account"
	card "github.com/Suinar/Bank-proto/repository/card"
	credit "github.com/Suinar/Bank-proto/repository/credit"
	currency "github.com/Suinar/Bank-proto/repository/currency"
	deposit "github.com/Suinar/Bank-proto/repository/deposit"
	user "github.com/Suinar/Bank-proto/repository/user"
	config "github.com/kVinsom/Bank-repository-service/internal/configs"
	handler "github.com/kVinsom/Bank-repository-service/internal/delivery/grps/handlers"
	cache "github.com/kVinsom/Bank-repository-service/internal/repository/cache"
	repository "github.com/kVinsom/Bank-repository-service/internal/repository/postgres_db"
	service "github.com/kVinsom/Bank-repository-service/internal/services"
	connectToCahce "github.com/kVinsom/Bank-repository-service/pkg/database/cahce"
	connectToDB "github.com/kVinsom/Bank-repository-service/pkg/database/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// main runs the service and reports a single fatal startup or serving error.
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run owns the application resources so they are closed on every return path.
func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	startupCtx, cancelStartup := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelStartup()

	db, err := connectToDB.NewPostgresDB(startupCtx, cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	rdb := connectToCahce.NewRedisDB(cfg)
	defer rdb.Close()
	if err := rdb.Ping(startupCtx).Err(); err != nil {
		return fmt.Errorf("ping Redis: %w", err)
	}

	repositories := repository.InitRepositories(db)
	caches := cache.InitCaches(rdb)
	services := service.InitServices(repositories, caches)
	handlers := handler.InitHandlers(services)

	grpcServer := grpc.NewServer()
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go shutdown(ctx, grpcServer, healthServer, cfg.ShutdownTimeout)

	log.Printf("repository service listening on %s (%s)", address, cfg.Environment)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("serve gRPC: %w", err)
	}
	return nil
}

// shutdown marks the service unhealthy and stops gRPC within the configured deadline.
func shutdown(ctx context.Context, server *grpc.Server, healthServer *health.Server, timeout time.Duration) {
	<-ctx.Done()
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
		log.Println("gRPC server stopped")
	case <-timer.C:
		log.Println("graceful shutdown timed out; forcing stop")
		server.Stop()
	}
}
