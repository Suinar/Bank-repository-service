package fixture

import (
	card "github.com/Suinar/Bank-proto/repository/card"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// NewCardCore creates a ready-to-use card core.
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

// NewCardCoreInputIdAndNumber creates a ready-to-use card core input id and number.
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

// NewCardCoreInputStatus creates a ready-to-use card core input status.
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

// NewCardProto creates a ready-to-use card proto.
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

// NewCardListProto creates a ready-to-use card list proto.
func NewCardListProto() *card.CardList {
	return &card.CardList{
		Cards: []*card.Card{
			NewCardProto(),
			NewCardProto(),
		},
	}
}

// NewCardNumberRequestProto creates a ready-to-use card number request proto.
func NewCardNumberRequestProto() *card.CardNumberRequest {
	return &card.CardNumberRequest{
		Number: TestCardNumber,
	}
}
