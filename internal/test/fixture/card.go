package fixture

import card "Bank-repository-service/proto/repository/card"

func NewCard() *card.Card {
	return &card.Card{
		Id:          TestId,
		UserId:      TestId,
		AccountId:   TestId,
		Number:      TestCardNumber,
		ExpiryMonth: TestTermMonths,
		ExpiryYear:  2030,
		Status:      card.CardStatus_CARD_STATUS_ACTIVE,
	}
}

func NewCardList() *card.CardList {
	return &card.CardList{
		Cards: []*card.Card{
			NewCard(),
			NewCard(),
		},
	}
}

func NewCardNumberRequest() *card.CardNumberRequest {
	return &card.CardNumberRequest{
		Number: TestCardNumber,
	}
}
