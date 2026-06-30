package fixture

import "Bank-repository-service/proto/repository/deposit"

func NewDeposit() *deposit.Deposit {
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

func NewDepositList() *deposit.DepositList {
	return &deposit.DepositList{
		Deposits: []*deposit.Deposit{
			NewDeposit(),
			NewDeposit(),
		},
	}
}
