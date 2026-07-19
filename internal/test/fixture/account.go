package fixture

import (
	account "github.com/Suinar/Bank-proto/repository/account"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// NewAccountCore creates a ready-to-use account core.
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

// NewAccountCoreInputId creates a ready-to-use account core input id.
func NewAccountCoreInputId(id int64) core.Account {
	return core.Account{
		Id:         id,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       TestName,
		Balance:    TestAmount,
		Status:     core.AccountStatusActive,
	}
}

// NewAccountCoreInputStatus creates a ready-to-use account core input status.
func NewAccountCoreInputStatus(status core.AccountStatus) core.Account {
	return core.Account{
		Id:         TestId,
		UserId:     TestId,
		CurrencyId: TestId,
		Name:       TestName,
		Balance:    TestAmount,
		Status:     status,
	}
}

// NewAccountUpdateInputCore creates a ready-to-use account update input core.
func NewAccountUpdateInputCore() *core.AccountUpdateInput {
	return &core.AccountUpdateInput{
		Name: StringPointer(TestName),
	}
}

// NewAccountProto creates a ready-to-use account proto.
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

// NewAccountListProto creates a ready-to-use account list proto.
func NewAccountListProto() *account.AccountList {
	return &account.AccountList{
		Accounts: []*account.Account{
			NewAccountProto(),
			NewAccountProto(),
		},
	}
}

// NewUpdateAccountInputProto creates a ready-to-use update account input proto.
func NewUpdateAccountInputProto() *account.AccountUpdateInput {
	return &account.AccountUpdateInput{
		Name: StringPointer(TestName),
	}
}

// NewUpdateAccountRequestProto creates a ready-to-use update account request proto.
func NewUpdateAccountRequestProto() *account.UpdateAccountRequest {
	return &account.UpdateAccountRequest{
		Id:    TestId,
		Input: NewUpdateAccountInputProto(),
	}
}
