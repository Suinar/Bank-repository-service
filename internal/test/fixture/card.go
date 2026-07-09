package fixture

import (
	core "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	card "github.com/Suinar/Bank-proto/repository/card"
)

func NewCardCore() core.Card {
	return core.Card{
		Id:          TestId,
		UserId:      TestId,
		AccountId:   TestId,
		Number:      TestCardNumber,
		ExpiryMonth: TestTermMonths,
		ExpiryYear:  30,
		Status:      core.CardStatusActive,
	}
}

func NewCardCoreInputIdAndNumber(id int64, number string) core.Card {
	return core.Card{
		Id:          id,
		UserId:      TestId,
		AccountId:   TestId,
		Number:      number,
		ExpiryMonth: TestTermMonths,
		ExpiryYear:  30,
		Status:      core.CardStatusActive,
	}
}

func NewCardCoreInputStatus(status core.CardStatus) core.Card {
	return core.Card{
		Id:          TestId,
		UserId:      TestId,
		AccountId:   TestId,
		Number:      TestCardNumber,
		ExpiryMonth: TestTermMonths,
		ExpiryYear:  30,
		Status:      status,
	}
}

func NewCardProto() *card.Card {
	return &card.Card{
		Id:          TestId,
		UserId:      TestId,
		AccountId:   TestId,
		Number:      TestCardNumber,
		ExpiryMonth: TestTermMonths,
		ExpiryYear:  30,
		Status:      card.CardStatus_CARD_STATUS_ACTIVE,
	}
}

func NewCardListProto() *card.CardList {
	return &card.CardList{
		Cards: []*card.Card{
			NewCardProto(),
			NewCardProto(),
		},
	}
}

func NewCardNumberRequestProto() *card.CardNumberRequest {
	return &card.CardNumberRequest{
		Number: TestCardNumber,
	}
}


