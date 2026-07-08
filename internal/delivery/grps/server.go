package grpc

import (
	config "Bank-repository-service/internal/configs"
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

	accountProto "github.com/Suinar/Bank-proto/repository/account"
	cardProto "github.com/Suinar/Bank-proto/repository/card"
	creditProto "github.com/Suinar/Bank-proto/repository/credit"
	currencyProto "github.com/Suinar/Bank-proto/repository/currency"
	depositProto "github.com/Suinar/Bank-proto/repository/deposit"
	userProto "github.com/Suinar/Bank-proto/repository/user"
)

func RunGrpcServer(cfg config.Config, services *service.Services) {
	lis, err := net.Listen(cfg.Network, cfg.GrpsPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	RegisterServices(grpcServer, services)

	log.Println("repository-service running on " + cfg.GrpsPort)

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
