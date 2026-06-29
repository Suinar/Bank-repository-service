package data

import "Bank-repository-service/proto/repository/card"

func NewCard() *card.Card {
	return &card.Card{
		Id:          1,
		UserId:      1,
		AccountId:   1,
		Number:      "1234567890123456",
		ExpiryMonth: 1,
		ExpiryYear:  1,
		Status:      card.CardStatus_CARD_STATUS_ACTIVE,
	}
}
