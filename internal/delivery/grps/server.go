package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	handler "Bank-repository-service/internal/delivery/grps/handlers"
	service "Bank-repository-service/internal/services"
	account "Bank-repository-service/proto/repository/account"
	card "Bank-repository-service/proto/repository/card"
	credit "Bank-repository-service/proto/repository/credit"
	currency `Bank-repository-service/proto/repository/currency`
	deposit "Bank-repository-service/proto/repository/deposit"
	user "Bank-repository-service/proto/repository/user"
)

func RunGrpcServer(services *service.Services) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	RegisterServices(grpcServer, services)

	log.Println("repository-service running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RegisterServices(
	grpcServer *grpc.Server,
	services *service.Services) {
	account.RegisterAccountRepositoryServer(
		grpcServer,
		handler.NewAccountHandler(services.AccountService),
	)
	card.RegisterCardRepositoryServer(
		grpcServer,
		handler.NewCardHandler(services.CardService),
	)
	credit.RegisterCreditRepositoryServer(
		grpcServer,
		handler.NewCreditHandler(services.CreditService),
	)
	currency.RegisterCurrencyRepositoryServer(
		grpcServer,
		handler.NewCurrencyHandler(services.CurrencyService),
	)
	deposit.RegisterDepositRepositoryServer(
		grpcServer,
		handler.NewDepositHandler(services.DepositService),
	)
	user.RegisterUserRepositoryServer(
		grpcServer,
		handler.NewUserHandler(services.UserService),
	)
}
