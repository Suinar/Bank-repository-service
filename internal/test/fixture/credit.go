package fixture

import (
	credit "github.com/Suinar/Bank-proto/repository/credit"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// NewCreditCore creates a ready-to-use credit core.
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

// NewCreditCoreInputId creates a ready-to-use credit core input id.
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

// NewCreditCoreInputAmount creates a ready-to-use credit core input amount.
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

// NewCreditProto creates a ready-to-use credit proto.
func NewCreditProto() *credit.Credit {
	return &credit.Credit{
		Id:             TestId,
		UserId:         TestId,
		CurrencyId:     TestId,
		Amount:         TestAmount,
		MonthlyPayment: 4_700,
		Status:         credit.CreditStatus_CREDIT_STATUS_ACTIVE,
	}
}

// NewCreditListProto creates a ready-to-use credit list proto.
func NewCreditListProto() *credit.CreditList {
	return &credit.CreditList{
		Credits: []*credit.Credit{
			NewCreditProto(),
			NewCreditProto(),
		},
	}
}
