package grpc

import (
	"log"
	"net"

	config "github.com/kVinsom/Bank-repository-service/internal/configs"
	accountHandler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/account"
	cardHandler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/card"
	creditHandler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/credit"
	currencyHandler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/currency"
	depositHandler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/deposit"
	userHandler "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/user"

	"google.golang.org/grpc"

	service "github.com/kVinsom/Bank-repository-service/internal/services"

	accountProto "github.com/Suinar/Bank-proto/repository/account"
	cardProto "github.com/Suinar/Bank-proto/repository/card"
	creditProto "github.com/Suinar/Bank-proto/repository/credit"
	currencyProto "github.com/Suinar/Bank-proto/repository/currency"
	depositProto "github.com/Suinar/Bank-proto/repository/deposit"
	userProto "github.com/Suinar/Bank-proto/repository/user"
)

// RunGrpcServer listens for gRPC requests and serves the registered APIs.
func RunGrpcServer(cfg config.Config, services *service.Services) {
	lis, err := net.Listen(cfg.Network, net.JoinHostPort(cfg.GRPCHost, cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	RegisterServices(grpcServer, services)

	log.Println("repository-services running on " + cfg.GRPCPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// RegisterServices attaches all repository handlers to the gRPC server.
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
