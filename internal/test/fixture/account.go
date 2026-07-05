package fixture

import (
	core "Bank-repository-service/pkg/core"
	account "Bank-repository-service/proto/repository/account"
)

func NewAccountCore() core.Account {
	return core.Account{
		Id:         TestId,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       TestName,
		Balance:    TestAmount,
		Status:     core.AccountStatusActive,
	}
}

func NewAccountUpdateInputCore() *core.AccountUpdateInput {
	return &core.AccountUpdateInput{
		Name: StringPointer(TestName),
	}
}

func NewAccountProto() *account.Account {
	return &account.Account{
		Id:         TestId,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       TestName,
		Balance:    TestAmount,
		Status:     account.AccountStatus_ACCOUNT_STATUS_ACTIVE,
	}
}

func NewAccountListProto() *account.AccountList {
	return &account.AccountList{
		Accounts: []*account.Account{
			NewAccountProto(),
			NewAccountProto(),
		},
	}
}

func NewUpdateAccountInputProto() *account.AccountUpdateInput {
	return &account.AccountUpdateInput{
		Name: StringPointer(TestName),
	}
}

func NewUpdateAccountRequestProto() *account.UpdateAccountRequest {
	return &account.UpdateAccountRequest{
		Id:    TestId,
		Input: NewUpdateAccountInputProto(),
	}
}
