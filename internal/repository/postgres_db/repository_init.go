package repository

import (
	account "Bank-repository-service/internal/repository/postgres_db/account"
	card "Bank-repository-service/internal/repository/postgres_db/card"
	credit "Bank-repository-service/internal/repository/postgres_db/credit"
	currency "Bank-repository-service/internal/repository/postgres_db/currency"
	deposit "Bank-repository-service/internal/repository/postgres_db/deposit"
	user "Bank-repository-service/internal/repository/postgres_db/user"
)

type Repositories struct {
	accountRepository  account.IAccountRepository
	cardRepository     card.ICardRepository
	creditRepository   credit.ICreditRepository
	currencyRepository currency.ICurrencyRepository
	depositRepository  deposit.IDepositRepository
	userRepository     user.IUserRepository
}

func InitRepositories(
	accountRepository account.IAccountRepository,
	cardRepository card.ICardRepository,
	creditRepository credit.ICreditRepository,
	currencyRepository currency.ICurrencyRepository,
	depositRepository deposit.IDepositRepository,
	UserRepository user.IUserRepository) *Repositories {
	return &Repositories{
		accountRepository:  accountRepository,
		cardRepository:     cardRepository,
		creditRepository:   creditRepository,
		currencyRepository: currencyRepository,
		depositRepository:  depositRepository,
		userRepository:     UserRepository,
	}
}
