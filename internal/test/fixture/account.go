package fixture

import (
	core "Bank-repository-service/pkg/core"
	account "Bank-repository-service/proto/repository/account"
)

func NewAccountCore(name string) core.Account {
	return core.Account{
		Id:         TestId,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       name,
		Balance:    TestAmount,
		Status:     core.AccountStatusActive,
	}
}

func NewAccountUpdateInputCore(name string) *core.AccountUpdateInput {
	return &core.AccountUpdateInput{
		Name: StringPointer(name),
	}
}

func NewAccountProto(name string) *account.Account {
	return &account.Account{
		Id:         TestId,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       name,
		Balance:    TestAmount,
		Status:     account.AccountStatus_ACCOUNT_STATUS_ACTIVE,
	}
}

func NewAccountListProto(nameFirst string, nameSecond string) *account.AccountList {
	return &account.AccountList{
		Accounts: []*account.Account{
			NewAccountProto(nameFirst),
			NewAccountProto(nameSecond),
		},
	}
}

func NewUpdateAccountInputProto(name string) *account.AccountUpdateInput {
	return &account.AccountUpdateInput{
		Name: StringPointer(name),
	}
}

func NewUpdateAccountRequestProto(name string) *account.UpdateAccountRequest {
	return &account.UpdateAccountRequest{
		Id:    TestId,
		Input: NewUpdateAccountInputProto(name),
	}
}
