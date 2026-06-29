package data

import "Bank-repository-service/proto/repository/account"

func NewAccount() *account.Account {
	return &account.Account{
		Id:         1,
		UserId:     1,
		CurrencyId: 1,
		Name:       "Test_Account",
		Balance:    1000,
		Status:     account.AccountStatus_ACCOUNT_STATUS_ACTIVE,
	}
}
