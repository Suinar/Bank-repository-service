package repository

import (
	account "Bank-repository-service/internal/repository/postgres_db/account"
	card "Bank-repository-service/internal/repository/postgres_db/card"
	credit "Bank-repository-service/internal/repository/postgres_db/credit"
	currency "Bank-repository-service/internal/repository/postgres_db/currency"
	deposit "Bank-repository-service/internal/repository/postgres_db/deposit"
	user "Bank-repository-service/internal/repository/postgres_db/user"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	AccountRepository  account.IAccountRepository
	CardRepository     card.ICardRepository
	CreditRepository   credit.ICreditRepository
	CurrencyRepository currency.ICurrencyRepository
	DepositRepository  deposit.IDepositRepository
	UserRepository     user.IUserRepository
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
