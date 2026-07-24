package repository

import (
	account "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/account"
	card "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/card"
	credit "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/credit"
	currency "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/currency"
	deposit "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/deposit"
	user "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/user"

	"github.com/jmoiron/sqlx"
)

// Repositories groups the PostgreSQL repository dependencies used by services.
type Repositories struct {
	AccountRepository  *account.AccountRepository
	CardRepository     *card.CardRepository
	CreditRepository   *credit.CreditRepository
	CurrencyRepository *currency.CurrencyRepository
	DepositRepository  *deposit.DepositRepository
	UserRepository     *user.UserRepository
}

// InitRepositories wires the dependencies required by repositories.
func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		AccountRepository:  account.NewAccountRepository(db),
		CardRepository:     card.NewCardRepository(db),
		CreditRepository:   credit.NewCreditRepository(db),
		CurrencyRepository: currency.NewCurrencyRepository(db),
		DepositRepository:  deposit.NewDepositRepository(db),
		UserRepository:     user.NewUserRepository(db),
	}
}
