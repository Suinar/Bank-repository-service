package grpc

import (
	accountHandler "Bank-repository-service/internal/delivery/grps/handlers/account"
	cardHandler "Bank-repository-service/internal/delivery/grps/handlers/card"
	creditHandler "Bank-repository-service/internal/delivery/grps/handlers/credit"
	currencyHandler "Bank-repository-service/internal/delivery/grps/handlers/currency"
	depositHandler "Bank-repository-service/internal/delivery/grps/handlers/deposit"
	userHandler "Bank-repository-service/internal/delivery/grps/handlers/user"
	"log"
	"net"

	"google.golang.org/grpc"

	service "Bank-repository-service/internal/services"
	accountProto "Bank-repository-service/proto/repository/account"
	cardProto "Bank-repository-service/proto/repository/card"
	creditProto "Bank-repository-service/proto/repository/credit"
	currencyProto "Bank-repository-service/proto/repository/currency"
	depositProto "Bank-repository-service/proto/repository/deposit"
	userProto "Bank-repository-service/proto/repository/user"
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
	accountProto.RegisterAccountRepositoryServer(
		grpcServer,
		accountHandler.NewAccountHandler(services.AccountService),
	)
	cardProto.RegisterCardRepositoryServer(
		grpcServer,
		cardHandler.NewCardHandler(services.CardService),
	)
	creditProto.RegisterCreditRepositoryServer(
		grpcServer,
		creditHandler.NewCreditHandler(services.CreditService),
	)
	currencyProto.RegisterCurrencyRepositoryServer(
		grpcServer,
		currencyHandler.NewCurrencyHandler(services.CurrencyService),
	)
	depositProto.RegisterDepositRepositoryServer(
		grpcServer,
		depositHandler.NewDepositHandler(services.DepositService),
	)
	userProto.RegisterUserRepositoryServer(
		grpcServer,
		userHandler.NewUserHandler(services.UserService),
	)
}
