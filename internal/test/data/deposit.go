package data

import "Bank-repository-service/proto/repository/deposit"

func NewDeposit() *deposit.Deposit {
	return &deposit.Deposit{
		Id:           1,
		UserId:       1,
		Amount:       1000,
		CurrencyId:   1,
		InterestRate: 1,
		TermMonths:   1,
		Status:       deposit.DepositStatus_DEPOSIT_STATUS_ACTIVE,
	}
}
