package main

import (
	config "github.com/Suinar/Bank-exhange-rate-service/internal/configs"
	handler "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers"
	cache "github.com/Suinar/Bank-exhange-rate-service/internal/repository/cache"
	repository "github.com/Suinar/Bank-exhange-rate-service/internal/repository/postgres_db"
	service "github.com/Suinar/Bank-exhange-rate-service/internal/services"
	connectToCahce "github.com/Suinar/Bank-exhange-rate-service/pkg/database/cahce"
	connectToDB "github.com/Suinar/Bank-exhange-rate-service/pkg/database/postgres"

	account "github.com/Suinar/Bank-proto/repository/account"
	card "github.com/Suinar/Bank-proto/repository/card"
	credit "github.com/Suinar/Bank-proto/repository/credit"
	currency "github.com/Suinar/Bank-proto/repository/currency"
	deposit "github.com/Suinar/Bank-proto/repository/deposit"
	user "github.com/Suinar/Bank-proto/repository/user"

	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	// Postgres
	db := connectToDB.NewPostgresDB(cfg)

	// Redis
	rdb := connectToCahce.NewRedisDB(cfg)

	defer func() {
		_ = rdb.Close()
		_ = db.Close()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// redis health check
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("redis not connected:", err)
	}

	repositories := repository.InitRepositories(db)
	caches := cache.InitCaches(rdb)

	services := service.InitServices(repositories, caches)

	handlers := handler.InitHandlers(services)

	// gRPC server
	grpcServer := grpc.NewServer()

	account.RegisterAccountRepositoryServer(grpcServer, handlers.AccountHandler)
	card.RegisterCardRepositoryServer(grpcServer, handlers.CardHandler)
	currency.RegisterCurrencyRepositoryServer(grpcServer, handlers.CurrencyHandler)
	credit.RegisterCreditRepositoryServer(grpcServer, handlers.CreditHandler)
	deposit.RegisterDepositRepositoryServer(grpcServer, handlers.DepositHandler)
	user.RegisterUserRepositoryServer(grpcServer, handlers.UserHandler)

	lis, err := net.Listen("tcp", ":"+cfg.ApiPort)
	if err != nil {
		log.Fatal(err)
	}

	// graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		log.Println("shutdown gRPC server...")

		stopped := make(chan struct{})

		go func() {
			grpcServer.GracefulStop()
			close(stopped)
		}()

		select {
		case <-stopped:
		case <-time.After(5 * time.Second):
			log.Println("force stop gRPC server")
			grpcServer.Stop()
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

