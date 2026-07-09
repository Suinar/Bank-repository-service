package fixture

import (
	core "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	credit "github.com/Suinar/Bank-proto/repository/credit"
)

func NewCreditCore() core.Credit {
	return core.Credit{
		Id:             TestId,
		UserId:         TestId,
		CurrencyId:     TestId,
		Amount:         TestAmount,
		MonthlyPayment: 4_700,
		Status:         core.CreditStatusActive,
	}
}

func NewCreditCoreInputId(id int64) core.Credit {
	return core.Credit{
		Id:             id,
		UserId:         TestId,
		CurrencyId:     TestId,
		Amount:         TestAmount,
		MonthlyPayment: 4_700,
		Status:         core.CreditStatusActive,
	}
}

func NewCreditCoreInputAmount(amount int64) core.Credit {
	return core.Credit{
		Id:             TestId,
		UserId:         TestId,
		CurrencyId:     TestId,
		Amount:         amount,
		MonthlyPayment: 4_700,
		Status:         core.CreditStatusActive,
	}
}

func NewCreditProto() *credit.Credit {
	return &credit.Credit{
		Id:             TestId,
		UserId:         TestId,
		Amount:         TestAmount,
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


