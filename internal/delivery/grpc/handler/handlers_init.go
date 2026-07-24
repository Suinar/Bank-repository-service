package handler

import (
	account "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/account"
	card "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/card"
	credit "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/credit"
	currency "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/currency"
	deposit "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/deposit"
	user "github.com/kVinsom/Bank-repository-service/internal/delivery/grpc/handler/user"
	service "github.com/kVinsom/Bank-repository-service/internal/services"
)

// Handlers groups the gRPC handlers registered by the server.
type Handlers struct {
	AccountHandler  *account.AccountHandler
	CardHandler     *card.CardHandler
	CreditHandler   *credit.CreditHandler
	CurrencyHandler *currency.CurrencyHandler
	DepositHandler  *deposit.DepositHandler
	UserHandler     *user.UserHandler
}

// InitHandlers wires the dependencies required by handlers.
func InitHandlers(services *service.Services) *Handlers {
	return &Handlers{
		AccountHandler:  account.NewAccountHandler(services.AccountService),
		CardHandler:     card.NewCardHandler(services.CardService),
		CreditHandler:   credit.NewCreditHandler(services.CreditService),
		CurrencyHandler: currency.NewCurrencyHandler(services.CurrencyService),
		DepositHandler:  deposit.NewDepositHandler(services.DepositService),
		UserHandler:     user.NewUserHandler(services.UserService),
	}
}
