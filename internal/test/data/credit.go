package data

import "Bank-repository-service/proto/repository/credit"

func NewCredit() *credit.Credit {
	return &credit.Credit{
		Id:           1,
		UserId:       1,
		Amount:       1000,
		CurrencyId:   1,
		InterestRate: 1,
		TermMonths:   1,
		Status:       credit.CreditStatus_CREDIT_STATUS_ACTIVE,
	}
}
