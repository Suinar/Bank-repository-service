package card

import (
	"context"

	card "github.com/Suinar/Bank-proto/repository/card"
	common "github.com/Suinar/Bank-proto/repository/common"
	repository "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/card"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CardService coordinates the application use cases for its domain.
type CardService struct {
	repository repository.ICardRepository
}

// NewCardService creates a ready-to-use card service.
func NewCardService(repository repository.ICardRepository) *CardService {
	return &CardService{repository: repository}
}

// GetAll returns all records available through CardService.
func (s *CardService) GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error) {
	cards, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &card.CardList{
		Cards: make([]*card.Card, len(cards)),
	}

	for i, c := range cards {
		res.Cards[i] = s.toProto(&c)
	}

	return res, nil
}

// GetByUser returns records matching the requested user lookup.
func (s *CardService) GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error) {
	cards, err := s.repository.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &card.CardList{
		Cards: make([]*card.Card, len(cards)),
	}

	for i, c := range cards {
		res.Cards[i] = s.toProto(&c)
	}

	return res, nil
}

// GetById returns records matching the requested id lookup.
func (s *CardService) GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	card, err := s.repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(card), nil
}

// GetByNumber returns records matching the requested number lookup.
func (s *CardService) GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error) {
	card, err := s.repository.GetByNumber(ctx, req.Number)
	if err != nil {
		return nil, err
	}

	return s.toProto(card), nil
}

// Blocking moves the requested record to its blocked state through CardService.
func (s *CardService) Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	card, err := s.repository.Blocking(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(card), nil
}

// Create persists a new record through CardService.
func (s *CardService) Create(ctx context.Context, req *card.Card) (*card.Card, error) {
	card, err := s.repository.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(card), nil
}

// Delete removes the requested record through CardService.
func (s *CardService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// toProto maps the domain model to its protobuf representation.
func (s *CardService) toProto(input *core.Card) *card.Card {
	if input == nil {
		return nil
	}

	return &card.Card{
		Id:          input.Id,
		UserId:      input.UserId,
		AccountId:   input.AccountId,
		Number:      input.Number,
		ExpiryMonth: int32(input.ExpiryMonth),
		ExpiryYear:  int32(input.ExpiryYear),
		Status:      card.CardStatus(input.Status),
	}
}

// fromProto maps a protobuf message to the domain model.
func (s *CardService) fromProto(input *card.Card) *core.Card {
	if input == nil {
		return nil
	}

	return &core.Card{
		Id:          input.Id,
		UserId:      input.UserId,
		AccountId:   input.AccountId,
		Number:      input.Number,
		ExpiryMonth: int8(input.ExpiryMonth),
		ExpiryYear:  int8(input.ExpiryYear),
		Status:      core.CardStatus(input.Status),
	}
}
