package service

import (
	cache "github.com/Suinar/Bank-exhange-rate-service/internal/repository/cache"
	repository "github.com/Suinar/Bank-exhange-rate-service/internal/repository/postgres_db"
	account "github.com/Suinar/Bank-exhange-rate-service/internal/services/account"
	card "github.com/Suinar/Bank-exhange-rate-service/internal/services/card"
	credit "github.com/Suinar/Bank-exhange-rate-service/internal/services/credit"
	currency "github.com/Suinar/Bank-exhange-rate-service/internal/services/currency"
	deposit "github.com/Suinar/Bank-exhange-rate-service/internal/services/deposit"
	user "github.com/Suinar/Bank-exhange-rate-service/internal/services/user"
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


