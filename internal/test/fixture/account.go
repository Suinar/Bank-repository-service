package fixture

import account "Bank-repository-service/proto/repository/account"

func NewAccount(name string) *account.Account {
	return &account.Account{
		Id:         TestId,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       name,
		Balance:    TestAmount,
		Status:     account.AccountStatus_ACCOUNT_STATUS_ACTIVE,
	}
}

func NewAccountList(nameFirst string, nameSecond string) *account.AccountList {
	return &account.AccountList{
		Accounts: []*account.Account{
			NewAccount(nameFirst),
			NewAccount(nameSecond),
		},
	}
}

func NewAccountUpdateInput(name string) *account.AccountUpdateInput {
	return &account.AccountUpdateInput{
		Name: String(name),
	}
}

func NewUpdateAccountRequest(name string) *account.UpdateAccountRequest {
	return &account.UpdateAccountRequest{
		Id:    TestId,
		Input: NewAccountUpdateInput(name),
	}
}
