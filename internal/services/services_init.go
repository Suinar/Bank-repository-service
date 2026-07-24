package service

import (
	cache "github.com/kVinsom/Bank-repository-service/internal/repositories/cache"
	repository "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres"
	account "github.com/kVinsom/Bank-repository-service/internal/services/account"
	card "github.com/kVinsom/Bank-repository-service/internal/services/card"
	credit "github.com/kVinsom/Bank-repository-service/internal/services/credit"
	currency "github.com/kVinsom/Bank-repository-service/internal/services/currency"
	deposit "github.com/kVinsom/Bank-repository-service/internal/services/deposit"
	user "github.com/kVinsom/Bank-repository-service/internal/services/user"
)

// Services groups the initialized application services.
type Services struct {
	AccountService  *account.AccountService
	CardService     *card.CardService
	CreditService   *credit.CreditService
	CurrencyService *currency.CurrencyService
	DepositService  *deposit.DepositService
	UserService     *user.UserService
}

// InitServices wires the dependencies required by services.
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
