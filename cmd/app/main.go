package main

import (
	config "Bank-repository-service/internal/configs"
	handler "Bank-repository-service/internal/delivery/grps/handlers"
	cache "Bank-repository-service/internal/repository/cache"
	repository "Bank-repository-service/internal/repository/postgres_db"
	service "Bank-repository-service/internal/services"
	connectToCahce "Bank-repository-service/pkg/database/cahce"
	connectToDB "Bank-repository-service/pkg/database/postgres"

	account "Bank-repository-service/proto/repository/account"
	card "Bank-repository-service/proto/repository/card"
	credit "Bank-repository-service/proto/repository/credit"
	currency "Bank-repository-service/proto/repository/currency"
	deposit "Bank-repository-service/proto/repository/deposit"
	user "Bank-repository-service/proto/repository/user"

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
