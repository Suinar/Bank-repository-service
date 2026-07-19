package fixture

import (
	deposit "github.com/Suinar/Bank-proto/repository/deposit"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// NewDepositCore creates a ready-to-use deposit core.
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

// NewDepositCoreInputId creates a ready-to-use deposit core input id.
func NewDepositCoreInputId(id int64) core.Deposit {
	return core.Deposit{
		Id:           id,
		UserId:       TestId,
		CurrencyId:   TestId,
		Amount:       TestAmount,
		InterestRate: 8.5,
		TermMonths:   TestTermMonths,
		Status:       core.DepositStatusActive,
	}
}

// NewDepositCoreInputAmount creates a ready-to-use deposit core input amount.
func NewDepositCoreInputAmount(amount int64) core.Deposit {
	return core.Deposit{
		Id:           TestId,
		UserId:       TestId,
		CurrencyId:   TestId,
		Amount:       amount,
		InterestRate: 8.5,
		TermMonths:   TestTermMonths,
		Status:       core.DepositStatusActive,
	}
}

// NewDepositProto creates a ready-to-use deposit proto.
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

// NewDepositListProto creates a ready-to-use deposit list proto.
func NewDepositListProto() *deposit.DepositList {
	return &deposit.DepositList{
		Deposits: []*deposit.Deposit{
			NewDepositProto(),
			NewDepositProto(),
		},
	}
}
