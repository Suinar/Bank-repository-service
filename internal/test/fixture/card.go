package fixture

import (
	core "Bank-repository-service/pkg/core"
	card "Bank-repository-service/proto/repository/card"
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
