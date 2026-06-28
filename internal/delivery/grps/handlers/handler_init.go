package handler

import (
	account "Bank-repository-service/internal/delivery/grps/handlers/account"
	card "Bank-repository-service/internal/delivery/grps/handlers/card"
	credit "Bank-repository-service/internal/delivery/grps/handlers/credit"
	currency "Bank-repository-service/internal/delivery/grps/handlers/currency"
	deposit "Bank-repository-service/internal/delivery/grps/handlers/deposit"
	user "Bank-repository-service/internal/delivery/grps/handlers/user"
)

type Handlers struct {
	accountHandler  account.IAccountHandler
	cardHandler     card.ICardHandler
	creditHandler   credit.ICreditHandler
	currencyHandler currency.ICurrencyHandler
	depositHandler  deposit.IDepositHandler
	userHandler     user.IUserHandler
}

func InitHandlers(
	accountHandler account.IAccountHandler,
	cardHandler card.ICardHandler,
	creditHandler credit.ICreditHandler,
	currencyHandler currency.ICurrencyHandler,
	depositHandler deposit.IDepositHandler,
	userHandler user.IUserHandler) *Handlers {
	return &Handlers{
		accountHandler:  accountHandler,
		cardHandler:     cardHandler,
		creditHandler:   creditHandler,
		currencyHandler: currencyHandler,
		depositHandler:  depositHandler,
		userHandler:     userHandler,
	}
}
