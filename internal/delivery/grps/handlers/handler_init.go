package handler

import (
	account "Bank-repository-service/internal/delivery/grps/handlers/account"
	card "Bank-repository-service/internal/delivery/grps/handlers/card"
	credit "Bank-repository-service/internal/delivery/grps/handlers/credit"
	currency "Bank-repository-service/internal/delivery/grps/handlers/currency"
	deposit "Bank-repository-service/internal/delivery/grps/handlers/deposit"
	user "Bank-repository-service/internal/delivery/grps/handlers/user"
	service "Bank-repository-service/internal/services"
)

type Handlers struct {
	AccountHandler  account.IAccountHandler
	CardHandler     card.ICardHandler
	CreditHandler   credit.ICreditHandler
	CurrencyHandler currency.ICurrencyHandler
	DepositHandler  deposit.IDepositHandler
	UserHandler     user.IUserHandler
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
