package fixture

import (
	core "Bank-repository-service/pkg/core"
	deposit "Bank-repository-service/proto/repository/deposit"
)

func NewDepositCore() core.Deposit {
	return core.Deposit{
		Id:           TestId,
		UserId:       TestId,
		CurrencyId:   TestId,
		Amount:       TestAmount,
		InterestRate: 8.5,
		TermMonths:   TestTermMonths,
		Status:       core.DepositStatusActive,
	}
}

func NewDepositProto() *deposit.Deposit {
	return &deposit.Deposit{
		Id:           TestId,
		UserId:       TestId,
		CurrencyId:   TestId,
		Amount:       TestAmount,
		InterestRate: 8.5,
		TermMonths:   TestTermMonths,
		Status:       deposit.DepositStatus_DEPOSIT_STATUS_ACTIVE,
	}
}

func NewDepositListProto() *deposit.DepositList {
	return &deposit.DepositList{
		Deposits: []*deposit.Deposit{
			NewDepositProto(),
			NewDepositProto(),
		},
	}
}
