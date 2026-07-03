package fixture

import (
	core "Bank-repository-service/pkg/core"
	credit "Bank-repository-service/proto/repository/credit"
)

func NewCreditCore() core.Credit {
	return core.Credit{
		Id:             TestId,
		UserId:         TestId,
		Amount:         TestAmount,
		InterestRate:   12.5,
		TermMonths:     TestTermMonths,
		MonthlyPayment: 4_700,
		Status:         core.CreditStatusActive,
	}
}

func NewCreditProto() *credit.Credit {
	return &credit.Credit{
		Id:             TestId,
		UserId:         TestId,
		Amount:         TestAmount,
		InterestRate:   12.5,
		TermMonths:     TestTermMonths,
		MonthlyPayment: 4_700,
		Status:         credit.CreditStatus_CREDIT_STATUS_ACTIVE,
	}
}

func NewCreditListProto() *credit.CreditList {
	return &credit.CreditList{
		Credits: []*credit.Credit{
			NewCreditProto(),
			NewCreditProto(),
		},
	}
}
