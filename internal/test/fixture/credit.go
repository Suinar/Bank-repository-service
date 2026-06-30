package fixture

import credit "Bank-repository-service/proto/repository/credit"

func NewCredit() *credit.Credit {
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

func NewCreditList() *credit.CreditList {
	return &credit.CreditList{
		Credits: []*credit.Credit{
			NewCredit(),
			NewCredit(),
		},
	}
}
