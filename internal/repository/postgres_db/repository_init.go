package repository

import (
	account "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/account"
	card "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/card"
	credit "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/credit"
	currency "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/currency"
	deposit "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/deposit"
	user "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/user"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	AccountRepository  *account.AccountRepository
	CardRepository     *card.CardRepository
	CreditRepository   *credit.CreditRepository
	CurrencyRepository *currency.CurrencyRepository
	DepositRepository  *deposit.DepositRepository
	UserRepository     *user.UserRepository
}

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



