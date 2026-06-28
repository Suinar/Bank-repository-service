package service

import (
	account "Bank-repository-service/internal/services/account"
	card "Bank-repository-service/internal/services/card"
	credit "Bank-repository-service/internal/services/credit"
	currency "Bank-repository-service/internal/services/currency"
	deposit "Bank-repository-service/internal/services/deposit"
	user "Bank-repository-service/internal/services/user"
)

type Services struct {
	AccountService  account.IAccountService
	CardService     card.ICardService
	CreditService   credit.ICreditService
	CurrencyService currency.ICurrencyService
	DepositService  deposit.IDepositService
	UserService     user.IUserService
}

func InitServices(
	accountService account.IAccountService,
	cardService card.ICardService,
	creditService credit.ICreditService,
	currencyService currency.ICurrencyService,
	depositService deposit.IDepositService,
	userService user.IUserService) *Services {
	return &Services{
		AccountService:  accountService,
		CardService:     cardService,
		CreditService:   creditService,
		CurrencyService: currencyService,
		DepositService:  depositService,
		UserService:     userService,
	}
}
