package handler

import (
	account "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers/account"
	card "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers/card"
	credit "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers/credit"
	currency "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers/currency"
	deposit "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers/deposit"
	user "github.com/Suinar/Bank-exhange-rate-service/internal/delivery/grps/handlers/user"
	service "github.com/Suinar/Bank-exhange-rate-service/internal/services"
)

type Handlers struct {
	AccountHandler  *account.AccountHandler
	CardHandler     *card.CardHandler
	CreditHandler   *credit.CreditHandler
	CurrencyHandler *currency.CurrencyHandler
	DepositHandler  *deposit.DepositHandler
	UserHandler     *user.UserHandler
}

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


