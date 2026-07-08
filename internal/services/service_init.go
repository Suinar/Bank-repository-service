package service

import (
	cache "Bank-repository-service/internal/repository/cache"
	repository "Bank-repository-service/internal/repository/postgres_db"
	account "Bank-repository-service/internal/services/account"
	card "Bank-repository-service/internal/services/card"
	credit "Bank-repository-service/internal/services/credit"
	currency "Bank-repository-service/internal/services/currency"
	deposit "Bank-repository-service/internal/services/deposit"
	user "Bank-repository-service/internal/services/user"
)

type Services struct {
	AccountService  *account.AccountService
	CardService     *card.CardService
	CreditService   *credit.CreditService
	CurrencyService *currency.CurrencyService
	DepositService  *deposit.DepositService
	UserService     *user.UserService
}

func InitServices(repositories *repository.Repositories, caches *cache.Caches) *Services {
	return &Services{
		AccountService:  account.NewAccountService(repositories.AccountRepository),
		CardService:     card.NewCardService(repositories.CardRepository),
		CreditService:   credit.NewCreditService(repositories.CreditRepository),
		CurrencyService: currency.NewCurrencyService(repositories.CurrencyRepository, caches.Currency),
		DepositService:  deposit.NewDepositService(repositories.DepositRepository),
		UserService:     user.NewUserService(repositories.UserRepository),
	}
}

